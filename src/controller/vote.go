package controller

import (
	"gin-cli/src/logic"
	"gin-cli/src/models"
	"gin-cli/src/utils"

	"github.com/gin-gonic/gin"
)

func VotePost(c *gin.Context) {
	p := new(models.VotePostParams)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	userID, err := utils.GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeInvalidAuth)
		return
	}
	err = logic.VotePost(userID, p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
