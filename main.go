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

	hero.Register(
		services.NewWeaponService(
			repositories.NewWeaponRepository(
				datasource.Database,
			),
		),
	)

	hero.Register(
		services.NewPodProgramService(
			repositories.NewPodProgramRepository(
				datasource.Database,
			),
		),
	)

	hero.Register(
		services.NewBossService(
			repositories.NewBossRepository(
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

	app.PartyFunc("/weapons", func(r iris.Party) {
		r.Use(middleware.BasicAuth)
		r.Get("/", hero.Handler(routes.Weapons))
		r.Get("/{id: uint}", hero.Handler(routes.WeaponByID))
		r.Post("/", hero.Handler(routes.InsertWeapon))
		r.Put("/{id: uint}", hero.Handler(routes.UpdateWeapon))
		r.Delete("/{id: uint}", hero.Handler(routes.DeleteWeapon))
	})

	app.PartyFunc("/podprograms", func(r iris.Party) {
		r.Use(middleware.BasicAuth)
		r.Get("/", hero.Handler(routes.PodPrograms))
		r.Get("/{id: uint}", hero.Handler(routes.PodProgramByID))
		r.Post("/", hero.Handler(routes.InsertPodProgram))
		r.Put("/{id: uint}", hero.Handler(routes.UpdatePodProgram))
		r.Delete("/{id: uint}", hero.Handler(routes.DeletePodProgram))
	})

	app.PartyFunc("/bosses", func(r iris.Party) {
		r.Use(middleware.BasicAuth)
		r.Get("/", hero.Handler(routes.Bosses))
		r.Get("/{id: uint}", hero.Handler(routes.BossByID))
		r.Post("/", hero.Handler(routes.InsertBoss))
		r.Put("/{id: uint}", hero.Handler(routes.UpdateBoss))
		r.Delete("/{id: uint}", hero.Handler(routes.DeleteBoss))
	})

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
