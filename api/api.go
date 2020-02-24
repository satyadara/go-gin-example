package api

import (
	"github.com/gin-gonic/gin"
	"satya-labs/api/user"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		user.ApplyRoutes(api)
	}
}
