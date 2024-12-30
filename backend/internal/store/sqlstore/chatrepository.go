package sqlstore

import (
	"fmt"

	"github.com/snakoner/lottery/internal/models"
)

type ChatRepository struct {
	store *Store
}

func (r *ChatRepository) Create(message *models.Message) error {
	if err := r.store.Connect(r.store.url); err != nil {
		return err
	}

	defer r.store.Close()

	query := `INSERT INTO chat(ts, user, content) 
		VALUES ($1, $2, $3)`

	_, err := r.store.dbx.Exec(query,
		message.Date,
		message.User,
		message.Content)

	return err
}

func (r *ChatRepository) GetMessages(limit int) ([]*models.Message, error) {
	if err := r.store.Connect(r.store.url); err != nil {
		return nil, err
	}

	defer r.store.Close()
	messages := []*models.Message{}
	var query string

	if limit != 0 {
		query = fmt.Sprintf(`SELECT * FROM chat ORDER BY id DESC LIMIT %d`, limit)
	} else {
		query = fmt.Sprintf(`SELECT * FROM chat`)
	}

	rows, err := r.store.dbx.Query(query)
	defer func() {
		if rows != nil {
			rows.Close()
		}
		r.store.Close()
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		message := models.Message{}
		err := rows.Scan(&message.ID,
			&message.Date,
			&message.User,
			&message.Content)

		if err != nil {
			return nil, err
		}

		messages = append(messages, &message)
	}

	return messages, nil
}
