package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseHandler is interface for response functionality
type ResponseHandler interface {
	SuccessResponse(ctx *gin.Context, status int, data interface{})
	ErrorResponse(ctx *gin.Context, code, details string)
	OkResponse(ctx *gin.Context, data interface{})
}

// ResponseService is an empty structure
type ResponseService struct{}

// NewResponseService creates a new ResponseService instance.
func NewResponseService() *ResponseService {
	return new(ResponseService)
}

// SuccessResponse returns a success response
func (res *ResponseService) SuccessResponse(ctx *gin.Context, status int, data interface{}) {
	r := NewResponse().SetStatus(status).SetData(data)
	ctx.JSON(status, r)
}

// OkResponse returns a "200 StatusOK" response
func (res *ResponseService) OkResponse(ctx *gin.Context, data interface{}) {
	status := http.StatusOK
	r := NewResponse().SetStatus(status).SetData(data)
	ctx.JSON(status, r)
}

// ErrorResponse returns a error response
func (res *ResponseService) ErrorResponse(ctx *gin.Context, code, details string) {
	e := NewError().SetCode(code).SetTitleByCode(code).SetTarget(TargetCommon).SetDetails(details)
	r := NewResponse().SetStatusByCode(code).AddError(e)
	ctx.AbortWithStatusJSON(r.Status, r)
}
