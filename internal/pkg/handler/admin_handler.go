package handler

import (
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/pkg/infrastucture/db"
	"be_soc/pkg/infrastucture/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type HTTPAdminHandler struct {
	Schema *graphql.Schema
}

func NewHTTPAdminHandler(db db.Database) *HTTPHandler {

	schema := schema.NewAdminSchema(db)
	return &HTTPHandler{Schema: schema}
}

func (h *HTTPAdminHandler) Handle(c *gin.Context) {
	req := dto.BaseRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         *h.Schema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
	})
	code := http.StatusOK
	if len(data.Errors) > 0 {
		code = http.StatusBadRequest
	}
	resp := dto.BaseResponse{
		Status: code,
		Error:  data.Errors,
		Result: data.Data,
	}

	c.JSON(code, resp)
}
