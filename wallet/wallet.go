package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func Create() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // bfc1b13b04378adae024a9a74886fe73069100cc39782130c8a0098e4595dff7
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("publickey is not *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubkey:", hexutil.Encode(publicKeyBytes)[4:]) //5524743e72d12691cffcaf5f5b3110d972bdf45ef309b01dd976ca5505201196f7eb432f82e3274968f0d8fb5330dc260d632e4ac05c67e3ed6ff037370e445f
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:])) //0x206d51d52e52d3790898f9241e38fbb65e70a3f189b3bfc24c56133b353c044c
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))        //0x1e38fbb65e70a3f189b3bfc24c56133b353c044c
}
