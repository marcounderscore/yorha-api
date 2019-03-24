package routes

import (
	"errors"
	"iris"
	"yorha-api/datamodels"
	"yorha-api/services"
)

// Weapons returns list of the Weapons. http://localhost:8080/weapons
func Weapons(service services.WeaponService) (results []datamodels.Weapon) {
	return service.GetAll()
}

// WeaponByID return the current Weapon. http://localhost:8080/weapons/1
func WeaponByID(service services.WeaponService, id uint) (weapon datamodels.Weapon, found bool) {
	return service.GetByID(id)
}

// InsertWeapon create a new record in the database. http://localhost:8080/weapons
func InsertWeapon(ctx iris.Context, service services.WeaponService) (datamodels.Weapon, error) {
	var weapon datamodels.Weapon

	if err := ctx.ReadJSON(&weapon); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.Weapon{}, errors.New("Json post request not correct at all")
	}

	return service.Create(weapon.Name, weapon.MinAttack, weapon.MaxAttack, weapon.Ability, weapon.Photo, weapon.ClassID)
}

// UpdateWeapon modify a current record in the database. http://localhost:8080/weapons/1
func UpdateWeapon(ctx iris.Context, service services.WeaponService, id uint) (datamodels.Weapon, bool, error) {
	var weapon datamodels.Weapon

	if err := ctx.ReadJSON(&weapon); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.Weapon{}, true, errors.New("Json put request not correct at all")
	}

	return service.Update(id, weapon.Name, weapon.MinAttack, weapon.MaxAttack, weapon.Ability, weapon.Photo, weapon.ClassID)
}

// DeleteWeapon deletes the current Weapon. http://localhost:8080/weapons/1
func DeleteWeapon(service services.WeaponService, id uint) (weapon datamodels.Weapon, found bool) {
	return service.DeleteByID(id)
}
