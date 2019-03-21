package routes

import (
	"errors"
	"iris"
	"yorha-api/datamodels"
	"yorha-api/services"
)

// Bosses returns list of the Bosses. http://localhost:8080/Bosses
func Bosses(service services.BossService) (results []datamodels.Boss) {
	return service.GetAll()
}

// BossByID return the current Boss. http://localhost:8080/Bosses/1
func BossByID(service services.BossService, id uint) (boss datamodels.Boss, found bool) {
	return service.GetByID(id)
}

// InsertBoss create a new record in the database. http://localhost:8080/Bosss
func InsertBoss(ctx iris.Context, service services.BossService) (datamodels.Boss, error) {
	var boss datamodels.Boss

	if err := ctx.ReadJSON(&boss); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.Boss{}, errors.New("Json post request not correct at all")
	}

	return service.Create(boss.Name, boss.Faction, boss.Zones, boss.Photo)
}

// UpdateBoss modify a current record in the database. http://localhost:8080/Bosss/1
func UpdateBoss(ctx iris.Context, service services.BossService, id uint) (datamodels.Boss, bool, error) {
	var boss datamodels.Boss

	if err := ctx.ReadJSON(&boss); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.Boss{}, true, errors.New("Json put request not correct at all")
	}

	return service.Update(id, boss.Name, boss.Faction, boss.Zones, boss.Photo)
}

// DeleteBoss deletes the current Boss. http://localhost:8080/Bosss/1
func DeleteBoss(service services.BossService, id uint) (boss datamodels.Boss, found bool) {
	return service.DeleteByID(id)
}
