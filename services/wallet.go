package services

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type WalletService interface {
	GetPrivateKey(common.Address) (string, error)
	AddAddress(common.Address, string)
	AddAddressPK(string)
}

type WalletServiceImp struct {
	Secrets map[common.Address]string
}

func NewWalletService() WalletService {
	wallet := WalletServiceImp{}
	wallet.Secrets = make(map[common.Address]string)
	return wallet
}

// returns private key of a managed address
func (wallet WalletServiceImp) GetPrivateKey(address common.Address) (string, error) {

	value, ok := wallet.Secrets[address]
	if ok {
		return value, nil
	} else {
		return ErrorTx, fmt.Errorf("key not found")
	}
}

func (wallet WalletServiceImp) AddAddress(address common.Address, privateKey string) {
	wallet.Secrets[address] = privateKey
}

func (wallet WalletServiceImp) AddAddressPK(privKey string) {

	privateKey, _ := crypto.HexToECDSA(privKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	pubAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	wallet.AddAddress(pubAddress, privKey)

}
