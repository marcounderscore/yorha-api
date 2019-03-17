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

	// Repository is created with the data from data source
	repo := repositories.NewAutomataRepository(datasource.Automatas)
	// Automata service will bind it to the automata app's dependencies.
	automataService := services.NewAutomataService(repo)
	hero.Register(automataService)

	app.PartyFunc("/automatas", func(r iris.Party) {
		r.Use(middleware.BasicAuth)
		r.Get("/", hero.Handler(routes.Automatas))
		r.Get("/{id: uint}", hero.Handler(routes.AutomataByID))
	})

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
