package repositories

import (
	"errors"
	"yorha-api/datamodels"

	"github.com/jinzhu/gorm"
)

// WeaponRepository handles the basic operations of an entity/model.
type WeaponRepository interface {
	SelectMany() (results []datamodels.Weapon)
	Select(id uint) (weapon datamodels.Weapon, found bool)
	Insert(weapon datamodels.Weapon) (insertedRecord datamodels.Weapon, err error)
	Update(id uint, weapon datamodels.Weapon) (updatedRecord datamodels.Weapon, found bool, err error)
	Delete(id uint) (weapon datamodels.Weapon, found bool)
}

// NewWeaponRepository returns a new memory-based repository,
func NewWeaponRepository(source *gorm.DB) WeaponRepository {
	return &weaponMemoryRepository{source: source}
}

type weaponMemoryRepository struct {
	source *gorm.DB
}

func (r *weaponMemoryRepository) SelectMany() (results []datamodels.Weapon) {
	r.source.Find(&results)

	return
}

func (r *weaponMemoryRepository) Select(id uint) (weapon datamodels.Weapon, found bool) {
	var weapons []datamodels.Weapon
	r.source.Find(&weapons)
	for _, item := range weapons {
		if item.ID == id {
			weapon = item
			return weapon, true
		}
	}

	return
}

func (r *weaponMemoryRepository) Insert(weapon datamodels.Weapon) (datamodels.Weapon, error) {
	r.source.Create(&datamodels.Weapon{
		Name:        weapon.Name,
		Class:       weapon.Class,
		Description: weapon.Description,
		Ability:     weapon.Ability,
		Photo:       weapon.Photo,
	})
	r.source.Last(&weapon)

	return weapon, nil
}

func (r *weaponMemoryRepository) Update(id uint, weapon datamodels.Weapon) (datamodels.Weapon, bool, error) {
	var weapons []datamodels.Weapon
	var current datamodels.Weapon
	r.source.Find(&weapons)
	for _, item := range weapons {
		if item.ID == id {
			current = item
			r.source.Model(&current).Updates(datamodels.Weapon{
				Name:        weapon.Name,
				Class:       weapon.Class,
				Description: weapon.Description,
				Ability:     weapon.Ability,
				Photo:       weapon.Photo,
			})
			return current, true, nil
		}
	}

	return datamodels.Weapon{}, false, errors.New("Not found record")
}

func (r *weaponMemoryRepository) Delete(id uint) (weapon datamodels.Weapon, found bool) {
	var weapons []datamodels.Weapon
	r.source.Find(&weapons)
	for _, item := range weapons {
		if item.ID == id {
			weapon = item
			r.source.Delete(&item)
			return weapon, true
		}
	}

	return
}
