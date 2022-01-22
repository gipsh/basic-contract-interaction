package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gipsh/basic-contract-interaction/services"
)

type ProductController struct {
	ProductService services.ProductService
}

type CreateProductInput struct {
	Name string `json:"name"`
}

type DelegateProductInput struct {
	ProductId int64  `json:"product_id"`
	Owner     string `json:"owner"`
}

type AcceptProductInput struct {
	ProductId int64  `json:"product_id"`
	Owner     string `json:"new_owner"`
}

func NewProductController() ProductController {
	pc := ProductController{}
	pc.ProductService = services.NewProductService()

	return pc
}

func (pc *ProductController) Products(c *gin.Context) {

	result, err := pc.ProductService.GetProducts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, result)

}

func (pc *ProductController) ProductByName(c *gin.Context) {

	var name string
	var err error

	if v, ok := c.Params.Get("name"); ok {
		name = v
	}

	result, err := pc.ProductService.GetProductByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	result, err := pc.ProductService.GetProductById(productId)
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

	tx, err := pc.ProductService.CreateProduct(input.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"tx": tx})

}

func (pc *ProductController) DelegateProduct(c *gin.Context) {

	var input DelegateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//owner = common.HexToAddress("0xCF6380c9B128941d20d9F812dA406A79424b4B7B")

	tx, err := pc.ProductService.DelegateProduct(input.ProductId, input.Owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"tx": tx})

}

func (pc *ProductController) AcceptProduct(c *gin.Context) {

	var input AcceptProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := pc.ProductService.AcceptProduct(input.ProductId, input.Owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"tx": tx})

}
