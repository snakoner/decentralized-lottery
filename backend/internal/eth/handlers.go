package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
	"github.com/snakoner/lottery/internal/lottery"
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

func (e *EthereumServer) getCurrentBlockNumber() (uint64, error) {
	blockNumber, err := e.http.cli.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	return blockNumber, err
}

func (e *EthereumServer) PastRounds(w http.ResponseWriter, r *http.Request) {
	roundNum := mux.Vars(r)["number"]
	round, err := strconv.ParseInt(roundNum, 10, 64)
	if err != nil {
		e.logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventSig := []byte("Bid(address,uint256,uint256,uint256)")
	eventHash := crypto.Keccak256Hash(eventSig)
	valueToFilter := big.NewInt(round)

	fmt.Println("value to filter: ", valueToFilter)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0), // Начальный блок
		ToBlock:   nil,           // Конечный блок (nil для текущего)
		Addresses: []common.Address{e.contractAddress},
		Topics:    [][]common.Hash{{eventHash}, {common.BigToHash(valueToFilter)}},
	}

	logs, err := e.http.cli.FilterLogs(context.Background(), query)
	if err != nil {
		e.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(len(logs))

	contractABI, err := abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":true,"name":"account","type":"address"},{"indexed":false,"name":"amount","type":"uint256"},{"indexed":false,"name":"timestamp","type":"uint256"},{"indexed":true,"name":"round","type":"uint256"}],"name":"Bid","type":"event"}]`))
	for _, vLog := range logs {
		fmt.Printf("Raw Log: %+v\n", vLog)

		// Распаковка данных события
		// var result struct {
		// 	Value  *big.Int
		// 	Sender common.Address
		// }
		result := &lottery.ContractBid{}
		err := contractABI.UnpackIntoInterface(&result, "Bid", vLog.Data)
		if err != nil {
			log.Printf("Failed to unpack log data: %v", err)
			continue
		}

		// fmt.Printf("Decoded Event: Value=%d, Sender=%s\n", result.Account[0], result.Sender.Hex())
		fmt.Println(result)
	}

}
