package services

import (
	"yorha-api/datamodels"
	"yorha-api/repositories"
)

// PodProgramService handles some of the CRUID operations of datamodel
type PodProgramService interface {
	GetAll() []datamodels.PodProgram
	GetByID(id uint) (datamodels.PodProgram, bool)
	Create(name string, program string, cooldown int, photo string) (datamodels.PodProgram, error)
	Update(id uint, name string, program string, cooldown int, photo string) (datamodels.PodProgram, bool, error)
	DeleteByID(id uint) (datamodels.PodProgram, bool)
}

// NewPodProgramService returns the default service
func NewPodProgramService(repo repositories.PodProgramRepository) PodProgramService {
	return &podProgramService{
		repo: repo,
	}
}

type podProgramService struct {
	repo repositories.PodProgramRepository
}

func (s *podProgramService) GetAll() []datamodels.PodProgram {
	return s.repo.SelectMany()
}

func (s *podProgramService) GetByID(id uint) (datamodels.PodProgram, bool) {
	return s.repo.Select(id)
}

func (s *podProgramService) Create(name string, program string, cooldown int, photo string) (datamodels.PodProgram, error) {
	return s.repo.Insert(datamodels.PodProgram{
		Name:     name,
		Program:  program,
		Cooldown: cooldown,
		Photo:    photo,
	})
}

func (s *podProgramService) Update(id uint, name string, program string, cooldown int, photo string) (datamodels.PodProgram, bool, error) {
	return s.repo.Update(id, datamodels.PodProgram{
		Name:     name,
		Program:  program,
		Cooldown: cooldown,
		Photo:    photo,
	})
}

func (s *podProgramService) DeleteByID(id uint) (datamodels.PodProgram, bool) {
	return s.repo.Delete(id)
}
