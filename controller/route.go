package controller

import "github.com/gin-gonic/gin"

func Route(r *gin.RouterGroup) {

	r.GET("health", getHealth)
	r.GET("currency", getCurrency)
}
