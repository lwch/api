package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const RequestTimeout = 10 * time.Second

type Context struct {
	w      http.ResponseWriter
	r      *http.Request
	values context.Context
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w:      w,
		r:      r,
		values: context.Background(),
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

func (ctx *Context) BodyFrom(r io.Reader) {
	io.Copy(ctx.w, r)
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

func (ctx *Context) SetContentType(str string) {
	ctx.w.Header().Set("Content-Type", str)
}

func (ctx *Context) SetContentDisposition(name string) {
	ctx.w.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
}

func (ctx *Context) AddValue(k, v interface{}) {
	ctx.values = context.WithValue(ctx.values, k, v)
}

func (ctx *Context) Value(k interface{}) interface{} {
	return ctx.values.Value(k)
}
