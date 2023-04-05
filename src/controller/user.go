package controller

import (
	"gin-cli/src/logic"
	"gin-cli/src/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// UserSignUp 注册
func UserSignUp(c *gin.Context) {
	// 获取参数
	var p models.UserSignUpParams
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("缺少必要的参数", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 验证通过执行逻辑
	if err := logic.SignUp(&p); err != nil {
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 业务处理
	ResponseSuccess(c, nil)
}

// UserLogin 登录
func UserLogin(c *gin.Context) {
	var p models.UserLoginParams
	// 参数校验
	if err := c.ShouldBindJSON(&p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 逻辑处理
	token, err := logic.Login(&p)
	if err != nil {
		ResponseCustomError(c, CodeInvalidParams, err.Error())
		return
	}
	ResponseSuccess(c, token)
}
