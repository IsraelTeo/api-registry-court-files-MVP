package service

import (
	"errors"

	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type JudgeService interface {
	GetByID(ID uint) (*model.Judge, error)
	GetAll() ([]model.Judge, error)
	Create(judge *model.Judge) (model.Judge, error)
	Update(ID uint, judge *model.Judge) (model.Judge, error)
	Delete(ID uint) error
}

type judgeService struct {
	judgeRepository repository.JudgeRepository
}

func NewJudgeService(judgeRepository repository.JudgeRepository) JudgeService {
	return &judgeService{judgeRepository: judgeRepository}
}

func (s *judgeService) GetByID(ID uint) (*model.Judge, error) {
	judge, err := s.judgeRepository.GetByID(ID)
	if err != nil {
		return nil, errors.New("error fetching judge by ID")
	}

	return judge, nil
}

func (s *judgeService) GetAll() ([]model.Judge, error) {
	judges, err := s.judgeRepository.GetAll()
	if err != nil {
		return nil, errors.New("error fetching all judges")
	}

	return judges, nil
}

func (s *judgeService) Create(judge *model.Judge) (model.Judge, error) {
	if err := s.judgeRepository.Create(judge); err != nil {
		return model.Judge{}, errors.New("error creating judge")
	}

	return *judge, nil
}

func (s *judgeService) Update(ID uint, judge *model.Judge) (model.Judge, error) {
	if judge.ID != ID {
		return model.Judge{}, errors.New("judge ID mismatch")
	}

	judgeFound, err := s.judgeRepository.GetByID(ID)
	if err != nil {
		return model.Judge{}, errors.New("judge not found to update")
	}

	judge.FullName = judgeFound.FullName
	judge.Specialty = judgeFound.Specialty
	judge.CourtID = judgeFound.CourtID

	if err := s.judgeRepository.Update(judge); err != nil {
		return model.Judge{}, errors.New("error updating judge")
	}

	return *judge, nil
}

func (s *judgeService) Delete(ID uint) error {
	if err := s.judgeRepository.Delete(ID); err != nil {
		return errors.New("error deleting judge")
	}

	return nil
}
