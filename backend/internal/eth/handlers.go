package eth

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/mux"
	"github.com/snakoner/lottery/internal/models"
)

func setHttpError(w http.ResponseWriter, err string, code int) {
	http.Error(w, err, code)
}

func (e *EthereumServer) BidHandler(w http.ResponseWriter, r *http.Request) {

}

// todo: filter by indexed round
// RoundHandler is an HTTP handler that processes requests for retrieving bid events
// corresponding to a specific lottery round.
//
// Parameters:
//   - w: The HTTP response writer used to send responses back to the client.
//   - r: The HTTP request containing the parameters for the round.
func (e *EthereumServer) RoundHandler(w http.ResponseWriter, r *http.Request) {
	roundNum := mux.Vars(r)["number"]
	round, err := strconv.ParseInt(roundNum, 10, 64)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	opts := &bind.FilterOpts{
		Start:   uint64(big.NewInt(0).Int64()),
		End:     nil,
		Context: context.Background(),
	}

	logs, err := e.http.inst.FilterBid(opts, nil)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bidEvents := &models.BidEvents{}
	for logs.Next() {
		if logs.Error() != nil {
			e.logger.Error("Log error: %v", logs.Error())
			continue
		}

		event := logs.Event // Автоматически распарсенное событие
		if round == event.Round.Int64() {
			bidEvents.Events = append(bidEvents.Events, &models.BidEvent{
				Account:   event.Account.Hex(),
				Amount:    event.Amount.Int64(),
				Timestamp: event.Timestamp.Int64(),
				Round:     event.Round.Int64(),
			})
		}
	}

	b, err := json.Marshal(bidEvents)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(b))

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}