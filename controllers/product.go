package controllers

import (
	"math/big"
	"net/http"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gipsh/basic-contract-interaction/contracts"
)

type ProductController struct {
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

func (pc *ProductController) getAllProducts(c *gin.Context) []interface{} {

	pSize, err := pc.ProductInstace.Size(&bind.CallOpts{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var result []interface{}
	for i := 1; i < int(pSize.Int64()); i++ {
		p, err := pc.ProductInstace.Products(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		result = append(result, p)
	}

	return result
}

func (pc *ProductController) Products(c *gin.Context) {

	result := pc.getAllProducts(c)

	c.JSON(200, result)

}

func (pc *ProductController) ProductByName(c *gin.Context) {

	products := pc.getAllProducts(c)

	var result []interface{}
	for _, p := range products {

		v := reflect.ValueOf(p).Elem().FieldByName("Name")
		if v.String() == "chori" {
			result = append(result, result...)
		}
	}

	c.JSON(200, result)

}

func (pc *ProductController) ProductById(c *gin.Context) {

	var productId int64
	var err error

	if v, ok := c.Params.Get("id"); ok {
		productId, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	result, err := pc.ProductInstace.Products(&bind.CallOpts{}, big.NewInt(productId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, result)

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
