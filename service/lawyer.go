package service

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type LawyerService interface {
	GetByID(uint) (*model.Lawyer, error)
	GetAll() (model.Lawyers, error)
	Create(*model.Lawyer) (model.Lawyer, error)
	Update(*model.Lawyer) (model.Lawyer, error)
	Delete(uint) error
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
		return nil, err
	}
	return lawyer, nil
}

func (s *lawyerService) GetAll() (model.Lawyers, error) {
	lawyers, err := s.lawyerRepository.GetAll()
	if err != nil {
		return nil, err
	}
	retur
	n lawyers, nil
}

func (s *lawyerService) Create(lawyer *model.Lawyer) (model.Lawyer, error) {
	if err := s.lawyerRepository.Create(lawyer); err != nil {
		return model.Lawyer{}, err
	}

	return *lawyer, nil
}

func (s *lawyerService) Update(lawyer *model.Lawyer) (model.Lawyer, error) {
	if err := s.lawyerRepository.Update(lawyer); err != nil {
		return model.Lawyer{}, err
	}

	return *lawyer, nil
}

func (s *lawyerService) Delete(ID uint) error {
	if err := s.lawyerRepository.Delete(ID); err != nil {
		return err
	}

	return nil
}
