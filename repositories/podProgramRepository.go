package repositories

import (
	"errors"
	"yorha-api/datamodels"

	"github.com/jinzhu/gorm"
)

// PodProgramRepository handles the basic operations of an entity/model.
type PodProgramRepository interface {
	SelectMany() (results []datamodels.PodProgram)
	Select(id uint) (podProgram datamodels.PodProgram, found bool)
	Insert(podProgram datamodels.PodProgram) (insertedRecord datamodels.PodProgram, err error)
	Update(id uint, podProgram datamodels.PodProgram) (updatedRecord datamodels.PodProgram, found bool, err error)
	Delete(id uint) (podProgram datamodels.PodProgram, found bool)
}

// NewPodProgramRepository returns a new memory-based repository,
func NewPodProgramRepository(source *gorm.DB) PodProgramRepository {
	return &podProgramMemoryRepository{source: source}
}

type podProgramMemoryRepository struct {
	source *gorm.DB
}

func (r *podProgramMemoryRepository) SelectMany() (results []datamodels.PodProgram) {
	r.source.Find(&results)

	return
}

func (r *podProgramMemoryRepository) Select(id uint) (podProgram datamodels.PodProgram, found bool) {
	r.source.First(&podProgram, id)
	if podProgram.ID != 0 {
		return podProgram, true
	}

	return
}

func (r *podProgramMemoryRepository) Insert(podProgram datamodels.PodProgram) (datamodels.PodProgram, error) {

	r.source.Create(&datamodels.PodProgram{
		Name:     podProgram.Name,
		Program:  podProgram.Program,
		Cooldown: podProgram.Cooldown,
		Photo:    podProgram.Photo,
	})
	r.source.Last(&podProgram)

	return podProgram, nil
}

func (r *podProgramMemoryRepository) Update(id uint, podProgram datamodels.PodProgram) (datamodels.PodProgram, bool, error) {
	var current datamodels.PodProgram
	r.source.First(&current, id)
	if current.ID != 0 {
		r.source.Model(&current).Updates(datamodels.PodProgram{
			Name:     podProgram.Name,
			Program:  podProgram.Program,
			Cooldown: podProgram.Cooldown,
			Photo:    podProgram.Photo,
		})
		return current, true, nil
	}

	return datamodels.PodProgram{}, false, errors.New("Not found record")
}

func (r *podProgramMemoryRepository) Delete(id uint) (podProgram datamodels.PodProgram, found bool) {
	r.source.First(&podProgram, id)
	if podProgram.ID != 0 {
		r.source.Delete(&podProgram)
		return podProgram, true
	}

	return
}
