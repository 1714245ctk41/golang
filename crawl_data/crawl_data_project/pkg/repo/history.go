package repo

import (
	"context"
	"crawl_data/pkg/model"
	"crawl_data/pkg/utils"
	"log"
	"time"
)

type HistoryRepository interface {
	InsertHistory(history model.History) error
}

type historyConnection struct{}

func NewHistoryRepository() HistoryRepository {
	return &historyConnection{}
}

var collectionHistory = "history"

func (vas *historyConnection) InsertHistory(history model.History) error {
	// user.CreatedAt = time.Now().Format(time.RFC1123Z)
	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	collection := client.Database(mongoDBNam).Collection(collectionHistory)

	history.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), history)
	if err != nil {
		log.Printf("Could not insert history: %v", err)
		return err
	}

	return nil
}
