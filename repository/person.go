package repository

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"gorm.io/gorm"
)

type PersonRepository interface {
	GetByID(ID uint) (*model.Person, error)
	GetAll() (model.Persons, error)
	Create(person *model.Person) error
	Update(person *model.Person) error
	Delete(ID uint) error
}

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) PersonRepository {
	return &personRepository{db}
}

func (r *personRepository) GetByID(ID uint) (*model.Person, error) {
	var person model.Person
	if err := r.db.First(&person, ID).Error; err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *personRepository) GetAll() (model.Persons, error) {
	var persons model.Persons
	err := r.db.Find(&persons).Error
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *personRepository) Create(person *model.Person) error {
	if err := r.db.Create(person).Error; err != nil {
		return err
	}

	return nil
}

func (r *personRepository) Update(person *model.Person) error {
	if err := r.db.Save(person).Error; err != nil {
		return err
	}

	return nil
}

func (r *personRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.Person{}, ID).Error; err != nil {
		return err
	}
	return nil
}
