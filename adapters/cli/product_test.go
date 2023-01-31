package cli_test

import (
	"fmt"
	"testing"
	mock_application "github.com/r31sr4r/go-ports-and-adapters/application/mocks"
	"github.com/golang/mock/gomock"
	cli "github.com/r31sr4r/go-ports-and-adapters/adapters/cli"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	productName := "Product 1"
	price := 10.0
	productStatus := "enabled"
	productID := "abc123"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()	

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with prince %f and status %s", 
		productID,
		productName,
		price,
		productStatus)

	result, err := cli.Run(service, "create", "", productName, price)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(service, "enable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(service, "disable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s\nName: %s\nPrice: %f\nStatus: %s",
		productID,
		productName,
		price,
		productStatus)
	result, err = cli.Run(service, "", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
	

}
