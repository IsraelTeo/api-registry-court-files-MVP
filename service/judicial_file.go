package service

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type JudicialFileService interface {
	GetByID(uint) (*model.JudicialFile, error)
	GetAll() (model.JudicialFiles, error)
	Create(*model.JudicialFile) (model.JudicialFile, error)
	Update(*model.JudicialFile) (model.JudicialFile, error)
	Delete(uint) error
	AddPerson(judicialFileID, personID uint) error
	AddLawyer(judicialFileID, lawyerID uint) error
}

type judicialFileService struct {
	judicialFileRepository repository.JudicialFileRepository
}

func NewJudicialFileService(judicialFileRepository repository.JudicialFileRepository) JudicialFileService {
	return &judicialFileService{judicialFileRepository: judicialFileRepository}
}

func (s *judicialFileService) GetByID(ID uint) (*model.JudicialFile, error) {
	file, err := s.judicialFileRepository.GetByID(ID)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *judicialFileService) GetAll() (model.JudicialFiles, error) {
	files, err := s.judicialFileRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (s *judicialFileService) Create(file *model.JudicialFile) (model.JudicialFile, error) {
	if err := s.judicialFileRepository.Create(file); err != nil {
		return model.JudicialFile{}, err
	}
	return *file, nil
}

func (s *judicialFileService) Update(file *model.JudicialFile) (model.JudicialFile, error) {
	if err := s.judicialFileRepository.Update(file); err != nil {
		return model.JudicialFile{}, err
	}
	return *file, nil
}

func (s *judicialFileService) Delete(ID uint) error {
	if err := s.judicialFileRepository.Delete(ID); err != nil {
		return err
	}
	return nil
}

func (s *judicialFileService) AddPerson(judicialFileID, personID uint) error {
	if err := s.judicialFileRepository.AddPerson(judicialFileID, personID); err != nil {
		return err
	}
	return nil
}

func (s *judicialFileService) AddLawyer(judicialFileID, lawyerID uint) error {
	if err := s.judicialFileRepository.AddLawyer(judicialFileID, lawyerID); err != nil {
		return err
	}
	return nil
}
