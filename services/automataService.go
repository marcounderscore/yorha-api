package services

import (
	"yorha-api/datamodels"
	"yorha-api/repositories"
)

// AutomataService handles some of the CRUID operations of the movie datamodel.
type AutomataService interface {
	GetAll() []datamodels.Automata
}

// NewAutomataService returns the default movie service.
func NewAutomataService(repo repositories.AutomataRepository) AutomataService {
	return &automataService{
		repo: repo,
	}
}

type automataService struct {
	repo repositories.AutomataRepository
}

// GetAll returns all movies.
func (s *automataService) GetAll() []datamodels.Automata {
	return s.repo.SelectMany()
}
