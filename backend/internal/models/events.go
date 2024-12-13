package models

type BidEvent struct {
	Account   string `json:"account"`
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	Round     int64  `json:"round"`
}

type BidEvents struct {
	Events []*BidEvent `json:"events"`
}
