package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lascape/gopkg/response"
	"github.com/lascape/goweb/server/http/params"
	"github.com/lascape/goweb/service/go_web"
	"github.com/sirupsen/logrus"
)

type BankAccount struct {
	Service *go_web.Service
}

func NewGoWebController() *BankAccount {
	return &BankAccount{Service: go_web.NewService()}
}

// GoWeb 服务
// @Summary 服务
// @Description 服务
// @Tags 测试
// @Accept json
// @Produce json
// @Param req query params.ReqGoWeb true "查询参数"
// @Success 200 {object} any "返回结果"
// @Router /api/v1/go/web [get]
func (c *BankAccount) GoWeb(ctx *gin.Context) {
	req, err := params.ValidReqGoWeb(ctx)
	if err != nil {
		logrus.Error(err)
		response.Error(ctx, err)
		return
	}

	err = c.Service.GoWeb(ctx.Request.Context(), req)
	if err != nil {
		logrus.Error(err)
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, nil)
}
