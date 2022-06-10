package router

import (
	"github.com/gin-gonic/gin"
	"zzgj_network/router/api"
)

// 路由注册
func Register(router *gin.Engine) {
	api.RegisterNetwork(router)
}
