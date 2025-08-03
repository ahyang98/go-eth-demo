package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"math/big"
	"time"
)

func TransferToken() {
	client, err := ethclient.Dial(constant.TokenApi)
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
		panic("public key is not ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Balance:", balance)

	nonceAt, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nonce:", nonceAt)

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("gas price:", gasPrice)
	gasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(15000000000))
	fmt.Println("gas price:", gasPrice)

	toAddress := common.HexToAddress("0x9914cd239b9480958d3C69702c417bCF7d6246B9")

	tokenAddress := common.HexToAddress("0xD2D9d80D20c169e975F3E8B2D6736E5b715a9E37")

	//transferFnSignature := []byte("transfer(address,uint256)")
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodId := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodId))

	paddedToAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedToAddress))

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) //1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedAmount...)

	fmt.Printf("data %+v\n", hexutil.Encode(data))

	gasLimit := uint64(1000000)
	tx := types.NewTransaction(nonceAt, tokenAddress, value, gasLimit, gasPrice, data)

	//chainID := big.NewInt(11155111)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("chainId:", chainID)

	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx send:", signTx.Hash().Hex())

	time.Sleep(30 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), signTx.Hash())
	if err != nil {
		fmt.Println("waiting for tx to be mined...")
	} else {
		fmt.Println("tx status:", receipt.Status) // 1 = 成功，0 = 失败
	}
}

//https://sepolia.etherscan.io/tx/0x67f83636428e84f223fb3ed2b2744776dea91191e1c45f14145d2b575c4de33b
