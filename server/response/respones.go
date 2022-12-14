package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Respones struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Respones{200, message, data})
}

func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Respones{400, message, 0})
}
