package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"

	"github.com/laoshancun/gitry/controller"
	"github.com/laoshancun/gitry/model"
	"github.com/laoshancun/gitry/service"
)

func main() {
	app := iris.New()

	// serve our app in public, public folder
	// contains the client-side vue.js application,
	// no need for any server-side template here,
	// actually if you're going to just use vue without any
	// back-end services, you can just stop afer this line and start the server.
	app.StaticWeb("/", "./wwwroot")

	// configure the http sessions.
	sess := sessions.New(sessions.Config{
		Cookie: "iris_session",
	})

	// configure the websocket server.
	ws := websocket.New(websocket.Config{})

	todosRouter := app.Party("/todos")
	// http://localhost:8080/todos/iris-ws.js
	// serve the javascript client library to communicate with
	// the iris high level websocket event system.
	todosRouter.Any("/iris-ws.js", websocket.ClientHandler())
	// create our mvc application targeted to /todos relative sub path.
	todosApp := mvc.New(todosRouter)

	// any dependencies bindings here...
	todosApp.Register(
		service.NewMemoryTodoService(),
		sess.Start,
		ws.Upgrade,
	)

	// controllers registration here...
	todosApp.Handle(new(controllers.TodoController))

	tmpl := iris.HTML("./views", ".html")
	tmpl.Layout("layouts/layout.html")
	tmpl.Reload(true)
	// tmpl.Binary(Asset, AssetNames)
	app.RegisterView(tmpl)

	app.Handle("GET", "/", func(ctx iris.Context) {
		page := model.IndexPage{Message: "Hello world!"}
		page.Title = "home page"
		ctx.ViewData("", page)
		if err := ctx.View("index.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	app.Handle("GET", "/", func(ctx iris.Context) {
		page := model.IndexPage{Message: "Hello world!"}
		page.Title = "home page"
		ctx.ViewData("", page)
		if err := ctx.View("index.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	app.Handle("GET", "/verifycode", func(ctx iris.Context) {
		catpcha := service.Captcha{}
		id, image := catpcha.GenerateCaptcha("")
		ctx.JSON(iris.Map{"state": true, "data": image, "token": id})
	})

	app.Handle("POST", "/verifycode", func(ctx iris.Context) {
		catpcha := service.Captcha{}
		code := ctx.PostValue("code")
		token := ctx.PostValue("token")
		state, err := catpcha.VerfiyCaptcha(token, code)
		ctx.JSON(iris.Map{"state": state, "msg": err})
	})

	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/app.yml")))
}
