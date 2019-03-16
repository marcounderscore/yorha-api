package routes

import (
	"yorha-api/datamodels"
	"yorha-api/services"
)

// Automatas returns list of the Automata. http://localhost:8080/automatas
func Automatas(service services.AutomataService) (results []datamodels.Automata) {
	return service.GetAll()
}
