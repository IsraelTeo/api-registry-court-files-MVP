package service

import (
	"errors"

	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type PersonService interface {
	GetByID(ID uint) (*model.Person, error)
	GetAll() ([]model.Person, error)
	Create(person *model.Person) (model.Person, error)
	Update(ID uint, person *model.Person) (model.Person, error)
	Delete(ID uint) error
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
		return nil, errors.New("error fetching person by ID")
	}

	return person, nil
}

func (s *personService) GetAll() ([]model.Person, error) {
	persons, err := s.personRepository.GetAll()
	if err != nil {
		return nil, errors.New("error fetching all persons")
	}

	return persons, nil
}

func (s *personService) Create(person *model.Person) (model.Person, error) {
	if err := s.personRepository.Create(person); err != nil {
		return model.Person{}, errors.New("error creating person")
	}

	return *person, nil
}

func (s *personService) Update(ID uint, person *model.Person) (model.Person, error) {
	if person.ID != ID {
		return model.Person{}, errors.New("person ID mismatch")
	}

	personFound, err := s.personRepository.GetByID(ID)
	if err != nil {
		return model.Person{}, errors.New("person not found to update")
	}

	personFound.FullName = person.FullName
	personFound.Role = person.Role

	if err := s.personRepository.Update(person); err != nil {
		return model.Person{}, errors.New("error updating person")
	}

	return *person, nil
}

func (s *personService) Delete(ID uint) error {
	if err := s.personRepository.Delete(ID); err != nil {
		return errors.New("error deleting person")
	}

	return nil
}
