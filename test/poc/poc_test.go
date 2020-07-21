package main

import (
	"log"
	"testing"

	"github.com/dionb/crudinator"
	"github.com/dionb/crudinator/test/fixtures"
	"gotest.tools/assert"
)

func init() {
	log.SetFlags(log.Llongfile)
}

type Resource1 struct {
	ID   string
	Name string
}

func TestBasicPoC(t *testing.T) {
	core, _ := crudinator.New()

	core.RegisterPersistentStore(&fixtures.FuzzerPS{})
	core.RegisterResource(Resource1{})

	for _, test := range []fixtures.Case{
		{
			Name: "GET 1",
			URL:  "/Resource1/SomeID",
		},
		{
			Name: "GET Many",
			URL:  "/Resource1",
		},
		{
			Name: "POST",
			URL:  "/Resource1/SomeID",
		},
	} {
		fixtures.RunTest(t, test, core.Router)
	}
}

func TestLocalPostgres(t *testing.T) {
	conf := crudinator.Config{
		PersistentStore: crudinator.PersistentStoreConfig{
			Engine:   "postgres",
			Username: "crudinator",
			Password: "password",
			Schema:   "crudinator",
		},
	}
	core, err := crudinator.New(conf)
	assert.NilError(t, err)

	core.RegisterResource(Resource1{})

	for _, test := range []fixtures.Case{
		{
			Name: "Get empty",
			URL:  "/Resource1/doesNotExist",
		},
	} {
		fixtures.RunTest(t, test, core.Router)
	}
}
