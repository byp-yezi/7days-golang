package main

import (
	"7days-golang/day2-context/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.Run(":9999")
}