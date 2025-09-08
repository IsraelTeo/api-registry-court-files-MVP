package service

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type PersonService interface {
	GetByID(uint) (*model.Person, error)
	GetAll() (model.Persons, error)
	Create(*model.Person) (model.Person, error)
	Update(*model.Person) (model.Person, error)
	Delete(uint) error
}

type personService struct {
	personRepository repository.PersonRepository
}

func NewPersonService(personRepository repository.PersonRepository) PersonService {
	return &personService{personRepository: personRepository}
}

func (s *personService) GetByID(ID uint) (*model.Person, error) {
	person, err := s.personRepository.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (s *personService) GetAll() (model.Persons, error) {
	persons, err := s.personRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (s *personService) Create(person *model.Person) (model.Person, error) {
	if err := s.personRepository.Create(person); err != nil {
		return model.Person{}, err
	}

	return *person, nil
}

func (s *personService) Update(person *model.Person) (model.Person, error) {
	if err := s.personRepository.Update(person); err != nil {
		return model.Person{}, err
	}

	return *person, nil
}

func (s *personService) Delete(ID uint) error {
	if err := s.personRepository.Delete(ID); err != nil {
		return err
	}

	return nil
}
