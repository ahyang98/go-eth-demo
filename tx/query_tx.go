package tx

import (
	"context"
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func QueryTX() {
	client, err := ethclient.Dial(constant.BlockApi)
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
		return
	}
	blockNumber := big.NewInt(8880355) //0x261ad2a790d51ef3bce5119e5033d3f4571ec83f6fb5285a9e38c9024d214c01
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		panic(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())

		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender:", sender.Hex())
		} else {
			log.Fatal(err)
		}
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return
		}
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		break
	}
	blockHash := common.HexToHash("0x261ad2a790d51ef3bce5119e5033d3f4571ec83f6fb5285a9e38c9024d214c01")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		panic(err)
	}
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			panic(err)
		}
		fmt.Println(tx.Hash().Hex()) //0x110d2887b654cf55168dc5d504f80ce95ee886d6c3e0f13e4068d803eb50d68e
		break
	}
	txHash := common.HexToHash("0x110d2887b654cf55168dc5d504f80ce95ee886d6c3e0f13e4068d803eb50d68e")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		panic(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex())

}
