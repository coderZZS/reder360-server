package controller

import (
	"gin-cli/src/logic"
	"gin-cli/src/models"
	"gin-cli/src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PostAdd 创建帖子
func PostAdd(c *gin.Context) {
	// 参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	uid, err := utils.GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = uid
	if err := logic.PostAdd(p); err != nil {
		zap.L().Error("logic.PostAdd(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func PostDetail(c *gin.Context) {
	id := c.Param("id")
	iId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.PostDetail(iId)
	if err != nil {
		zap.L().Error("PostDetail failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	ResponseSuccess(c, data)
}

func PostList(c *gin.Context) {
	page, pageSize, err := utils.GetPageParams(c)
	if err != nil {
		ResponseCustomError(c, CodeInvalidParams, err.Error())
		return
	}
	data, err := logic.PostList(page, pageSize)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// PostListByType 根据时间或者分数查询列表
// @Summary 根据时间或者分数查询列表
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamsPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /api/post/list/type [get]
func PostListByType(c *gin.Context) {
	p := &models.ParamsPostList{
		Page:     1,
		PageSize: 10,
		Order:    models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 获取数据
	data, err := logic.PostListByType(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
