package crudinator

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"
)

// func (ctx *Context) makeDefaultGetHandler(t reflect.Type) http.Handler {

// 	return func(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {

// 	}

// }

type DefaultHandler struct {
	resourceType reflect.Type
}

func makeDefaultGetHandleFunc(res interface{}, core *Core) httprouter.Handle {
	resourceType := reflect.TypeOf(res)
	return func(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ctx := newContext(core)

		key := params.ByName("id")
		ps := ctx.PersistentStore
		resourceName := resourceType.Name()
		value := reflect.New(resourceType).Interface()

		if HandleError(rw, ps.Get(key, resourceName, value)) {
			return
		}

		encoder := json.NewEncoder(rw)
		HandleError(rw, encoder.Encode(value))
	}
}

func injectMiddlewareStd(core *Core, handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		req = injectReqCtx(core, req)
		handleFunc(rw, req)
	}
}
