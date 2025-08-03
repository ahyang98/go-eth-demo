package block

import (
	"context"
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubscribeBlock() {
	client, err := ethclient.Dial(constant.TokenWS)
	if err != nil {
		panic(err)
	}
	headers := make(chan *types.Header)
	subscribeNewHead, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case err := <-subscribeNewHead.Err():
			panic(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				panic(err)
			}
			PrintBlock(block)
		}
	}
}
