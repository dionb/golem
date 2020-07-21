package crudinator

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/dionb/crudinator/storage/gocraftPStore"
	"github.com/gocraft/web"
)

type Core struct {
	Validators      map[reflect.Type]map[string]ValidatorFunc
	Resources       map[string]interface{}
	Router          *web.Router
	EventSink       EventSink
	PersistentStore PersistentStore
	OauthProvider   AuthProvider
	Config          Config
}

// New will return a new Core object with the config parsed and EventSink, PersistentStore, Oauth Provider
func New(conf ...Config) (*Core, error) {
	core := Core{}

	if len(conf) > 0 {
		core.Config = conf[0]
	} else {
		core.Config = ParseConfig()
	}

	core.Resources = make(map[string]interface{})

	core.Router = web.New(Context{})
	core.Router.Middleware(makeInitContextMw(&core))

	var ps PersistentStore
	switch core.Config.PersistentStore.Engine {
	case "mysql":
		fallthrough
	case "postgres":
		ps = gocraftPStore.New().(PersistentStore)

	}
	err := core.RegisterPersistentStore(ps)
	if err != nil {
		return nil, fmt.Errorf("initialising persistent store: %w", err)
	}

	return &core, nil
}

func (core *Core) RegisterValidator(name string, validator ValidatorFunc) {
	// core.Validators[name] = validator
}

func (core *Core) RegisterResource(res interface{}) {
	resType := reflect.TypeOf(res)
	if resType.Kind() != reflect.Struct {
		log.Println("only struct types can be registered")
		os.Exit(2)
	}

	name := resType.Name()

	if _, ok := core.Resources[name]; ok {
		log.Println("a resource with the name \"" + name + "\" already exists")
		os.Exit(2)
	}

	core.Resources[name] = res

	if getter, ok := res.(StdGetHandler); ok {
		core.Router.Get("/"+name+"/:id", injectMiddlewareStd(core, getter.Get))
	} else {
		core.Router.Get("/"+name+"/:id", makeDefaultGetHandleFunc(res, core))
	}

	if lister, ok := res.(StdListHandler); ok {
		core.Router.Get("/"+name, injectMiddlewareStd(core, lister.List))
	} else {
		core.Router.Get("/"+name, makeDefaultListHandleFunc(res, core))
	}

	if setter, ok := res.(StdSetHandler); ok {
		core.Router.Post("/"+name, injectMiddlewareStd(core, setter.Set))
	} else {
		core.Router.Post("/"+name, makeDefaultSetHandleFunc(res, core))
	}
}

func (core *Core) RegisterPersistentStore(ps PersistentStore) error {
	err := ps.Connect(core.Config.PersistentStore)
	if err != nil {
		return fmt.Errorf("verifying connection details: %w", err)
	}

	core.PersistentStore = ps
	return nil
}
