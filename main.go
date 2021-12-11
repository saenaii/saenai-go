package main

import (
	"github.com/kataras/iris/v12"

	"{{app_name}}/router"
)

func main() {
	app := iris.Default()

	router.Register(app)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Listen(":8080")
}
