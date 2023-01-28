package application_test

import (
	"github.com/r31sr4r/go-ports-and-adapters/application"	
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10.0

	err := product.Enable()
	require.Nil(t, err)	

	product.Price = 0.0
	err = product.Enable()
	require.Equal(t, "Price must be greater than zero to enable the product", err.Error())
}

