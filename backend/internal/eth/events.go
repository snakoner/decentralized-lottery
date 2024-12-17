package eth

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/snakoner/lottery/internal/lottery"
)

func (e *EthereumServer) SubscribeBidEvent(ctx context.Context) {
	ch := make(chan *lottery.LotteryBid)
	event, err := e.wss.inst.WatchBid(&bind.WatchOpts{}, ch, []common.Address{}, nil)
	if err != nil {
		e.logger.Error(err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			e.logger.Info("stop listening bid event")
			return
		case err := <-event.Err():
			e.logger.Error(err)
		case vLog := <-ch:
			e.newEvents = append(e.newEvents, vLog)
			e.logger.Info(fmt.Sprintf("[Bid event]: account: %v, amount: %v, ts: %v, round: %v ", vLog.Account, vLog.Amount, vLog.Timestamp, vLog.Round))
			// make write to db
		}
	}
}
