package main

import (
	"github.com/kataras/iris"
)

var addr = iris.Addr("localhost:8080")

func main() {
	app := iris.New()
	app.Get("/", index)
	app.Run(addr)
}

func index(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello World"})
}