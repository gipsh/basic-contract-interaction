package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"os"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gipsh/basic-contract-interaction/contracts"
)

type ProductItem struct {
	Name     string
	Status   uint8
	Owner    common.Address
	NewOwner common.Address
}

const (
	ErrorTx = "Error"
)

type ProductService interface {
	GetProducts() ([]interface{}, error)
	GetProductById(productId int64) (ProductItem, error)
	GetProductByName(name string) ([]interface{}, error)
	CreateProduct(name string) (string, error) //txHash
	DelegateProduct(productId int64, newOwner string) (string, error)
	AcceptProduct(productId int64, newOwner string) (string, error)
}

type ProductSerivceImp struct {
	Client  *ethclient.Client
	ChainId *big.Int
}

func NewProductService() ProductService {

	var err error
	ps := ProductSerivceImp{}

	ps.Client, err = ethclient.Dial(os.Getenv("PROVIDER"))
	if err != nil {
		log.Fatal(err)
	}

	ps.ChainId, err = ps.Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return ps

}

func (ps *ProductSerivceImp) buildSignedTxWithKey(privKey string) (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ps.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ps.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, ps.ChainId)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Signer = ps.keySigner(ps.ChainId, privateKey)

	return auth, nil

}

func (ps *ProductSerivceImp) buildSignedTx() (*bind.TransactOpts, error) {

	return ps.buildSignedTxWithKey(os.Getenv("PRIVATE_KEY"))

}

func (ps *ProductSerivceImp) keySigner(chainID *big.Int, key *ecdsa.PrivateKey) (signerfn bind.SignerFn) {
	signerfn = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		keyAddr := crypto.PubkeyToAddress(key.PublicKey)
		if address != keyAddr {
			return nil, errors.New("not authorized to sign this account")
		}
		return types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	}
	return
}

func (ps *ProductSerivceImp) getInstance() (*contracts.Product, error) {

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	instance, err := contracts.NewProduct(contractAddress, ps.Client)
	if err != nil {
		return nil, err
	}

	return instance, nil

}

func (ps *ProductSerivceImp) getAllProducts() ([]interface{}, error) {

	instance, err := ps.getInstance()
	if err != nil {
		return nil, err
	}

	pSize, err := instance.Size(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	var result []interface{}
	for i := 1; i < int(pSize.Int64()); i++ {
		p, err := instance.Products(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

func (ps ProductSerivceImp) GetProducts() ([]interface{}, error) {

	result, err := ps.getAllProducts()
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (ps ProductSerivceImp) GetProductById(productId int64) (ProductItem, error) {

	instance, err := ps.getInstance()
	if err != nil {
		return ProductItem{}, err
	}

	p, err := instance.Products(&bind.CallOpts{}, big.NewInt(int64(productId)))
	if err != nil {
		return ProductItem{}, err
	}

	return p, nil

}

func (ps ProductSerivceImp) GetProductByName(name string) ([]interface{}, error) {

	items, err := ps.getAllProducts()
	if err != nil {
		return nil, err
	}

	var result []interface{}
	for _, p := range items {
		v := reflect.ValueOf(p).Elem().FieldByName("Name")
		if v.String() == name {
			result = append(result, result...)
		}
	}

	return result, nil

}

func (ps ProductSerivceImp) CreateProduct(name string) (string, error) {

	instance, err := ps.getInstance()
	if err != nil {
		return ErrorTx, err
	}

	signedTxOpts, err := ps.buildSignedTx()
	if err != nil {
		return ErrorTx, err
	}

	result, err := instance.CreateProduct(signedTxOpts, name)
	if err != nil {
		return ErrorTx, err
	}

	return result.Hash().Hex(), nil

}

func (ps ProductSerivceImp) DelegateProduct(productId int64, newOwner string) (string, error) {

	instance, err := ps.getInstance()
	if err != nil {
		return ErrorTx, err
	}

	signedTxOpts, err := ps.buildSignedTx()
	if err != nil {
		return ErrorTx, err
	}

	result, err := instance.DelegateProduct(signedTxOpts, big.NewInt(productId), common.HexToAddress(newOwner))
	if err != nil {
		return ErrorTx, err
	}

	return result.Hash().Hex(), nil

}

func (ps ProductSerivceImp) AcceptProduct(productId int64, newOwner string) (string, error) {

	instance, err := ps.getInstance()
	if err != nil {
		return ErrorTx, err
	}

	signedTxOpts, err := ps.buildSignedTxWithKey(newOwner)
	if err != nil {
		return ErrorTx, err
	}

	result, err := instance.AcceptProduct(signedTxOpts, big.NewInt(productId))
	if err != nil {
		return ErrorTx, err
	}

	return result.Hash().Hex(), nil

}
