package service

import (
	"crawl_data/pkg/model"
	"crawl_data/pkg/repo"
)

type HistoryService interface {
	InsertHistory(history model.History) error
}

type historyService struct {
	historyRepository repo.HistoryRepository
}

func NewHistoryService(historyRepository repo.HistoryRepository) HistoryService {
	return &historyService{
		historyRepository: historyRepository,
	}
}

func (vas *historyService) InsertHistory(history model.History) error {
	err := vas.historyRepository.InsertHistory(history)
	if err != nil {
		return err
	}
	return nil
}
