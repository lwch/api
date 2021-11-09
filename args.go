package api

import (
	"mime/multipart"
	"strconv"
	"strings"
)

func (ctx *Context) XInt(name string) int {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return int(n)
}

func (ctx *Context) OInt(name string, def int) int {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	n, _ := strconv.ParseInt(v, 10, 64)
	return int(n)
}

func (ctx *Context) XUInt(name string) uint {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	n, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return uint(n)
}

func (ctx *Context) OUInt(name string, def uint) uint {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	n, _ := strconv.ParseInt(v, 10, 64)
	return uint(n)
}

func (ctx *Context) XStr(name string) string {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	return v
}

func (ctx *Context) OStr(name, def string) string {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	return v
}

func transCsv(data []string) []string {
	for i, v := range data {
		data[i] = strings.ReplaceAll(v, "%2c%", ",")
	}
	return data
}

func (ctx *Context) XCsv(name string) []string {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	return transCsv(strings.Split(v, ","))
}

func (ctx *Context) OCsv(name string, def []string) []string {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	return transCsv(strings.Split(v, ","))
}

func (ctx *Context) XBool(name string) bool {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return b
}

func (ctx *Context) OBool(name string, def bool) bool {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	b, _ := strconv.ParseBool(v)
	return b
}

func (ctx *Context) File(name string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.r.FormFile(name)
}
