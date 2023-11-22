package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthResp struct {
	ServerHealth string `json:"server_health"`
}

// getHealth godoc
// @Summary      Get API server health
// @Description  Get API server health
// @Success      200  {object}  string
// @Router       /health [get]
func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, healthResp{ServerHealth: "UP"})
}
