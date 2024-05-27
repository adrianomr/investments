package controllers

import (
	"errors"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"

	"github.com/adrianomr/investments/src/application/controllers"
	"github.com/adrianomr/investments/src/domain/usecases/mock"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/stretchr/testify/assert"
)

func TestCdbController(t *testing.T) {
	t.Run("Should return new cdb controller", func(t *testing.T) {
		result := controllers.NewCdbController()
		assert.NotNil(t, result)
		assert.NotNil(t, result.Create)
		assert.NotNil(t, result.CreateOrder)
		assert.NotNil(t, result.Routes())
	})
}

func TestCreateCdbOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	create := mock.NewMockICdbCreate(ctrl)
	createOrder := mock.NewMockICdbOrderCreate(ctrl)
	restController := controllers.CdbController{Create: create, CreateOrder: createOrder}
	defer ctrl.Finish()

	t.Run("Should return bad request when cdb id invalid", func(t *testing.T) {

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs/invalid/orders",
			Path:   "/api/cdbs/{cdb_id}/orders",
			Body:   "{\"amount\": 150.55, \"type\": \"BUY\"}",
		}, restController.CreateCdbOrder)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	})

	t.Run("Should fail to create cdb order when no body sent", func(t *testing.T) {

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs/b0dae81c-e55c-4ca8-b635-2f3087d6b590/orders",
			Path:   "/api/cdbs/{cdb_id}/orders",
		}, restController.CreateCdbOrder)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	})

	t.Run("Should fail to create cdb order when no type sent", func(t *testing.T) {

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs/b0dae81c-e55c-4ca8-b635-2f3087d6b590/orders",
			Path:   "/api/cdbs/{cdb_id}/orders",
			Body:   "{\"amount\": 150.55}",
		}, restController.CreateCdbOrder)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	})

	t.Run("Should fail to create cdb order when no amount sent", func(t *testing.T) {

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs/b0dae81c-e55c-4ca8-b635-2f3087d6b590/orders",
			Path:   "/api/cdbs/{cdb_id}/orders",
			Body:   "{\"type\": \"BUY\"}",
		}, restController.CreateCdbOrder)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	})

	t.Run("Should fail to create cdb order when usecase failed", func(t *testing.T) {

		createOrder.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, errors.New("usecase error"))

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs/b0dae81c-e55c-4ca8-b635-2f3087d6b590/orders",
			Path:   "/api/cdbs/{cdb_id}/orders",
			Body:   "{\"amount\": 150.55, \"type\": \"BUY\"}",
		}, restController.CreateCdbOrder)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode())
	})

	t.Run("Should create cdb order", func(t *testing.T) {

		createOrder.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, nil)

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs/b0dae81c-e55c-4ca8-b635-2f3087d6b590/orders",
			Path:   "/api/cdbs/{cdb_id}/orders",
			Body:   "{\"amount\": 150.55, \"type\": \"BUY\"}",
		}, restController.CreateCdbOrder)

		assert.Equal(t, http.StatusCreated, resp.StatusCode())
	})
}

func TestCreateCdb(t *testing.T) {
	ctrl := gomock.NewController(t)
	create := mock.NewMockICdbCreate(ctrl)
	createOrder := mock.NewMockICdbOrderCreate(ctrl)
	restController := controllers.CdbController{Create: create, CreateOrder: createOrder}
	defer ctrl.Finish()

	t.Run("Should fail to create cdb when no body sent", func(t *testing.T) {

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs",
			Path:   "/api/cdbs",
		}, restController.CreateCdb)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	})

	t.Run("Should fail to create cdb when usecase failed", func(t *testing.T) {

		create.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, errors.New("usecase error"))

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs",
			Path:   "/api/cdbs",
			Body:   "{\"percentage\": 0.5, \"type\": \"cdi\"}",
		}, restController.CreateCdb)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode())
	})

	t.Run("Should create cdb", func(t *testing.T) {

		create.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, nil)

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdbs",
			Path:   "/api/cdbs",
			Body:   "{\"percentage\": 0.5, \"type\": \"cdi\"}",
		}, restController.CreateCdb)

		assert.Equal(t, http.StatusCreated, resp.StatusCode())
	})
}
