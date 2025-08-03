package block

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

func PrintBlock(block *types.Block) {
	fmt.Printf("latest block number %d\n", block.Number().Uint64())
	fmt.Printf("block hash %s\n", block.Hash().Hex())
	fmt.Printf("block timestamp %s\n", time.Unix(int64(block.Time()), 0).UTC().In(time.Local))
	fmt.Printf("block tx counts %d\n", block.Transactions().Len())
	fmt.Println("block nonce:", block.Nonce())
}
