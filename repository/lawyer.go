package repository

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"gorm.io/gorm"
)

type LawyerRepository interface {
	GetByID(ID uint) (*model.Lawyer, error)
	GetAll() ([]model.Lawyer, error)
	Create(lawyer *model.Lawyer) error
	Update(lawyer *model.Lawyer) error
	Delete(ID uint) error
}

type lawyerRepository struct {
	db *gorm.DB
}

func NewLawyerRepository(db *gorm.DB) LawyerRepository {
	return &lawyerRepository{db}
}

func (r *lawyerRepository) GetByID(ID uint) (*model.Lawyer, error) {
	var lawyer model.Lawyer
	if err := r.db.First(&lawyer, ID).Error; err != nil {
		return nil, err
	}

	return &lawyer, nil
}

func (r *lawyerRepository) GetAll() ([]model.Lawyer, error) {
	var lawyers []model.Lawyer
	if err := r.db.Find(&lawyers).Error; err != nil {
		return nil, err
	}

	return lawyers, nil
}

func (r *lawyerRepository) Create(lawyer *model.Lawyer) error {
	if err := r.db.Create(lawyer).Error; err != nil {
		return err
	}

	return nil
}

func (r *lawyerRepository) Update(lawyer *model.Lawyer) error {
	if err := r.db.Save(lawyer).Error; err != nil {
		return err
	}

	return nil
}

func (r *lawyerRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.Lawyer{}, ID).Error; err != nil {
		return err
	}

	return nil
}
