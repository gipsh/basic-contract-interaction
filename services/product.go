package services

import (
	"context"
	"crypto/ecdsa"
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

	ProductInstace *contracts.Product
	SignedTxOpts   *bind.TransactOpts
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

func (ps *ProductSerivceImp) buildSignedTxFrom(from common.Address) (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
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
	auth.From = from
	auth.Signer = ps.keySigner(ps.ChainId, privateKey)

	return auth, nil

}

func (ps *ProductSerivceImp) buildSignedTx() (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
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

func (ps *ProductSerivceImp) keySigner(chainID *big.Int, key *ecdsa.PrivateKey) (signerfn bind.SignerFn) {
	signerfn = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		// keyAddr := crypto.PubkeyToAddress(key.PublicKey)
		// if address != keyAddr {
		// 	return nil, errors.New("not authorized to sign this account")
		// }
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

	signedTxOpts, err := ps.buildSignedTxFrom(common.HexToAddress(newOwner))
	if err != nil {
		return ErrorTx, err
	}

	result, err := instance.AcceptProduct(signedTxOpts, big.NewInt(productId))
	if err != nil {
		return ErrorTx, err
	}

	return result.Hash().Hex(), nil

}

/*

Se desea tener una api donde en ella se puedan realizar las siguientes operaciones:

1 - Ver todos los productos disponibles con el detalle correspondiente.
2 - Crear un producto nuevo desde una wallet configurada en el backend.
3 - Delegar un producto a una wallet desde una wallet configurada en el backend.
4 - Aceptar una delegación de un producto.
5 - Ver todos los productos que tiene delegados determinada wallet.


Cada una de estas operaciones deberá interactuar con un smart contract que se encuentra en la testnet mumbai de polygon, es decir la api será una pasarela donde se comunicará con dicha blockchain. La address de dicho smart contract es:0xd9E0b2C0724F3a01AaECe3C44F8023371f845196 y la  especificación del modelo de datos se encuentra allí.

La api se podrá realizar en node, python, java, go. Es deseable que sea en python3 y debe utilizar la lib web3

No es necesario realizar todos los puntos, si los realiza es mejor y si tiene que optar por priorizar asuma que a menor valor numérico en el listado mayor prioridad

Se valora todo el background que considere colocar como tests, readmes, documentación, etc.

*/