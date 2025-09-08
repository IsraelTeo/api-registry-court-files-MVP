package service

import (
	"errors"

	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type LawyerService interface {
	GetByID(ID uint) (*model.Lawyer, error)
	GetAll() ([]model.Lawyer, error)
	Create(lawyer *model.Lawyer) (model.Lawyer, error)
	Update(ID uint, lawyer *model.Lawyer) (model.Lawyer, error)
	Delete(ID uint) error
}

type lawyerService struct {
	lawyerRepository repository.LawyerRepository
}

func NewLawyerService(lawyerRepository repository.LawyerRepository) LawyerService {
	return &lawyerService{lawyerRepository: lawyerRepository}
}

func (s *lawyerService) GetByID(ID uint) (*model.Lawyer, error) {
	lawyer, err := s.lawyerRepository.GetByID(ID)
	if err != nil {
		return nil, errors.New("error fetching lawyer by ID")
	}

	return lawyer, nil
}

func (s *lawyerService) GetAll() ([]model.Lawyer, error) {
	lawyers, err := s.lawyerRepository.GetAll()
	if err != nil {
		return nil, errors.New("error fetching all lawyers")
	}

	return lawyers, nil
}

func (s *lawyerService) Create(lawyer *model.Lawyer) (model.Lawyer, error) {
	if err := s.lawyerRepository.Create(lawyer); err != nil {
		return model.Lawyer{}, errors.New("error creating lawyer")
	}

	return *lawyer, nil
}

func (s *lawyerService) Update(ID uint, lawyer *model.Lawyer) (model.Lawyer, error) {
	if lawyer.ID != ID {
		return model.Lawyer{}, errors.New("lawyer ID mismatch")
	}

	lawyerFound, err := s.lawyerRepository.GetByID(ID)
	if err != nil {
		return model.Lawyer{}, errors.New("lawyer not found to update")
	}

	lawyer.FullName = lawyerFound.FullName
	lawyer.BarNumber = lawyerFound.BarNumber

	if err := s.lawyerRepository.Update(lawyer); err != nil {
		return model.Lawyer{}, errors.New("error updating lawyer")
	}

	return *lawyer, nil
}

func (s *lawyerService) Delete(ID uint) error {
	if err := s.lawyerRepository.Delete(ID); err != nil {
		return errors.New("error deleting lawyer")
	}

	return nil
}
