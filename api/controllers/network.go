package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
	"zzgj_network/library"
)

type NetworkControl struct{}

type NetworkPostReq struct {
	Img string `json:"img"`
}

type NetworkResp struct {
}

func (*NetworkControl) ImgUpload(ctx *gin.Context) {
	params := new(NetworkPostReq)

	response := library.NewResponse(ctx)
	if err := library.NewRequest(ctx).Bind(params); nil != err {
		response.Error(err)
		return
	}
	time.Sleep(time.Second / 10) // 休眠100毫秒

	response.Success(true)
}
