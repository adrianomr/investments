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

func TestCdiController(t *testing.T) {
	t.Run("Should return new cdi controller", func(t *testing.T) {
		result := controllers.NewCdiController()
		assert.NotNil(t, result)
		assert.NotNil(t, result.Create)
		assert.NotNil(t, result.Routes())
	})
}

func TestCreateCdi(t *testing.T) {
	ctrl := gomock.NewController(t)
	create := mock.NewMockICdiCreate(ctrl)
	restController := controllers.CdiController{Create: create}
	defer ctrl.Finish()

	t.Run("Should fail to create cdi when no body sent", func(t *testing.T) {

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdis",
			Path:   "/api/cdis",
		}, restController.CreateCdi)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	})

	t.Run("Should fail to create cdi when usecase failed", func(t *testing.T) {

		create.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, errors.New("usecase error"))

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdis",
			Path:   "/api/cdis",
			Body:   "{\"rate\": 0.5, \"date\": \"2024-05-27T19:47:46Z\"}",
		}, restController.CreateCdi)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode())
	})

	t.Run("Should create cdi", func(t *testing.T) {

		create.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, nil)

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodPost,
			Url:    "/api/cdis",
			Path:   "/api/cdis",
			Body:   "{\"rate\": 0.5, \"date\": \"2024-05-27T19:47:46Z\"}",
		}, restController.CreateCdi)

		assert.Equal(t, http.StatusCreated, resp.StatusCode())
	})
}
