package repositories

import (
	"yorha-api/datamodels"

	"github.com/jinzhu/gorm"
)

// AutomataRepository handles the basic operations of an entity/model.
type AutomataRepository interface {
	SelectMany() (results []datamodels.Automata)
	Select(id uint) (automata datamodels.Automata, found bool)
}

// NewAutomataRepository returns a new memory-based repository,
func NewAutomataRepository(source *gorm.DB) AutomataRepository {
	return &automataMemoryRepository{source: source}
}

type automataMemoryRepository struct {
	source *gorm.DB
}

func (r *automataMemoryRepository) SelectMany() (results []datamodels.Automata) {
	//Gorm method to retrieve all data from this source
	r.source.Find(&results)

	return
}

func (r *automataMemoryRepository) Select(id uint) (automata datamodels.Automata, found bool) {
	//Due unexpected result with "First" method the "Find" method with "for" loop is used instead
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
