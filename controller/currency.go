package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExReq struct {
	Source string `form:"source" binding:"required,oneof=USD TWD JPY"`
	Target string `form:"target" binding:"required,oneof=USD TWD JPY"`
	Amount string `form:"amount" bind:"required"`
}

type ExResp struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

// getCurrency godoc
// @Summary      Currency Exchange
// @Description  Currency Exchange
// @Tags         Currency
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /currency [get]
func getCurrency(c *gin.Context) {
	exReq := ExReq{}
	err := c.ShouldBind(&exReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, ExResp{
			Msg:    "failed",
			Amount: "0",
		})
		return
	}

	c.JSON(http.StatusOK, ExResp{
		Msg:    "success",
		Amount: "101",
	})
}
