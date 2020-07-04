package crudinator

import (
	"github.com/gocraft/web"
)

func makeInitContextMw(core *Core) func(*Context, web.ResponseWriter, *web.Request, web.NextMiddlewareFunc) {
	return func(ctx *Context, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		fillContext(core, ctx)

		next(rw, req)
	}
}
