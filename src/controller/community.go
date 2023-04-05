package controller

import (
	"fmt"
	"gin-cli/src/logic"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityTagList 获取社区所有标签
func CommunityTagList(c *gin.Context) {
	data, err := logic.GetCommunityTagList()
	if err != nil {
		zap.L().Error("logic.CommunityTagList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetail(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id, "----id")
	iId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetCommunityDetail(iId)
	if err != nil {
		ResponseCustomError(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, data)
}
