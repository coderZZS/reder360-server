package controller

import (
	"fmt"
	"gin-cli/src/logic"
	"gin-cli/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUp(c *gin.Context) {
	// 获取参数
	fmt.Println("3333")
	var p models.UserSignUpParams
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 验证通过执行逻辑
	if err := logic.SignUp(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 业务处理
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
