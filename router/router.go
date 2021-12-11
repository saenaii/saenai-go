package router

import (
	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application) {
	app.PartyFunc("/channel", func(p iris.Party) {
		p.Get("/exists", channel.Exists)
	})
}