package user

import (
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context)  {
	c.JSON(200, gin.H{"message": "moshi-moshi"})
}
