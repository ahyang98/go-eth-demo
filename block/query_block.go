package block

import (
	"context"
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Query() {
	client, err := ethclient.Dial(constant.BlockApi)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	PrintBlock(block)

}
