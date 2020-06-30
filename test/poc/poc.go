package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/dionb/crudinator"
	"github.com/dionb/crudinator/test/fixtures"
)

func main() {
	log.SetFlags(log.Llongfile)
	t := testing.T{}
	basicPoCTest(&t)
}

func basicPoCTest(t *testing.T) {
	core := crudinator.New()

	core.RegisterPersistentStore(&fixtures.FuzzerPS{})
	core.RegisterResource(Resource1{})

	http.ListenAndServe("localhost:8080", core.Router)
}

type Resource1 struct {
	ID   string
	Name string
}
