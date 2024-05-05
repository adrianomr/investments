package controllers

import (
	"net/http"

	"github.com/adrianomr/investments/src/domain/usecases"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

type InvestmentsController struct {
	InvestmentsGetAll usecases.IInvestmentGetAll
}

func NewInvestmentsController() *InvestmentsController {
	return &InvestmentsController{
		InvestmentsGetAll: usecases.NewInvestmentGetAll(),
	}
}

func (c *InvestmentsController) Routes() []restserver.Route {
	return []restserver.Route{
		{
			URI:      "investments",
			Method:   http.MethodGet,
			Function: c.GetAll,
			Prefix:   restserver.PublicApi,
		},
	}
}

// @Summary Get all
// @Tags Investments
// @Accept json
// @Produce json
// @Success 200 {array} models.Investments
// @Param X-AccountantId header uint64 true "id da contabilidade" minimum(0)
// @Param X-TenantId header uint64 true "id do dono do negócio" minimum(0)
// @Param X-UserId header uint64 true "id do usuário" minimum(0)
// @Router /private-api/rest/investments [get]
func (c *InvestmentsController) GetAll(ctx restserver.WebContext) {
	result, err := c.InvestmentsGetAll.Execute(ctx.Context())
	if err != nil {
		ctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	ctx.JsonResponse(http.StatusOK, result)
}
