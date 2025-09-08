package repository

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"gorm.io/gorm"
)

type JudicialFileRepository interface {
	GetByID(ID uint) (*model.JudicialFile, error)
	GetAll() (model.JudicialFiles, error)
	Create(file *model.JudicialFile) error
	Update(file *model.JudicialFile) error
	Delete(ID uint) error
	AddPerson(judicialFileID, personID uint) error
	AddLawyer(judicialFileID, lawyerID uint) error
}

type judicialFileRepository struct {
	db *gorm.DB
}

func NewJudicialFileRepository(db *gorm.DB) JudicialFileRepository {
	return &judicialFileRepository{db}
}

func (r *judicialFileRepository) GetByID(ID uint) (*model.JudicialFile, error) {
	var file model.JudicialFile
	err := r.db.Preload("Persons").Preload("Lawyers").First(&file, ID).Error
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *judicialFileRepository) GetAll() (model.JudicialFiles, error) {
	var files model.JudicialFiles
	if err := r.db.Preload("Persons").Preload("Lawyers").Find(&files).Error; err != nil {
		return nil, err
	}

	return files, nil
}

func (r *judicialFileRepository) Create(file *model.JudicialFile) error {
	if err := r.db.Create(file).Error; err != nil {
		return err
	}

	return nil
}

func (r *judicialFileRepository) Update(file *model.JudicialFile) error {
	if err := r.db.Save(file).Error; err != nil {
		return err
	}

	return nil
}

func (r *judicialFileRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.JudicialFile{}, ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *judicialFileRepository) AddPerson(judicialFileID, personID uint) error {
	var file model.JudicialFile
	if err := r.db.First(&file, judicialFileID).Error; err != nil {
		return err
	}

	var person model.Person
	if err := r.db.First(&person, personID).Error; err != nil {
		return err
	}

	return r.db.Model(&file).Association("Persons").Append(&person)
}

func (r *judicialFileRepository) AddLawyer(judicialFileID, lawyerID uint) error {
	var file model.JudicialFile
	if err := r.db.First(&file, judicialFileID).Error; err != nil {
		return err
	}

	var lawyer model.Lawyer
	if err := r.db.First(&lawyer, lawyerID).Error; err != nil {
		return err
	}

	return r.db.Model(&file).Association("Lawyers").Append(&lawyer)
}
