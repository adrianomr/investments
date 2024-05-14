package controllers

import (
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"

	"github.com/adrianomr/investments/src/application/controllers"
	"github.com/adrianomr/investments/src/domain/usecases/mock"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/stretchr/testify/assert"
)

func TestAppController(t *testing.T) {
	t.Run("Should return new cdb controller", func(t *testing.T) {
		result := controllers.NewCdbController()
		assert.NotNil(t, result)
		assert.NotNil(t, result.Routes())
	})
}

func TestShouldCreateCdbOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	create := mock.NewMockICdbOrderCreate(ctrl)
	restController := controllers.CdbController{Create: create}
	defer ctrl.Finish()

	t.Run("Should create cdb order", func(t *testing.T) {

		create.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, nil)

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs/orders",
			Path:   "/api/cdbs/orders",
		}, restController.CreateCdb)

		assert.Equal(t, http.StatusCreated, resp.StatusCode())
	})
}
