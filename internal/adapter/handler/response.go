package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// response represents a response body format
type Response struct {
	Error   int    `json:"error" example:"0"`
	Message string `json:"message" example:"Message"`
	Data    any    `json:"data,omitempty"`
	Count   int    `json:"count,omitempty"`
	Page    int    `json:"page,omitempty"`
	Size    int    `json:"size,omitempty"`
}

type SuccessOptions struct {
	Count int
	Page  int
	Size  int
}

type SuccessOption func(req *SuccessOptions)

func WithPagination(count, page, size int) SuccessOption {
	return func(req *SuccessOptions) {
		req.Count = count
		req.Page = page
		req.Size = size
	}
}

func SuccessResponse(ctx *gin.Context, data any, opts ...SuccessOption) {
	res := &SuccessOptions{}
	for _, opt := range opts {
		opt(res)
	}
	ctx.JSON(http.StatusOK, Response{
		Error:   0,
		Message: "Success",
		Data:    data,
		Page:    res.Page,
		Count:   res.Count,
		Size:    res.Size,
	})
}

func ErrorResponse(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, Response{
		Error:   code,
		Message: err.Error(),
	})
}
