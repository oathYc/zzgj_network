package api

import (
	"github.com/gin-gonic/gin"
	"zzgj_network/api/controllers"
)

func RegisterNetwork(router *gin.Engine) {
	api := new(controllers.NetworkControl)
	v1 := router.Group("/api/v1/network")
	{
		v1.POST("/img-upload", api.ImgUpload)
	}
}
