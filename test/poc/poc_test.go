package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/dionb/crudinator"
	"github.com/dionb/crudinator/test/fixtures"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestBasicGET(t *testing.T) {
	core := crudinator.New()

	core.RegisterPersistentStore(&fixtures.FuzzerPS{})
	core.RegisterResource(Resource1{})

	req, _ := http.NewRequest("GET", "/Resource1", nil)
	rw := fixtures.MockRW{}
	core.Router.ServeHTTP(&rw, req)

	// http.ListenAndServe("localhost:8080", core.Router)
	// t.Fail()
}

func TestBasicPOST(t *testing.T) {
	core := crudinator.New()

	core.RegisterPersistentStore(&fixtures.FuzzerPS{})
	core.RegisterResource(Resource1{})

	req, _ := http.NewRequest("POST", "/Resource1/id", nil)
	rw := fixtures.MockRW{}
	core.Router.ServeHTTP(&rw, req)

	// http.ListenAndServe("localhost:8080", core.Router)
	t.Fail()
}

type Resource1 struct {
	ID   string
	Name string
}
