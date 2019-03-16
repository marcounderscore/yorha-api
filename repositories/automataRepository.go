package repositories

import (
	"yorha-api/datamodels"
)

// AutomataRepository handles the basic operations of Automata entity/model.
type AutomataRepository interface {
	SelectMany() (results []datamodels.Automata)
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
