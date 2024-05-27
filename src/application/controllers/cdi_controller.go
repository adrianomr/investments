package controllers

import (
	"github.com/adrianomr/investments/src/domain/models"
	"net/http"

	"github.com/adrianomr/investments/src/domain/usecases"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

type CdiController struct {
	Create usecases.ICdiCreate
}

func NewCdiController() *CdiController {
	return &CdiController{
		Create: usecases.NewCdiCreate(),
	}
}

func (c *CdiController) Routes() []restserver.Route {
	return []restserver.Route{
		{
			URI:      "/cdis",
			Method:   http.MethodPost,
			Function: c.CreateCdi,
			Prefix:   restserver.AuthenticatedApi,
		},
	}
}

// @Summary Get all
// @Tags Cdi
// @Accept json
// @Produce json
// @Success 200 {array} models.Cdi
// @Param X-UserId header uint64 true "id do usuário" minimum(0)
// @Router /cdis/{cdi_id}/orders [get]
func (c *CdiController) CreateCdi(ctx restserver.WebContext) {

	cdi := &models.Cdi{}
	err := ctx.DecodeBody(cdi)
	if err != nil {
		ctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	result, err := c.Create.Execute(ctx.Context(), cdi)

	if err != nil {
		ctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	ctx.JsonResponse(http.StatusCreated, result)
}
