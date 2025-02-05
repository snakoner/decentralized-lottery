package eth

import (
	"context"
	"encoding/json"
	"log"
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

// todo: filter by indexed round

// RoundHandler is an HTTP handler that processes requests to retrieve bid events
// for a specific lottery round. It accepts a round number from the URL, filters
// the bid events from the Ethereum contract based on the provided round number,
// and returns the corresponding events in the response.
//
// Parameters:
//   - w: The HTTP response writer used to send responses back to the client.
//   - r: The HTTP request containing the round number in the URL.
//
// Response:
//   - A JSON object containing the list of bid events for the specified round if successful.
//   - A corresponding HTTP error status if there is an issue (e.g., 400 for bad requests or 500 for server errors).
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

	logs, err := e.http.inst.FilterBid(opts, nil, nil)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bidEvents := &models.BidResponse{}
	for logs.Next() {
		if logs.Error() != nil {
			e.logger.Error("Log error: %v", logs.Error())
			continue
		}

		block, err := e.http.cli.BlockByHash(context.Background(), logs.Event.Raw.BlockHash)
		if err != nil {
			log.Fatalf("Failed to retrieve block: %v", err)
		}

		// Extract the timestamp
		timestamp := block.Time()
		event := logs.Event
		if round == event.Round.Int64() {
			bidEvents.Events = append(bidEvents.Events, &models.BidEvent{
				Account:   event.Account.Hex(),
				Amount:    event.Amount.Int64(),
				Timestamp: int64(timestamp),
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

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// WinnerHandler is an HTTP handler that processes requests to retrieve information
// about the selected winner for a specific lottery round. It accepts the round number
// from the URL parameter, filters the winner selection events from the Ethereum contract,
// and returns the selected winner's details in the response.
//
// Parameters:
//   - w: The HTTP response writer used to send responses back to the client.
//   - r: The HTTP request containing the round number in the URL.
//
// Response:
//   - A JSON object containing the winner's details if successful.
//   - A corresponding HTTP error status if there is an issue (e.g., 400 for bad requests or 500 for server errors).
func (e *EthereumServer) WinnerHandler(w http.ResponseWriter, r *http.Request) {
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

	logs, err := e.http.inst.FilterWinnerSelected(opts, nil, nil)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// for now roundFinished: false
	winnerSelectedResponse := &models.WinnerSelectedResponse{}
	for logs.Next() {
		if logs.Error() != nil {
			e.logger.Error("Log error: %v", logs.Error())
			continue
		}

		block, err := e.http.cli.BlockByHash(context.Background(), logs.Event.Raw.BlockHash)
		if err != nil {
			log.Fatalf("Failed to retrieve block: %v", err)
		}

		// Extract the timestamp
		timestamp := block.Time()
		event := logs.Event
		if round == event.Round.Int64() {
			winnerSelectedResponse.Winner = models.WinnerSelectedEvent{
				Account:       event.Account.Hex(),
				Amount:        event.Amount.Int64(),
				Round:         event.Round.Int64(),
				Timestamp:     int64(timestamp),
				RoundFinished: true,
			}
		}
	}

	b, err := json.Marshal(winnerSelectedResponse)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (e *EthereumServer) AllTimeRewardHander(w http.ResponseWriter, r *http.Request) {
	opts := &bind.FilterOpts{
		Start:   uint64(big.NewInt(0).Int64()),
		End:     nil,
		Context: context.Background(),
	}

	logs, err := e.http.inst.FilterWinnerSelected(opts, nil, nil)
	allTImeReward := new(big.Int)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for logs.Next() {
		if logs.Error() != nil {
			e.logger.Error("Log error: %v", logs.Error())
			continue
		}

		allTImeReward = allTImeReward.Add(allTImeReward, logs.Event.Amount)
	}

	ether := new(big.Float).Quo(new(big.Float).SetInt(allTImeReward), big.NewFloat(1e18)).Text('f', 18)

	b, err := json.Marshal(&models.AllTimeRewardResponse{Reward: ether})
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
