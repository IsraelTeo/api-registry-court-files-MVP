package repository

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"gorm.io/gorm"
)

type JudgeRepository interface {
	GetByID(ID uint) (*model.Judge, error)
	GetAll() ([]model.Judge, error)
	Create(judge *model.Judge) error
	Update(judge *model.Judge) error
	Delete(ID uint) error
}

type judgeRepository struct {
	db *gorm.DB
}

func NewJudgeRepository(db *gorm.DB) JudgeRepository {
	return &judgeRepository{db}
}

func (r *judgeRepository) GetByID(ID uint) (*model.Judge, error) {
	var judge model.Judge
	if err := r.db.First(&judge, ID).Error; err != nil {
		return nil, err
	}

	return &judge, nil
}

func (r *judgeRepository) GetAll() ([]model.Judge, error) {
	var judges []model.Judge
	if err := r.db.Find(&judges).Error; err != nil {
		return nil, err
	}

	return judges, nil
}

func (r *judgeRepository) Create(judge *model.Judge) error {
	if err := r.db.Create(judge).Error; err != nil {
		return err
	}

	return nil
}

func (r *judgeRepository) Update(judge *model.Judge) error {
	if err := r.db.Save(judge).Error; err != nil {
		return err
	}

	return nil
}

func (r *judgeRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.Judge{}, ID).Error; err != nil {
		return err
	}

	return nil
}
