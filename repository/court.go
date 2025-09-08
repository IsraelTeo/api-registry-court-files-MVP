package repository

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"gorm.io/gorm"
)

type CourtRepository interface {
	GetByID(ID uint) (*model.Court, error)
	GetAll() (model.Courts, error)
	Create(court *model.Court) error
	Update(court *model.Court) error
	Delete(ID uint) error
}

type courtRepository struct {
	db *gorm.DB
}

func NewCourtRepository(db *gorm.DB) CourtRepository {
	return &courtRepository{db}
}

func (r *courtRepository) GetByID(ID uint) (*model.Court, error) {
	var court model.Court
	if err := r.db.First(&court, ID).Error; err != nil {
		return nil, err
	}

	return &court, nil
}

func (r *courtRepository) GetAll() (model.Courts, error) {
	var courts model.Courts
	if err := r.db.Find(&courts).Error; err != nil {
		return nil, err
	}

	return courts, nil
}

func (r *courtRepository) Create(court *model.Court) error {
	if err := r.db.Create(court).Error; err != nil {
		return err
	}

	return nil
}

func (r *courtRepository) Update(court *model.Court) error {
	if err := r.db.Save(court).Error; err != nil {
		return err
	}

	return nil
}

func (r *courtRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.Court{}, ID).Error; err != nil {
		return err
	}

	return nil
}
