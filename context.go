package crudinator

import (
	"context"
	"net/http"
	"reflect"
)

type ContextKey string

const (
	PersistentStoreKey ContextKey = "CRUDPersistentStore"
)

type Context struct {
	Validators      *map[reflect.Type]map[string]ValidatorFunc
	EventSink       EventSink
	PersistentStore PersistentStore
	context.Context
}

func newContext(core *Core) Context {
	ctx := Context{}
	fillContext(core, &ctx)
	return ctx
}

func fillContext(core *Core, ctx *Context) {
	sessionStore := core.PersistentStore.Session()
	if sessionStore == nil {
		sessionStore = core.PersistentStore
	}
	ctx.PersistentStore = sessionStore
	ctx.Validators = &core.Validators
	ctx.EventSink = core.EventSink
}

func ExtractPersistentStore(ctx context.Context) PersistentStore {
	val := ctx.Value(PersistentStoreKey)
	return val.(PersistentStore)
}

func injectReqCtx(core *Core, req *http.Request) *http.Request {
	ctx := req.Context()
	coreCTX := newContext(core)

	ctx = context.WithValue(ctx, PersistentStoreKey, coreCTX.PersistentStore)

	return req.WithContext(ctx)
}
