package controllers

import (
	"net/http"
	"testing"

	"github.com/adrianomr/investments/src/application/controllers"
	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/domain/usecases/mock"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAppController(t *testing.T) {
	t.Run("Should return new investment controller", func(t *testing.T) {
		result := controllers.NewInvestmentsController()
		assert.NotNil(t, result)
		assert.NotNil(t, result.Routes())
	})
}

func TestGetAllInvestment(t *testing.T) {
	ctrl := gomock.NewController(t)
	investmentGetAll := mock.NewMockIInvestmentGetAll(ctrl)
	restController := controllers.InvestmentsController{InvestmentsGetAll: investmentGetAll}
	defer ctrl.Finish()

	t.Run("Should get all investment", func(t *testing.T) {
		expected := []models.Investment{
			{ID: 1, Name: "investment test api 1"},
			{ID: 2, Name: "investment test api 2"},
		}

		investmentGetAll.EXPECT().Execute(gomock.Any()).Return(expected, nil)

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodGet,
			Url:    "/investments",
			Path:   "/investments",
		}, restController.GetAll)

		var result []models.Investment
		assert.NoError(t, resp.DecodeBody(&result))
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})
}
