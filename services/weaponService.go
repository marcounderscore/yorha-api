package services

import (
	"yorha-api/datamodels"
	"yorha-api/repositories"
)

// WeaponService handles some of the CRUID operations of datamodel
type WeaponService interface {
	GetAll() []datamodels.Weapon
	GetByID(id uint) (datamodels.Weapon, bool)
	Create(name string, minAttack int, maxAttack int, ability string, photo string, classID uint) (datamodels.Weapon, error)
	Update(id uint, name string, minAttack int, maxAttack int, ability string, photo string, classID uint) (datamodels.Weapon, bool, error)
	DeleteByID(id uint) (datamodels.Weapon, bool)
}

// NewWeaponService returns the default service
func NewWeaponService(repo repositories.WeaponRepository) WeaponService {
	return &weaponService{
		repo: repo,
	}
}

type weaponService struct {
	repo repositories.WeaponRepository
}

func (s *weaponService) GetAll() []datamodels.Weapon {
	return s.repo.SelectMany()
}

func (s *weaponService) GetByID(id uint) (datamodels.Weapon, bool) {
	return s.repo.Select(id)
}

func (s *weaponService) Create(name string, minAttack int, maxAttack int, ability string, photo string, classID uint) (datamodels.Weapon, error) {
	return s.repo.Insert(datamodels.Weapon{
		Name:      name,
		MinAttack: minAttack,
		MaxAttack: maxAttack,
		Ability:   ability,
		Photo:     photo,
		ClassID:   classID,
	})
}

func (s *weaponService) Update(id uint, name string, minAttack int, maxAttack int, ability string, photo string, classID uint) (datamodels.Weapon, bool, error) {
	return s.repo.Update(id, datamodels.Weapon{
		Name:      name,
		MinAttack: minAttack,
		MaxAttack: maxAttack,
		Ability:   ability,
		Photo:     photo,
		ClassID:   classID,
	})
}

func (s *weaponService) DeleteByID(id uint) (datamodels.Weapon, bool) {
	return s.repo.Delete(id)
}
