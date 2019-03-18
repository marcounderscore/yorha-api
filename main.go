package main

import (
	"yorha-api/datasource"
	"yorha-api/repositories"
	"yorha-api/services"
	"yorha-api/web/middleware"
	"yorha-api/web/routes"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	hero.Register(
		services.NewAutomataService(
			repositories.NewAutomataRepository(
				datasource.Database,
			),
		),
	)

	app.PartyFunc("/automatas", func(r iris.Party) {
		r.Use(middleware.BasicAuth)
		r.Get("/", hero.Handler(routes.Automatas))
		r.Get("/{id: uint}", hero.Handler(routes.AutomataByID))
		r.Post("/", hero.Handler(routes.InsertAutomata))
		r.Put("/{id: uint}", hero.Handler(routes.UpdateAutomata))
		r.Delete("/{id: uint}", hero.Handler(routes.DeleteAutomata))
	})

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
