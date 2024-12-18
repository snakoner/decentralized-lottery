package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/snakoner/lottery/internal/config"
	"github.com/snakoner/lottery/internal/lottery"
)

type Signer struct {
	privateKey *ecdsa.PrivateKey
}
type EthereumServer struct {
	http struct {
		cli  *ethclient.Client
		inst *lottery.Lottery
	}
	wss struct {
		cli  *ethclient.Client
		inst *lottery.Lottery
	}
	cli             *ethclient.Client
	logger          *logrus.Logger
	contractAddress common.Address
	signer          *Signer
	newEvents       []*lottery.LotteryBid
}

func New(config *config.Config) (*EthereumServer, error) {
	address := common.HexToAddress(config.Contract.Address)
	httpCli, err := ethclient.Dial(fmt.Sprintf("%s%s", config.Infura.HttpProvider, config.Private.Provider))
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.Private.Evm)
	if err != nil {
		return nil, err
	}

	httpInst, err := lottery.NewLottery(address, httpCli)

	wssCli, err := ethclient.Dial(fmt.Sprintf("%s%s", config.Infura.WssProvider, config.Private.Provider))
	if err != nil {
		return nil, err
	}

	wssInst, err := lottery.NewLottery(address, wssCli)
	if err != nil {
		return nil, err
	}

	return &EthereumServer{
		logger:          logrus.New(),
		contractAddress: address,
		signer: &Signer{
			privateKey: privateKey,
		},
		http: struct {
			cli  *ethclient.Client
			inst *lottery.Lottery
		}{
			cli:  httpCli,
			inst: httpInst,
		},

		wss: struct {
			cli  *ethclient.Client
			inst *lottery.Lottery
		}{
			cli:  wssCli,
			inst: wssInst,
		},
	}, nil
}

func (e *EthereumServer) getTimeLeft() (int64, error) {
	res, err := e.http.inst.GetTimeLeft(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	return res.Int64(), nil
}

func (e *EthereumServer) getParticipantsNum() (int64, error) {
	res, err := e.http.inst.GetParticipantsNumber(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	return res.Int64(), nil
}

func (e *EthereumServer) restartLottery(newDuration int64) error {
	auth := bind.NewKeyedTransactor(e.signer.privateKey)
	auth.GasLimit = uint64(3000000) // in units
	gasPrice, err := e.http.cli.SuggestGasPrice(context.Background())
	if err != nil {
		e.logger.Error(err)
		return err
	}

	duration := new(big.Int)
	duration.SetInt64(newDuration)
	auth.GasPrice = gasPrice

	_, err = e.http.inst.RestartEmpty(auth, duration)
	if err != nil {
		e.logger.Error(err)
		return err
	}

	return nil
}

func (e *EthereumServer) start() error {
	auth := bind.NewKeyedTransactor(e.signer.privateKey)
	auth.GasLimit = uint64(3000000) // in units
	gasPrice, err := e.http.cli.SuggestGasPrice(context.Background())
	if err != nil {
		e.logger.Error(err)
		return err
	}

	auth.GasPrice = gasPrice

	_, err = e.http.inst.Start(auth)
	if err != nil {
		e.logger.Error(err)
		return err
	}

	return nil
}

func (e *EthereumServer) SelectWinner(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			e.logger.Info("stop SelectWinner() function")
			return
		case <-time.After(time.Second * 10):
			timeLeft, err := e.getTimeLeft()
			if err != nil {
				e.logger.Error(err)
				return
			}

			if timeLeft == 0 {
				e.logger.Info("time is over")
				parts, err := e.getParticipantsNum()
				if err != nil {
					e.logger.Error(err)
					return
				}

				if parts == 0 {
					// dont set new duration
					if err := e.restartLottery(0); err != nil {
						e.logger.Error(err)
						return
					}
				} else {
					if err := e.start(); err != nil {
						e.logger.Error(err)
						return
					}
				}
			}
		}
	}
}

func (e *EthereumServer) Stop() {
	if e.http.cli != nil {
		e.http.cli.Close()
	}
	if e.wss.cli != nil {
		e.wss.cli.Close()
	}
}
