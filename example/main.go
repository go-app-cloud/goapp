package main

import (
	"github.com/go-app-cloud/goapp"
	"log"
)

const (
	html = ".html"
)

func main() {
	app := goapp.Default()
	//app.HandleDir("/public", "./assets/public", goapp.DirOptions{Asset: Asset, AssetInfo: AssetInfo, AssetNames: AssetNames})
	app.HandleDir("/public", "./assets/public")
	//app.RegisterView(goapp.HTML("./assets/view", html).Binary(Asset, AssetNames))
	app.RegisterView(goapp.HTML("./assets/view", html).Reload(true))

	app.Get("/", func(ctx goapp.Context) {
		_ = ctx.View("index.html")
	})

	if err := app.Run(goapp.Addr(":9990"), nil); err != nil {
		log.Fatal(err)
	}
}
