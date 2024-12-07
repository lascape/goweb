package params

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// ReqGoWeb 请求参数，用于获取银行账户列表
type ReqGoWeb struct {
	ID int64 `json:"id" form:"id"` // 账户 ID
}

func ValidReqGoWeb(ctx *gin.Context) (req ReqGoWeb, err error) {
	err = ctx.ShouldBind(&req)
	if err != nil {
		return req, err
	}
	if req.ID == 0 {
		return req, errors.New("invalid id")
	}
	return req, nil
}
