package service

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type CourtService interface {
	GetByID(uint) (*model.Court, error)
	GetAll() (model.Courts, error)
	Create(*model.Court) (model.Court, error)
	Update(*model.Court) (model.Court, error)
	Delete(uint) error
}

type courtService struct {
	courtRepository repository.CourtRepository
}

func NewCourtService(courtRepository repository.CourtRepository) CourtService {
	return &courtService{courtRepository: courtRepository}
}

func (s *courtService) GetByID(ID uint) (*model.Court, error) {
	court, err := s.courtRepository.GetByID(ID)
	if err != nil {
		return nil, err
	}
	return court, nil
}

func (s *courtService) GetAll() (model.Courts, error) {
	courts, err := s.courtRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return courts, nil
}

func (s *courtService) Create(court *model.Court) (model.Court, error) {
	if err := s.courtRepository.Create(court); err != nil {
		return model.Court{}, err
	}
	return *court, nil
}

func (s *courtService) Update(court *model.Court) (model.Court, error) {
	if err := s.courtRepository.Update(court); err != nil {
		return model.Court{}, err
	}
	return *court, nil
}

func (s *courtService) Delete(ID uint) error {
	if err := s.courtRepository.Delete(ID); err != nil {
		return err
	}
	return nil
}
