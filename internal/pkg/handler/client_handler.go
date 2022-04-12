package handler

import (
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/pkg/infrastucture/db"
	"be_soc/pkg/infrastucture/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type HTTPClientHandler struct {
	Schema *graphql.Schema
}

func NewHTTPClientHandler(db db.Database) *HTTPHandler {

	schema := schema.NewClientSchema(db)
	return &HTTPHandler{Schema: schema}
}

func (h *HTTPClientHandler) Handle(c *gin.Context) {
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
	//fmt.Println(req.Query)
	//fmt.Println(req.Mutation)
	exce := ""
	if len(req.Query) > 0 {
		exce = req.Query
	} else {
		exce = req.Mutation
	}
	//fmt.Println(exce)
	data := graphql.Do(graphql.Params{
		Schema:        *h.Schema,
		RequestString: exce,
		Context:       c,
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
	c.JSON(http.StatusOK, resp)
}
