package routes

import (
	"iris"
	"yorha-api/datamodels"
	"yorha-api/services"
)

// Automatas returns list of the automatas. http://localhost:8080/automatas
func Automatas(service services.AutomataService) (results []datamodels.Automata) {
	return service.GetAll()
}

// AutomataByID return the current automata. http://localhost:8080/automatas/1
func AutomataByID(service services.AutomataService, id uint) (automata datamodels.Automata, found bool) {
	return service.GetByID(id)
}

// InsertAutomata create a new record in the database. http://localhost:8080/automatas
func InsertAutomata(ctx iris.Context, service services.AutomataService) (datamodels.Automata, error) {
	var automata datamodels.Automata

	if err := ctx.ReadJSON(&automata); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return datamodels.Automata{}, err
	}

	name := automata.Name
	occupation := automata.Occupation
	race := automata.Race
	photo := automata.Photo

	return service.Create(name, occupation, race, photo)
}
