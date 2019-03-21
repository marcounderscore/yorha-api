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

	return
}

func (r *bossMemoryRepository) Select(id uint) (boss datamodels.Boss, found bool) {
	r.source.First(&boss, id)
	if boss.ID != 0 {
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
