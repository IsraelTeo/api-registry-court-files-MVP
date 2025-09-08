package service

import (
	"errors"

	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type CourtService interface {
	GetByID(ID uint) (*model.Court, error)
	GetAll() ([]model.Court, error)
	Create(*model.Court) (model.Court, error)
	Update(ID uint, court *model.Court) (model.Court, error)
	Delete(ID uint) error
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
		return nil, errors.New("error fetching court by ID")
	}

	return court, nil
}

func (s *courtService) GetAll() ([]model.Court, error) {
	courts, err := s.courtRepository.GetAll()
	if err != nil {
		return nil, errors.New("error fetching all courts")
	}

	return courts, nil
}

func (s *courtService) Create(court *model.Court) (model.Court, error) {
	if err := s.courtRepository.Create(court); err != nil {
		return model.Court{}, errors.New("error creating court")
	}

	return *court, nil
}

func (s *courtService) Update(ID uint, court *model.Court) (model.Court, error) {
	if court.ID != ID {
		return model.Court{}, errors.New("court ID mismatch")
	}

	courtFound, err := s.courtRepository.GetByID(ID)
	if err != nil {
		return model.Court{}, errors.New("court not found to update")
	}

	court.Name = courtFound.Name
	court.Headquarters = courtFound.Headquarters
	court.Judges = courtFound.Judges

	if err := s.courtRepository.Update(court); err != nil {
		return model.Court{}, errors.New("error updating court")
	}

	return *court, nil
}

func (s *courtService) Delete(ID uint) error {
	if err := s.courtRepository.Delete(ID); err != nil {
		return errors.New("error deleting court")
	}

	return nil
}
