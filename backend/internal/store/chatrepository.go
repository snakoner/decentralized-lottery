package store

import "github.com/snakoner/lottery/internal/models"

type ChatRepository interface {
	Create(*models.Message) error
	GetMessages(int) ([]*models.Message, error)
}
