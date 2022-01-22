package services_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gipsh/basic-contract-interaction/services"
	"github.com/gipsh/basic-contract-interaction/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock.NewMockProductService(ctrl)

	productService.EXPECT().GetProducts().Return([]interface{}{
		services.ProductItem{
			Name:     "chori",
			Status:   0,
			Owner:    common.HexToAddress("0xdeadbeaf"),
			NewOwner: common.HexToAddress(""),
		},
	}, nil)

	items, err := productService.GetProducts()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 1, len(items))

}
