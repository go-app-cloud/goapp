package main

import (
	"github.com/go-app-cloud/goapp"
	"log"
)

func main() {
	app := goapp.Default()
	app.Get("/", func(ctx goapp.Context) {
		_, _ = ctx.JSON(goapp.Map{
			"success": "true",
		})
	})
	if err := app.Run(goapp.Addr(":9990"), nil); err != nil {
		log.Fatal(err)
	}
}
