package user

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.RouterGroup)  {
	user := r.Group("/user")
	{
		user.GET("/", create)
	}
}
