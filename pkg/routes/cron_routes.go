package routes

import "github.com/gin-gonic/gin"

func InitCronRoutes(r *gin.RouterGroup) gin.IRoutes {
	cron := r.Group("/cron")
	{
		cron.GET("/list")
	}

	return r
}
