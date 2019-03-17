package repositories

import (
	"errors"
	"yorha-api/datamodels"

	"github.com/jinzhu/gorm"
)

// AutomataRepository handles the basic operations of an entity/model.
type AutomataRepository interface {
	SelectMany() (results []datamodels.Automata)
	Select(id uint) (automata datamodels.Automata, found bool)
	Insert(automata datamodels.Automata) (insertedRecord datamodels.Automata, err error)
	Update(id uint, automata datamodels.Automata) (updatedRecord datamodels.Automata, found bool, err error)
}

// NewAutomataRepository returns a new memory-based repository,
func NewAutomataRepository(source *gorm.DB) AutomataRepository {
	return &automataMemoryRepository{source: source}
}

type automataMemoryRepository struct {
	source *gorm.DB
}

func (r *automataMemoryRepository) SelectMany() (results []datamodels.Automata) {
	r.source.Find(&results)

	return
}

func (r *automataMemoryRepository) Select(id uint) (automata datamodels.Automata, found bool) {
	var automatas []datamodels.Automata
	r.source.Find(&automatas)
	for _, item := range automatas {
		if item.ID == id {
			automata = item
			return automata, true
		}
	}

	return
}

func (r *automataMemoryRepository) Insert(automata datamodels.Automata) (datamodels.Automata, error) {
	r.source.Create(&datamodels.Automata{
		Name:       automata.Name,
		Occupation: automata.Occupation,
		Race:       automata.Race,
		Photo:      automata.Photo,
	})
	r.source.Last(&automata)

	return automata, nil
}

func (r *automataMemoryRepository) Update(id uint, automata datamodels.Automata) (datamodels.Automata, bool, error) {
	var automatas []datamodels.Automata
	var current datamodels.Automata
	r.source.Find(&automatas)
	for _, item := range automatas {
		if item.ID == id {
			current = item
			r.source.Model(&current).Updates(datamodels.Automata{
				Name:       automata.Name,
				Occupation: automata.Occupation,
				Race:       automata.Race,
				Photo:      automata.Photo,
			})
			return current, true, nil
		}
	}

	return datamodels.Automata{}, false, errors.New("Not found record")
}
