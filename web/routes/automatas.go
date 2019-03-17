package routes

import (
	"yorha-api/datamodels"
	"yorha-api/services"
)

// Automatas returns list of the Automata. http://localhost:8080/automatas
func Automatas(service services.AutomataService) (results []datamodels.Automata) {
	return service.GetAll()
}

// AutomataByID return the current automata
func AutomataByID(service services.AutomataService, id uint) (automata datamodels.Automata, found bool) {
	return service.GetByID(id)
}
