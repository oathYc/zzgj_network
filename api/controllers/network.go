package controllers

import (
	"errors"
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
	if params.Img == "" {
		response.Error(errors.New("上传图片不能为空"))
		return
	}
	time.Sleep(time.Second / 10) // 休眠100毫秒

	response.Success(true)
}
