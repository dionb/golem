package crudinator

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/gocraft/web"
)

// func (ctx *Context) makeDefaultGetHandler(t reflect.Type) http.Handler {

// 	return func(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {

// 	}

// }

type DefaultHandler struct {
	resourceType reflect.Type
}

func makeDefaultListHandleFunc(res interface{}, core *Core) HandleFunc {
	resourceType := reflect.TypeOf(res)
	return func(ctx *Context, rw web.ResponseWriter, req *web.Request) {

		filters := make(map[string]interface{}, 0)
		decoder := json.NewDecoder(req.Body)
		if req.ContentLength > 0 && HandleError(rw, decoder.Decode(&filters)) {
			return
		}

		ps := ctx.PersistentStore
		resourceName := resourceType.Name()

		value := reflect.New(reflect.SliceOf(resourceType))

		if HandleError(rw, ps.List(resourceName, filters, value.Interface())) {
			return
		}

		log.Println(value)

		encoder := json.NewEncoder(rw)
		HandleError(rw, encoder.Encode(value))
	}
}

func makeDefaultGetHandleFunc(res interface{}, core *Core) HandleFunc {
	resourceType := reflect.TypeOf(res)
	return func(ctx *Context, rw web.ResponseWriter, req *web.Request) {

		key := req.PathParams["id"]
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

func makeDefaultPostHandleFunc(res interface{}, core *Core) HandleFunc {
	resourceType := reflect.TypeOf(res)
	return func(ctx *Context, rw web.ResponseWriter, req *web.Request) {

		key := req.PathParams["id"]
		ps := ctx.PersistentStore
		resourceName := resourceType.Name()
		value := reflect.New(resourceType).Interface()

		decoder := json.NewDecoder(req.Body)
		if HandleError(rw, decoder.Decode(value)) {
			return
		}

		if HandleError(rw, ps.Set(key, resourceName, value)) {
			return
		}

		encoder := json.NewEncoder(rw)
		HandleError(rw, encoder.Encode(value))
	}
}

func makeDefaultDeleteHandleFunc(res interface{}, core *Core) HandleFunc {
	resourceType := reflect.TypeOf(res)
	return func(ctx *Context, rw web.ResponseWriter, req *web.Request) {

		key := req.PathParams["id"]
		ps := ctx.PersistentStore
		resourceName := resourceType.Name()

		if HandleError(rw, ps.Delete(key, resourceName)) {
			return
		}
	}
}

func injectMiddlewareStd(core *Core, handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		req = injectReqCtx(core, req)
		handleFunc(rw, req)
	}
}
