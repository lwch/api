package api

import (
	"encoding/json"
	"net/http"
	"time"
)

const RequestTimeout = 10 * time.Second

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w: w,
		r: r,
	}
}

func (ctx *Context) OK(payload interface{}) {
	ctx.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.w).Encode(map[string]interface{}{
		"code":    0,
		"payload": payload,
	})
}

func (ctx *Context) ERR(code int, msg string) {
	ctx.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.w).Encode(map[string]interface{}{
		"code": code,
		"msg":  msg,
	})
}

func (ctx *Context) Body(data []byte) {
	ctx.w.Write(data)
}

func (ctx *Context) NotFound(what string) {
	panic(NotFound(what))
}

func (ctx *Context) Timeout() {
	panic(Timeout{})
}

func (ctx *Context) URI() string {
	return ctx.r.URL.Path
}

func (ctx *Context) ServeFile(dir string) {
	http.ServeFile(ctx.w, ctx.r, dir)
}

func (ctx *Context) HTTPNotFound(what string) {
	http.Error(ctx.w, what+" not found", http.StatusNotFound)
}

func (ctx *Context) HTTPServiceUnavailable(msg string) {
	http.Error(ctx.w, msg, http.StatusServiceUnavailable)
}

func (ctx *Context) HTTPTimeout() {
	http.Error(ctx.w, "timeout", http.StatusRequestTimeout)
}

func (ctx *Context) HTTPConflict(msg string) {
	http.Error(ctx.w, msg, http.StatusConflict)
}

func (ctx *Context) HTTPForbidden(msg string) {
	http.Error(ctx.w, msg, http.StatusForbidden)
}

func (ctx *Context) Token() string {
	return ctx.r.Header.Get("X-Token")
}