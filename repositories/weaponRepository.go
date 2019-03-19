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

	var response []datamodels.Weapon
	for _, item := range results {
		var class datamodels.Class
		r.source.First(&class, item.ClassID)
		r.source.Model(&item).Association("Class").Append(class)
		response = append(response, item)
	}

	return response
}

func (r *weaponMemoryRepository) Select(id uint) (weapon datamodels.Weapon, found bool) {
	r.source.First(&weapon, id)
	if weapon.ID != 0 {
		var class datamodels.Class
		r.source.First(&class, weapon.ClassID)
		r.source.Model(&weapon).Association("Class").Append(class)
		return weapon, true
	}

	return
}

func (r *weaponMemoryRepository) Insert(weapon datamodels.Weapon) (datamodels.Weapon, error) {
	var class datamodels.Class
	r.source.First(&class, weapon.ClassID)
	if class.ID == 0 {
		return datamodels.Weapon{}, errors.New("Foreign key constraint fail for Race id")
	}

	r.source.Create(&datamodels.Weapon{
		Name:        weapon.Name,
		Class:       weapon.Class,
		Description: weapon.Description,
		Ability:     weapon.Ability,
		Photo:       weapon.Photo,
		ClassID:     weapon.ClassID,
	})
	r.source.Last(&weapon)

	return weapon, nil
}

func (r *weaponMemoryRepository) Update(id uint, weapon datamodels.Weapon) (datamodels.Weapon, bool, error) {
	var class datamodels.Class
	r.source.First(&class, weapon.ClassID)
	if class.ID == 0 {
		return datamodels.Weapon{}, true, errors.New("Foreign key constraint fail for Class id")
	}

	var current datamodels.Weapon
	r.source.First(&current, id)
	if current.ID != 0 {
		r.source.Model(&current).Updates(datamodels.Weapon{
			Name:        weapon.Name,
			Class:       weapon.Class,
			Description: weapon.Description,
			Ability:     weapon.Ability,
			Photo:       weapon.Photo,
			ClassID:     weapon.ClassID,
		})
		return current, true, nil
	}

	return datamodels.Weapon{}, false, errors.New("Not found record")
}

func (r *weaponMemoryRepository) Delete(id uint) (weapon datamodels.Weapon, found bool) {
	r.source.First(&weapon, id)
	if weapon.ID != 0 {
		r.source.Delete(&weapon)
		return weapon, true
	}

	return
}
