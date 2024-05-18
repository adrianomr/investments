package controllers

import (
	"github.com/adrianomr/investments/src/domain/models"
	"net/http"

	"github.com/adrianomr/investments/src/domain/usecases"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

type CdbController struct {
	Create      usecases.ICdbCreate
	CreateOrder usecases.ICdbOrderCreate
}

func NewCdbController() *CdbController {
	return &CdbController{
		Create:      usecases.NewCdbCreate(),
		CreateOrder: usecases.NewCdbOrderCreate(),
	}
}

func (c *CdbController) Routes() []restserver.Route {
	return []restserver.Route{
		{
			URI:      "/cdbs",
			Method:   http.MethodPost,
			Function: c.CreateCdb,
			Prefix:   restserver.AuthenticatedApi,
		},
		{
			URI:      "/cdbs/{cdb_id}/orders",
			Method:   http.MethodPost,
			Function: c.CreateCdbOrder,
			Prefix:   restserver.AuthenticatedApi,
		},
	}
}

// @Summary Get all
// @Tags Cdb
// @Accept json
// @Produce json
// @Success 200 {array} models.Cdb
// @Param X-UserId header uint64 true "id do usuário" minimum(0)
// @Router /cdbs/{cdb_id}/orders [get]
func (c *CdbController) CreateCdb(ctx restserver.WebContext) {
	userId := ctx.AuthenticationContext().GetUserID()
	cdb := &models.Cdb{}
	ctx.DecodeBody(cdb)
	cdb.UserID = userId
	result, err := c.Create.Execute(ctx.Context(), cdb)
	if err != nil {
		ctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	ctx.JsonResponse(http.StatusCreated, result)
}

// @Summary Get all
// @Tags Cdb
// @Accept json
// @Produce json
// @Success 200 {array} models.Cdb
// @Param X-UserId header uint64 true "id do usuário" minimum(0)
// @Router /cdbs/{cdb_id}/orders [get]
func (c *CdbController) CreateCdbOrder(ctx restserver.WebContext) {
	userId := ctx.AuthenticationContext().GetUserID()
	order := &models.CdbOrder{}
	ctx.DecodeBody(order)
	order.UserID = userId
	result, err := c.CreateOrder.Execute(ctx.Context(), order)
	if err != nil {
		ctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	ctx.JsonResponse(http.StatusCreated, result)
}