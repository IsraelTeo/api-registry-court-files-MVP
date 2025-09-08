package service

import (
	"errors"
	"time"

	"github.com/IsraelTeo/api-registry-court-files-MVP/dto"
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type JudicialFileService interface {
	GetByID(ID uint) (*model.JudicialFile, error)
	GetAll() ([]model.JudicialFile, error)
	Create(filDTO *dto.JudicialFile) (model.JudicialFile, error)
	Update(ID uint, file *model.JudicialFile) (model.JudicialFile, error)
	Delete(ID uint) error
	AddPerson(judicialFileID, personID uint) error
	AddLawyer(judicialFileID, lawyerID uint) error
}

type judicialFileService struct {
	judicialFileRepository repository.JudicialFileRepository
	personRepository       repository.PersonRepository
	lawyerRepository       repository.LawyerRepository
	courtRepository        repository.CourtRepository
}

func NewJudicialFileService(
	judicialFileRepository repository.JudicialFileRepository,
	personRepository repository.PersonRepository,
	lawyerRepository repository.LawyerRepository,
	courtRepository repository.CourtRepository,
) JudicialFileService {
	return &judicialFileService{
		judicialFileRepository: judicialFileRepository,
		personRepository:       personRepository,
		lawyerRepository:       lawyerRepository,
		courtRepository:        courtRepository,
	}
}

func (s *judicialFileService) GetByID(ID uint) (*model.JudicialFile, error) {
	file, err := s.judicialFileRepository.GetByID(ID)
	if err != nil {
		return nil, errors.New("error fetching judicial file by ID")
	}

	return file, nil
}

func (s *judicialFileService) GetAll() ([]model.JudicialFile, error) {
	files, err := s.judicialFileRepository.GetAll()
	if err != nil {
		return nil, errors.New("error fetching all judicial files")
	}

	return files, nil
}

func (s *judicialFileService) Create(fileDTO *dto.JudicialFile) (model.JudicialFile, error) {
	courtFound, err := s.courtRepository.GetByID(fileDTO.CourtID)
	if err != nil {
		return model.JudicialFile{}, errors.New("court not found to create judicial file")
	}

	persons := []model.Person{}
	for _, personID := range fileDTO.PersonsIDs {
		person, err := s.personRepository.GetByID(personID)
		if err != nil {
			return model.JudicialFile{}, errors.New("person not found to create judicial file")
		}
		persons = append(persons, *person)
	}

	lawyers := []model.Lawyer{}
	for _, lawyerID := range fileDTO.LawyersIDs {
		lawyer, err := s.lawyerRepository.GetByID(lawyerID)
		if err != nil {
			return model.JudicialFile{}, errors.New("lawyer not found to create judicial file")
		}
		lawyers = append(lawyers, *lawyer)
	}

	file := model.JudicialFile{
		FileNumber:         fileDTO.FileNumber,
		NotificationNumber: fileDTO.NotificationNumber,
		DigitizationNumber: fileDTO.DigitizationNumber,
		DocumentType:       fileDTO.DocumentType,
		Headquarters:       courtFound.Headquarters,
		Persons:            persons,
		Lawyers:            lawyers,
		Court:              courtFound.Name,
		NotificationDate:   fileDTO.NotificationDate,
		CreationDate:       time.Now(),
		CourtID:            fileDTO.CourtID,
	}

	if err := s.judicialFileRepository.Create(&file); err != nil {
		return model.JudicialFile{}, errors.New("error creating judicial file")
	}

	return file, nil
}

func (s *judicialFileService) Update(ID uint, file *model.JudicialFile) (model.JudicialFile, error) {
	if file.ID != ID {
		return model.JudicialFile{}, errors.New("judicial file ID mismatch")
	}

	fileFound, err := s.judicialFileRepository.GetByID(ID)
	if err != nil {
		return model.JudicialFile{}, errors.New("error fetching judicial file to update")
	}

	file.FileNumber = fileFound.FileNumber
	file.NotificationNumber = fileFound.NotificationNumber
	file.DigitizationNumber = fileFound.DigitizationNumber
	file.DocumentType = fileFound.DocumentType
	file.Headquarters = fileFound.Headquarters
	file.Court = fileFound.Court
	file.NotificationDate = fileFound.NotificationDate
	file.UpdateDate = time.Now()
	file.CourtID = fileFound.CourtID
	file.Persons = fileFound.Persons
	file.Lawyers = fileFound.Lawyers

	if err := s.judicialFileRepository.Update(file); err != nil {
		return model.JudicialFile{}, errors.New("error updating judicial file")
	}

	return *file, nil
}

func (s *judicialFileService) Delete(ID uint) error {
	if err := s.judicialFileRepository.Delete(ID); err != nil {
		return errors.New("error deleting judicial file")
	}

	return nil
}

func (s *judicialFileService) AddPerson(judicialFileID, personID uint) error {
	if err := s.judicialFileRepository.AddPerson(judicialFileID, personID); err != nil {
		return errors.New("error adding person to judicial file")
	}

	return nil
}

func (s *judicialFileService) AddLawyer(judicialFileID, lawyerID uint) error {
	if err := s.judicialFileRepository.AddLawyer(judicialFileID, lawyerID); err != nil {
		return errors.New("error adding lawyer to judicial file")
	}

	return nil
}
