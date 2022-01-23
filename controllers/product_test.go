package controllers

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/gipsh/basic-contract-interaction/services"
	"github.com/gipsh/basic-contract-interaction/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	productService := mock.NewMockProductService(ctrl)

	productController := ProductController{
		ProductService: productService,
	}
	items := []services.ProductItem{
		{
			Name:   "PROD1",
			Status: 1,
		},
		{
			Name:   "PROD2",
			Status: 0,
		},
	}

	s := make([]interface{}, len(items))
	for i, v := range items {
		s[i] = v
	}

	productService.EXPECT().GetProducts().Return(s, nil)

	assert.NotNil(t, productController)

	productController.Products(c)

	assert.Equal(t, 200, w.Code)

}

func TestGetProductById(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	productService := mock.NewMockProductService(ctrl)

	productController := ProductController{
		ProductService: productService,
	}

	productService.EXPECT().GetProductById(int64(1)).Return(services.ProductItem{
		Name:   "chori",
		Status: 0,
		Owner:  common.HexToAddress("0xdead"),
	}, nil)

	assert.NotNil(t, productController)

	productController.ProductById(c)

	assert.Equal(t, 200, w.Code)

}

func TestGetProductByIdError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	productService := mock.NewMockProductService(ctrl)

	productController := ProductController{
		ProductService: productService,
	}

	productService.EXPECT().GetProductById(int64(1)).Return(services.ProductItem{}, fmt.Errorf(services.ErrorTx))

	assert.NotNil(t, productController)

	productController.ProductById(c)

	assert.Equal(t, 400, w.Code)

}
