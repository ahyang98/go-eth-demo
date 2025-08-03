package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ahyang98/go-eth-demo/constant"
	"github.com/ahyang98/go-eth-demo/token/store"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

func InvokeWithoutAbi() {
	const TokenAddress = "0xbc686328e0c5859299a205c995d7fc6ce5f497ec"
	client, err := ethclient.Dial(constant.TokenApi)
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA(constant.PrivateKey1)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyEcdsa, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("wrong format ecdsa public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyEcdsa)
	nonceAt, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	gasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(10000000000))

	//tokenJson, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	//if err != nil {
	//	panic(err)
	//}
	//methodName := "setItem"
	//var key [32]byte
	//var value [32]byte
	//copy(key[:], []byte("demo_save_key7"))
	//copy(value[:], []byte("demo_save_value7"))
	methodSignature := []byte("setItem(bytes32,bytes32)")
	methodSelector := crypto.Keccak256(methodSignature)[:4]

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_no_use_abi"))
	copy(value[:], []byte("demo_save_value_no_use_abi_11111"))

	//input, err := tokenJson.Pack(methodName, key, value)
	//if err != nil {
	//	panic(err)
	//}

	// 组合调用数据
	var input []byte
	input = append(input, methodSelector...)
	input = append(input, key[:]...)
	input = append(input, value[:]...)

	chainId := big.NewInt(int64(11155111))
	to := common.HexToAddress(TokenAddress)
	tx := types.NewTransaction(nonceAt, to, big.NewInt(0), 3000000, gasPrice, input)
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		panic(err)
	}
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction sent: %s\n", signTx.Hash().Hex())

	_, err = waitForReceipt(client, signTx.Hash())
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Minute)

	//callInput, err := tokenJson.Pack("items", key)
	//if err != nil {
	//	panic(err)
	//}
	//
	//callMsg := ethereum.CallMsg{
	//	To:   &to,
	//	Data: callInput,
	//}
	itemsSignature := []byte("items(bytes32)")
	itemsSelector := crypto.Keccak256(itemsSignature)[:4]

	var callInput []byte
	callInput = append(callInput, itemsSelector...)
	callInput = append(callInput, key[:]...)

	//to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		panic(err)
	}
	var unpacked [32]byte
	//err = tokenJson.UnpackIntoInterface(&unpacked, "items", result)
	//if err != nil {
	//	panic(err)
	//}
	copy(unpacked[:], result)

	fmt.Printf("the value %+v\n", hexutil.Encode(value[:]))
	fmt.Printf("the item %+v\n", hexutil.Encode(unpacked[:]))
}

func InvokeWithAbi() {
	const TokenAddress = "0xbc686328e0c5859299a205c995d7fc6ce5f497ec"
	client, err := ethclient.Dial(constant.TokenApi)
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA(constant.PrivateKey1)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyEcdsa, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("wrong format ecdsa public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyEcdsa)
	nonceAt, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	gasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(10000000000))

	tokenJson, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	methodName := "setItem"
	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key7"))
	copy(value[:], []byte("demo_save_value7"))
	input, err := tokenJson.Pack(methodName, key, value)
	if err != nil {
		panic(err)
	}
	chainId := big.NewInt(int64(11155111))
	to := common.HexToAddress(TokenAddress)
	tx := types.NewTransaction(nonceAt, to, big.NewInt(0), 3000000, gasPrice, input)
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		panic(err)
	}
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction sent: %s\n", signTx.Hash().Hex())

	_, err = waitForReceipt(client, signTx.Hash())
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Minute)

	callInput, err := tokenJson.Pack("items", key)
	if err != nil {
		panic(err)
	}

	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		panic(err)
	}
	var unpacked [32]byte
	err = tokenJson.UnpackIntoInterface(&unpacked, "items", result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the value %+v\n", hexutil.Encode(value[:]))
	fmt.Printf("the item %+v\n", hexutil.Encode(unpacked[:]))
}

func InvokeWithGenGo() {
	const TokenAddress = "0xbc686328e0c5859299a205c995d7fc6ce5f497ec"
	client, err := ethclient.Dial(constant.TokenApi)
	if err != nil {
		panic(err)
	}
	token, err := store.NewToken(common.HexToAddress(TokenAddress), client)
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA(constant.PrivateKey1)
	if err != nil {
		panic(err)
	}
	var key [32]byte
	var value [32]byte
	//
	copy(key[:], []byte("demo_save_key6"))
	copy(value[:], []byte("demo_save_value5"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		panic(err)
	}
	tx, err := token.SetItem(opt, key, value)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	receipt, err := waitForReceipt(client, tx.Hash())
	if err != nil {
		panic(err)
	}

	fmt.Println("status", receipt.Status)

	time.Sleep(time.Minute)

	callOpt := &bind.CallOpts{Context: context.Background()}
	item, err := token.Items(callOpt, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the value %+v\n", hexutil.Encode(value[:]))
	fmt.Printf("the item %+v\n", hexutil.Encode(item[:]))
}
