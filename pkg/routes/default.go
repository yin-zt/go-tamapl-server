package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-tamapl-server/pkg/config"
)

func InitRoutes() *gin.Engine {
	gin.SetMode(config.GinMode)
	r := gin.Default()
	apiGroup := r.Group("/" + config.GinUrlPrefix)

	InitCronRoutes(apiGroup)
	return r
}
