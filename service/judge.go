package service

import (
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
)

type JudgeService interface {
	GetByID(uint) (*model.Judge, error)
	GetAll() (model.Judges, error)
	Create(*model.Judge) (model.Judge, error)
	Update(*model.Judge) (model.Judge, error)
	Delete(uint) error
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
		return nil, err
	}
	return judge, nil
}

func (s *judgeService) GetAll() (model.Judges, error) {
	judges, err := s.judgeRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return judges, nil
}

func (s *judgeService) Create(judge *model.Judge) (model.Judge, error) {
	if err := s.judgeRepository.Create(judge); err != nil {
		return model.Judge{}, err
	}
	return *judge, nil
}

func (s *judgeService) Update(judge *model.Judge) (model.Judge, error) {
	if err := s.judgeRepository.Update(judge); err != nil {
		return model.Judge{}, err
	}
	return *judge, nil
}

func (s *judgeService) Delete(ID uint) error {
	if err := s.judgeRepository.Delete(ID); err != nil {
		return err
	}
	return nil
}
