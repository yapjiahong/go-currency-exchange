package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-currency-converter/service"
	"net/http"
	"strconv"
	"strings"
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

func (e ExReq) parseAmount() (*service.CurrencyDTO, error) {
	amountStr := strings.ReplaceAll(e.Amount, "$", "")
	amountStr = strings.ReplaceAll(amountStr, ",", "")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %s", e.Amount)
	}
	return &service.CurrencyDTO{
		From:   service.EnumConverter[e.Source],
		To:     service.EnumConverter[e.Target],
		Amount: amount,
	}, nil
}

// getCurrency godoc
// @Summary      Currency ConvertCurrency
// @Description  Currency ConvertCurrency
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
			Msg:    err.Error(),
			Amount: "0",
		})
		return
	}

	dto, err := exReq.parseAmount()
	if err != nil {
		c.JSON(http.StatusBadRequest, ExResp{
			Msg:    err.Error(),
			Amount: "0",
		})
		return
	}

	res, err := service.ConvertCurrency(*dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, ExResp{
			Msg:    err.Error(),
			Amount: "0",
		})
		return
	}

	c.JSON(http.StatusOK, ExResp{
		Msg:    "success",
		Amount: res,
	})
}
