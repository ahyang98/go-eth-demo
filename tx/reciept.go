package tx

import (
	"context"
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

func QueryReceipt() {
	client, err := ethclient.Dial(constant.BlockApi)
	if err != nil {
		panic(err)
	}
	blockNumber := big.NewInt(8880355) //0x261ad2a790d51ef3bce5119e5033d3f4571ec83f6fb5285a9e38c9024d214c01
	blockHash := common.HexToHash("0x261ad2a790d51ef3bce5119e5033d3f4571ec83f6fb5285a9e38c9024d214c01")
	blockReceiptsByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		panic(err)
	}
	blockReceiptsByNumber, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", blockReceiptsByNumber[0])
	fmt.Printf("%+v\n", blockReceiptsByHash[0])
	//fmt.Println(*(blockReceiptsByNumber[0]) == *(blockReceiptsByHash[0]))
	for _, receipt := range blockReceiptsByHash {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.TxHash.Hex()) //0x110d2887b654cf55168dc5d504f80ce95ee886d6c3e0f13e4068d803eb50d68e
		fmt.Println(receipt.TransactionIndex)
		fmt.Println(receipt.ContractAddress.Hex())
		break
	}
	txHash := common.HexToHash("0x110d2887b654cf55168dc5d504f80ce95ee886d6c3e0f13e4068d803eb50d68e")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		panic(err)
	}
	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs)
	fmt.Println(receipt.TxHash.Hex()) //0x110d2887b654cf55168dc5d504f80ce95ee886d6c3e0f13e4068d803eb50d68e
	fmt.Println(receipt.TransactionIndex)
	fmt.Println(receipt.ContractAddress.Hex())

}
