package eth

import (
	"net/http"
)

// type BidResponse struct {
// 	FromAddress string `json:"fromAddress"`
// 	TicketNum   int    `json:"ticketNum"`
// }

func (e *EthereumServer) BidHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("xBidHandler()")
	// bidResponse := &BidResponse{}
	// err := json.NewDecoder(r.Body).Decode(bidResponse)
	// if err != nil {
	// 	e.logger.Error(err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// for len(e.newEvents) == 0 {
	// }

	// lastEvent := e.newEvents[0]

	// fmt.Println("last event: ", lastEvent)

	// e.newEvents = e.newEvents[:]

	// e.logger.WithFields(logrus.Fields{"BidHandler()": bidResponse})
	// w.WriteHeader(http.StatusOK)
}
