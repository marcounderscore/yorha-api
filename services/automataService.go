package services

import (
	"yorha-api/datamodels"
	"yorha-api/repositories"
)

// AutomataService handles some of the CRUID operations of datamodel
type AutomataService interface {
	GetAll() []datamodels.Automata
	GetByID(id uint) (datamodels.Automata, bool)
}

// NewAutomataService returns the default service
func NewAutomataService(repo repositories.AutomataRepository) AutomataService {
	return &automataService{
		repo: repo,
	}
}

type automataService struct {
	repo repositories.AutomataRepository
}

// GetAll returns all
func (s *automataService) GetAll() []datamodels.Automata {
	return s.repo.SelectMany()
}

// GetByID returns one based on its id
func (s *automataService) GetByID(id uint) (datamodels.Automata, bool) {
	return s.repo.Select(id)
}
