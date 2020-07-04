package crudinator

import (
	"log"
	"os"
	"reflect"

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
func New() *Core {
	core := Core{}

	core.Config = ParseConfig()

	core.Resources = make(map[string]interface{})

	core.Router = web.New(Context{})
	core.Router.Middleware(makeInitContextMw(&core))

	return &core
}

func (core *Core) RegisterValidator(name string, validator ValidatorFunc) {

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

}

func (core *Core) RegisterPersistentStore(ps PersistentStore) {
	core.PersistentStore = ps
}
