package application_test

import (
	"github.com/r31sr4r/go-ports-and-adapters/application"	
	"github.com/stretchr/testify/require"
	"testing"
	uuid "github.com/satori/go.uuid"
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

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0.0

	err := product.Disable()
	require.Nil(t, err)	

	product.Price = 10.0
	err = product.Disable()
	require.Equal(t, "Price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10.0

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "Status must be 'enabled' or 'disabled'", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10.0
	_, err = product.IsValid()
	require.Equal(t, "Price must be greater than zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10.0

	product_ID_getter := product.GetID()

	require.Equal(t, product.ID, product_ID_getter)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10.0

	product_name_getter := product.GetName()

	require.Equal(t, product.Name, product_name_getter)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10.0

	product_status_getter := product.GetStatus()

	require.Equal(t, product.Status, product_status_getter)
}

