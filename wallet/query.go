package wallet

import (
	"context"
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

func QueryBalance() {
	client, err := ethclient.Dial(constant.BlockApi)
	if err != nil {
		panic(err)
	}
	account := common.HexToAddress("0xD3D0124bDD8163C317bb0571baD421E9521Ac894")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(balance)
	blockNumber := big.NewInt(8888000)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println(balanceAt)
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
	pendingBalanceAt, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		panic(err)
	}
	fmt.Println(pendingBalanceAt)
}

//60700957809979104
//13574331160256632
//0.013574331160256632
//60700957809979104
