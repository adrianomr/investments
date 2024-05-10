package controllers

import (
	"github.com/adrianomr/investments/src/domain/models"
	"net/http"

	"github.com/adrianomr/investments/src/domain/usecases"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

type CdbController struct {
	Create usecases.ICdbOrderCreate
}

func NewCdbController() *CdbController {
	return &CdbController{
		Create: usecases.NewCdbOrderCreate(),
	}
}

func (c *CdbController) Routes() []restserver.Route {
	return []restserver.Route{
		{
			URI:      "cdb",
			Method:   http.MethodGet,
			Function: c.CreateCdb,
			Prefix:   restserver.PublicApi,
		},
	}
}

// @Summary Get all
// @Tags Cdb
// @Accept json
// @Produce json
// @Success 200 {array} models.Cdb
// @Param X-AccountantId header uint64 true "id da contabilidade" minimum(0)
// @Param X-TenantId header uint64 true "id do dono do negócio" minimum(0)
// @Param X-UserId header uint64 true "id do usuário" minimum(0)
// @Router /private-api/rest/cdb [get]
func (c *CdbController) CreateCdb(ctx restserver.WebContext) {
	userId := ctx.AuthenticationContext().GetUserID()
	order := &models.CdbOrder{}
	ctx.DecodeBody(order)
	order.UserID = userId
	result, err := c.Create.Execute(ctx.Context(), order)
	if err != nil {
		ctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	ctx.JsonResponse(http.StatusOK, result)
}
