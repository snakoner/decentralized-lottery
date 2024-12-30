package sqlstore

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/snakoner/lottery/internal/store"
)

// Store ...
type Store struct {
	dbx            *sql.DB
	url            string
	chatRepository *ChatRepository
}

func New(url string) *Store {
	return &Store{
		url: url,
	}
}

func (s *Store) Connect(url string) error {
	dbx, err := sql.Open("postgres", url)
	if err != nil {
		s.dbx = nil

		return err
	}

	s.dbx = dbx

	return nil
}

func (s *Store) Close() error {
	if s.dbx != nil {
		return s.dbx.Close()
	}

	return errors.New("s.dbx is nil")
}

// User ...
func (s *Store) Chat() store.ChatRepository {
	if s.chatRepository != nil {
		return s.chatRepository
	}

	s.chatRepository = &ChatRepository{
		store: s,
	}

	return s.chatRepository
}
