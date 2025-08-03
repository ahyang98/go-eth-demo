package wallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func ETHTransfer() {
	client, err := ethclient.Dial(constant.BlockApi)
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA(constant.PrivateKey1)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	value := big.NewInt(0.0001 * constant.ETH2Wei)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("gasprice %+v\n", gasPrice)
	toAddress := common.HexToAddress("0x9914cd239b9480958d3C69702c417bCF7d6246B9")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		panic(err)
	}
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tx send: %s", signTx.Hash().Hex())
}
