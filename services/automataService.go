package services

import (
	"yorha-api/datamodels"
	"yorha-api/repositories"
)

// AutomataService handles some of the CRUID operations of datamodel
type AutomataService interface {
	GetAll() []datamodels.Automata
	GetByID(id uint) (datamodels.Automata, bool)
	Create(name string, occupation string, race string, photo string) (datamodels.Automata, error)
	Update(id uint, name string, occupation string, race string, photo string) (datamodels.Automata, bool, error)
	DeleteByID(id uint) (datamodels.Automata, bool)
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

func (s *automataService) GetAll() []datamodels.Automata {
	return s.repo.SelectMany()
}

func (s *automataService) GetByID(id uint) (datamodels.Automata, bool) {
	return s.repo.Select(id)
}

func (s *automataService) Create(name string, occupation string, race string, photo string) (datamodels.Automata, error) {
	return s.repo.Insert(datamodels.Automata{
		Name:       name,
		Occupation: occupation,
		Race:       race,
		Photo:      photo,
	})
}

func (s *automataService) Update(id uint, name string, occupation string, race string, photo string) (datamodels.Automata, bool, error) {
	return s.repo.Update(id, datamodels.Automata{
		Name:       name,
		Occupation: occupation,
		Race:       race,
		Photo:      photo,
	})
}

func (s *automataService) DeleteByID(id uint) (datamodels.Automata, bool) {
	return s.repo.Delete(id)
}
