package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/igor-marchi/go-hexagonal/application"
	mock_application "github.com/igor-marchi/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockIProductInterface(controller)
	persistence := mock_application.NewMockIProductPersistenceInterface(controller)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}