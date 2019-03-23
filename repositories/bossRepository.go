package repositories

import (
	"errors"
	"yorha-api/datamodels"

	"github.com/jinzhu/gorm"
)

// BossRepository handles the basic operations of an entity/model.
type BossRepository interface {
	SelectMany() (results []datamodels.Boss)
	Select(id uint) (boss datamodels.Boss, found bool)
	Insert(boss datamodels.Boss) (insertedRecord datamodels.Boss, err error)
	Update(id uint, boss datamodels.Boss) (updatedRecord datamodels.Boss, found bool, err error)
	Delete(id uint) (boss datamodels.Boss, found bool)
}

// NewBossRepository returns a new memory-based repository,
func NewBossRepository(source *gorm.DB) BossRepository {
	return &bossMemoryRepository{source: source}
}

type bossMemoryRepository struct {
	source *gorm.DB
}

func (r *bossMemoryRepository) SelectMany() (results []datamodels.Boss) {
	r.source.Find(&results)

	var response []datamodels.Boss
	var zones []datamodels.Zone
	r.source.Find(&zones)
	for _, item := range results {
		for _, innerItem := range zones {
			if innerItem.BossID == item.ID {
				var zone datamodels.Zone
				r.source.First(&zone, innerItem.ID)
				r.source.Model(&item).Association("Zones").Append([]datamodels.Zone{zone})
			}
		}
		response = append(response, item)
	}

	return response
}

func (r *bossMemoryRepository) Select(id uint) (boss datamodels.Boss, found bool) {
	r.source.First(&boss, id)
	if boss.ID != 0 {
		var zones []datamodels.Zone
		r.source.Find(&zones)
		for _, item := range zones {
			if item.BossID == id {
				var zone datamodels.Zone
				r.source.First(&zone, item.ID)
				r.source.Model(&boss).Association("Zones").Append([]datamodels.Zone{zone})
			}
		}
		return boss, true
	}

	return
}

func (r *bossMemoryRepository) Insert(boss datamodels.Boss) (datamodels.Boss, error) {

	r.source.Create(&datamodels.Boss{
		Name:    boss.Name,
		Faction: boss.Faction,
		Zones:   boss.Zones,
		Photo:   boss.Photo,
	})
	r.source.Last(&boss)

	return boss, nil
}

func (r *bossMemoryRepository) Update(id uint, boss datamodels.Boss) (datamodels.Boss, bool, error) {
	var current datamodels.Boss
	r.source.First(&current, id)
	if current.ID != 0 {
		r.source.Model(&current).Updates(datamodels.Boss{
			Name:    boss.Name,
			Faction: boss.Faction,
			Zones:   boss.Zones,
			Photo:   boss.Photo,
		})
		return current, true, nil
	}

	return datamodels.Boss{}, false, errors.New("Not found record")
}

func (r *bossMemoryRepository) Delete(id uint) (boss datamodels.Boss, found bool) {
	r.source.First(&boss, id)
	if boss.ID != 0 {
		r.source.Delete(&boss)
		return boss, true
	}

	return
}
