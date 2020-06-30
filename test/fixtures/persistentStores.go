package fixtures

import (
	"log"
	"reflect"

	"github.com/dionb/crudinator"
	fuzz "github.com/google/gofuzz"
)

type NopPS struct{ crudinator.PersistentStore }

type FuzzerPS struct {
	crudinator.PersistentStore
}

func (fps *FuzzerPS) Get(key interface{}, tableName string, dst interface{}) error {
	fuzzer := fuzz.New()
	fuzzer.Fuzz(dst)
	val := reflect.ValueOf(dst).Elem()
	id := key.(string)
	// log.Println(val.Elem().Type().Name())
	idField, found := val.Type().FieldByName("ID")
	if !found {
		log.Println("no ID field found in resource to be fuzzed")
		return nil
	}
	val.FieldByIndex(idField.Index).SetString(tableName + "." + id)
	return nil
}

func (fps *FuzzerPS) Session() crudinator.PersistentStore {
	return fps
}
