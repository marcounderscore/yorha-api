package services

import (
	"yorha-api/datamodels"
	"yorha-api/repositories"
)

// BossService handles some of the CRUID operations of datamodel
type BossService interface {
	GetAll() []datamodels.Boss
	GetByID(id uint) (datamodels.Boss, bool)
	Create(name string, faction string, zones []datamodels.Zone, photo string) (datamodels.Boss, error)
	Update(id uint, name string, faction string, zones []datamodels.Zone, photo string) (datamodels.Boss, bool, error)
	DeleteByID(id uint) (datamodels.Boss, bool)
}

// NewBossService returns the default service
func NewBossService(repo repositories.BossRepository) BossService {
	return &bossService{
		repo: repo,
	}
}

type bossService struct {
	repo repositories.BossRepository
}

func (s *bossService) GetAll() []datamodels.Boss {
	return s.repo.SelectMany()
}

func (s *bossService) GetByID(id uint) (datamodels.Boss, bool) {
	return s.repo.Select(id)
}

func (s *bossService) Create(name string, faction string, zones []datamodels.Zone, photo string) (datamodels.Boss, error) {
	return s.repo.Insert(datamodels.Boss{
		Name:    name,
		Faction: faction,
		Zones:   zones,
		Photo:   photo,
	})
}

func (s *bossService) Update(id uint, name string, faction string, zones []datamodels.Zone, photo string) (datamodels.Boss, bool, error) {
	return s.repo.Update(id, datamodels.Boss{
		Name:    name,
		Faction: faction,
		Zones:   zones,
		Photo:   photo,
	})
}

func (s *bossService) DeleteByID(id uint) (datamodels.Boss, bool) {
	return s.repo.Delete(id)
}
