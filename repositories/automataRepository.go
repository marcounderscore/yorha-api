package repositories

import (
	"yorha-api/datamodels"
)

// AutomataRepository handles the basic operations of an entity/model.
type AutomataRepository interface {
	SelectMany() (results []datamodels.Automata)
	Select(id uint) (automata datamodels.Automata, found bool)
}

// NewAutomataRepository returns a new memory-based repository,
func NewAutomataRepository(source []datamodels.Automata) AutomataRepository {
	return &automataMemoryRepository{source: source}
}

type automataMemoryRepository struct {
	source []datamodels.Automata
}

func (r *automataMemoryRepository) SelectMany() (results []datamodels.Automata) {
	results = r.source

	return
}

func (r *automataMemoryRepository) Select(id uint) (automata datamodels.Automata, found bool) {
	for _, item := range r.source {
		if item.ID == id {
			automata = item
			return automata, true
		}
	}

	return
}
