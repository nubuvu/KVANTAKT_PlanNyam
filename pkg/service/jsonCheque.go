package service

import (
	"KVANTAKT_PlanNyam/pkg/repository"
)

type JsonChequeService struct {
	repo repository.JsonCheque
}

func NewJsonChequeService(repo repository.JsonCheque) *JsonChequeService {
	return &JsonChequeService{repo: repo}
}

func (s *JsonChequeService) ParseJsonCheque(userId int, item []byte) ([]int, error) {
	return s.repo.ParseJsonCheque(userId, item)
}
