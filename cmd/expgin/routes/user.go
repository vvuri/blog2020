package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user_id": "someid"})

}
