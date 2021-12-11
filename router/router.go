package router

import (
	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application) {
	app.PartyFunc("/path", func(p iris.Party) {
		p.Get("/path", func() {})
	})
}