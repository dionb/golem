package fixtures

import (
	"errors"
	"log"
	"reflect"

	"github.com/dionb/crudinator"
	fuzz "github.com/google/gofuzz"
)

type NopPS struct{ crudinator.PersistentStore }

type FuzzerPS struct {
	crudinator.PersistentStore
}

type interfaceLoader struct {
	v   interface{}
	typ reflect.Type
}

func (fps *FuzzerPS) Get(key interface{}, tableName string, dst interface{}) error {
	fuzzer := fuzz.New()
	fuzzer.Fuzz(dst)
	val := reflect.ValueOf(dst).Elem()
	id := key.(string)
	idField, found := val.Type().FieldByName("ID")
	if !found {
		log.Println("no ID field found in resource to be fuzzed")
		return nil
	}
	val.FieldByIndex(idField.Index).SetString(tableName + "." + id)
	return nil
}

func (fps *FuzzerPS) List(tableName string, filters map[string]interface{}, dst interface{}) error {
	val := reflect.ValueOf(dst).Elem()
	log.Printf("%#v\n", val)

	if val.Kind() != reflect.Slice {
		return errors.New("need slice to list")
	}

	var elemType reflect.Type
	elemType = val.Type().Elem()

	log.Println(elemType.Name())
	fuzzer := fuzz.New()
	for i := 0; i < 20; i++ {
		entry := reflect.New(elemType)
		fuzzer.Fuzz(entry.Interface())
		val.Set(reflect.Append(val, entry.Elem()))
	}
	return nil
}

func (fps *FuzzerPS) Set(key string, tablename string, value interface{}) {
	log.Printf("saving: %s.%s: %#v", tablename, key, value)
}

func (fps *FuzzerPS) Session() crudinator.PersistentStore {
	return fps
}
