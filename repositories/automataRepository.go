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
	Delete(id uint) (automata datamodels.Automata, found bool)
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

	var response []datamodels.Automata
	for _, item := range results {
		var race datamodels.Race
		r.source.First(&race)
		r.source.Model(&item).Association("Race").Append(race)
		response = append(response, item)
	}

	return response
}

func (r *automataMemoryRepository) Select(id uint) (automata datamodels.Automata, found bool) {
	r.source.First(&automata, id)
	if automata.ID != 0 {
		var race datamodels.Race
		r.source.First(&race)
		r.source.Model(&automata).Association("Race").Append(race)
		return automata, true
	}

	return
}

func (r *automataMemoryRepository) Insert(automata datamodels.Automata) (datamodels.Automata, error) {
	var race datamodels.Race
	r.source.First(&race, automata.RaceID)
	if race.ID == 0 {
		return datamodels.Automata{}, errors.New("Foreign key constraint fail for Race id")
	}

	r.source.Create(&datamodels.Automata{
		Name:       automata.Name,
		Occupation: automata.Occupation,
		Photo:      automata.Photo,
		RaceID:     automata.RaceID,
	})
	r.source.Last(&automata)

	return automata, nil
}

func (r *automataMemoryRepository) Update(id uint, automata datamodels.Automata) (datamodels.Automata, bool, error) {
	var race datamodels.Race
	r.source.First(&race, automata.RaceID)
	if race.ID == 0 {
		return datamodels.Automata{}, true, errors.New("Foreign key constraint fail for Race id")
	}

	var current datamodels.Automata
	r.source.First(&current, id)
	if current.ID != 0 {
		r.source.Model(&current).Updates(datamodels.Automata{
			Name:       automata.Name,
			Occupation: automata.Occupation,
			RaceID:     automata.RaceID,
			Photo:      automata.Photo,
		})
		return current, true, nil
	}

	return datamodels.Automata{}, false, errors.New("Not found record")
}

func (r *automataMemoryRepository) Delete(id uint) (automata datamodels.Automata, found bool) {
	var automatas []datamodels.Automata
	r.source.Find(&automatas)
	for _, item := range automatas {
		if item.ID == id {
			automata = item
			r.source.Delete(&item)
			return automata, true
		}
	}

	return
}
