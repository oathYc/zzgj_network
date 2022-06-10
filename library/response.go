package library

import (
	"github.com/gin-gonic/gin"
	"net/http"
	api_http "zzgj_network/api"
)

type Response struct {
	ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	traceId := "<empty>"
	medCtx := api_http.ParserMedContext(c)
	if nil != medCtx {
		traceId = medCtx.GetTraceId()
	}

	// 响应请求唯一标识
	c.Header("X-Trace-Id", traceId)

	return &Response{c}
}

// 响应成功请求
func (r *Response) Success(data interface{}) {
	r.SuccessMsg(data, "")
}

// 响应成功请求 自定义msg
func (r *Response) SuccessMsg(data interface{}, msg string) {
	if len(msg) == 0 {
		msg = "success"
	}

	// 响应数据
	r.ctx.JSON(http.StatusOK, api_http.ResponseEntity{
		Code:    Success,
		Message: msg,
		Data:    data,
	})
}

type ErrorInit struct {
	ErrorCode int
	Error     error
}

// 响应错误
func (r *Response) Error(err error) {
	// 屏蔽掉其他错误
	r.ctx.JSON(http.StatusOK, api_http.ResponseEntity{
		Code:    ErrInternal,
		Message: err.Error(),
		Data:    nil,
	})
	return
}
