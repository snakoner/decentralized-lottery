package models

type Message struct {
	ID      int    `json:"id" db:"id"`
	Date    int64  `json:"date" db:"ts"`
	User    string `json:"user" db:"user_addr"`
	Content string `json:"content" db:"content"`
}
