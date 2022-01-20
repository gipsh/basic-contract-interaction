package controllers

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/gipsh/basic-contract-interaction/contracts"
)

type ProductController struct {
	Singer         bind.SignerFn
	ProductInstace *contracts.Product
	SignedTxOpts   *bind.TransactOpts
}

type CreateProductInput struct {
	Name string `json:"name"`
}

type DelegateProductInput struct {
	ProductId int64  `json:"product_id"`
	Owner     string `json:"owner"`
}

type AcceptProductInput struct {
	ProductId int64 `json:"product_id"`
}

func NewProductController() ProductController {

	pc := ProductController{}

	client, err := ethclient.Dial(os.Getenv("PROVIDER"))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	pc.Singer = KeySigner(chainID, privateKey)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Signer = pc.Singer

	pc.SignedTxOpts = auth

	contractAddress := common.HexToAddress(os.Getenv("PRIVATE_ADDRESS"))

	pc.ProductInstace, err = contracts.NewProduct(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return pc
}

func KeySigner(chainID *big.Int, key *ecdsa.PrivateKey) (signerfn bind.SignerFn) {
	signerfn = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		keyAddr := crypto.PubkeyToAddress(key.PublicKey)
		if address != keyAddr {
			return nil, errors.New("not authorized to sign this account")
		}
		return types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	}
	return
}

func (pc *ProductController) Products(c *gin.Context) {

	result, err := pc.ProductInstace.Products(&bind.CallOpts{}, big.NewInt(0))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, result)

}

func (pc *ProductController) ProductByName(c *gin.Context) {

}

func (pc *ProductController) ProductById(c *gin.Context) {

}

func (pc *ProductController) CreateProduct(c *gin.Context) {

	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pc.ProductInstace.CreateProduct(pc.SignedTxOpts, input.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, result)

}

func (pc *ProductController) DelegateProduct(c *gin.Context) {

	var input DelegateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	owner := common.HexToAddress(input.Owner)
	productId := big.NewInt(input.ProductId)
	result, err := pc.ProductInstace.DelegateProduct(pc.SignedTxOpts, productId, owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, result)

}

func (pc *ProductController) AcceptProduct(c *gin.Context) {

	var input AcceptProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productId := big.NewInt(input.ProductId)
	result, err := pc.ProductInstace.AcceptProduct(pc.SignedTxOpts, productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, result)

}
