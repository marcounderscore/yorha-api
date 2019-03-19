package routes

import (
	"errors"
	"iris"
	"yorha-api/datamodels"
	"yorha-api/services"
)

// PodPrograms returns list of the podPrograms. http://localhost:8080/podprograms
func PodPrograms(service services.PodProgramService) (results []datamodels.PodProgram) {
	return service.GetAll()
}

// PodProgramByID return the current podProgram. http://localhost:8080/podprograms/1
func PodProgramByID(service services.PodProgramService, id uint) (podProgram datamodels.PodProgram, found bool) {
	return service.GetByID(id)
}

// InsertPodProgram create a new record in the database. http://localhost:8080/podprograms
func InsertPodProgram(ctx iris.Context, service services.PodProgramService) (datamodels.PodProgram, error) {
	var podProgram datamodels.PodProgram

	if err := ctx.ReadJSON(&podProgram); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.PodProgram{}, errors.New("Json post request not correct at all")
	}

	return service.Create(podProgram.Name, podProgram.Program, podProgram.Cooldown, podProgram.Photo)
}

// UpdatePodProgram modify a current record in the database. http://localhost:8080/podprograms/1
func UpdatePodProgram(ctx iris.Context, service services.PodProgramService, id uint) (datamodels.PodProgram, bool, error) {
	var podProgram datamodels.PodProgram

	if err := ctx.ReadJSON(&podProgram); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.PodProgram{}, true, errors.New("Json put request not correct at all")
	}

	return service.Update(id, podProgram.Name, podProgram.Program, podProgram.Cooldown, podProgram.Photo)
}

// DeletePodProgram deletes the current PodProgram. http://localhost:8080/podprograms/1
func DeletePodProgram(service services.PodProgramService, id uint) (podProgram datamodels.PodProgram, found bool) {
	return service.DeleteByID(id)
}
