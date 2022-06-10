package boot

import (
	"github.com/gin-gonic/gin"
	"zzgj_network/library/global"
	"zzgj_network/router"
)

func InitHttpServer() error {

	r := gin.Default()
	router.Register(r)

	err := r.Run(global.Config.Server.LocalHttpAddr)
	return err
}
