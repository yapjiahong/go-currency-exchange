package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Msg struct {
	Message string `json:"message"`
}

func Throw(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, &Msg{
		Message: err.Error(),
	})
}
