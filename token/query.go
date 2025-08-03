package token

import (
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

func Query() {
	client, err := ethclient.Dial(constant.BlockApi)
	if err != nil {
		panic(err)
	}
	tokenAddress := common.HexToAddress("0xD2D9d80D20c169e975F3E8B2D6736E5b715a9E37")
	instance, err := NewToken(tokenAddress, client)
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress("0xD3D0124bDD8163C317bb0571baD421E9521Ac894")
	balanceOf, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wei: %s\n", balanceOf)
	fbalance := new(big.Float)
	fbalance.SetString(balanceOf.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
}
