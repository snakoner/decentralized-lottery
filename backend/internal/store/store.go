package store

type Store interface {
	Chat() ChatRepository
}
