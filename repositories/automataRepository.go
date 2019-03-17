package repositories

import (
	"yorha-api/datamodels"
)

// AutomataRepository handles the basic operations of Automata entity/model.
type AutomataRepository interface {
	SelectMany() (results []datamodels.Automata)
	Select(id uint) (automata datamodels.Automata, found bool)
}

// NewAutomataRepository returns a new movie memory-based repository,
func NewAutomataRepository(source []datamodels.Automata) AutomataRepository {
	return &automataMemoryRepository{source: source}
}

type automataMemoryRepository struct {
	source []datamodels.Automata
}

const (
	// ReadOnlyMode will Lock(read) the data.
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

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
