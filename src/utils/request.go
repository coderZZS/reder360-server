package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

var CtxUserIDKey = "userID"

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	fmt.Println(uid, "--------------------udddd")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return userID, nil
}

func GetPageParams(c *gin.Context) (page, pageSize int64, err error) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")
	page, _ = strconv.ParseInt(pageStr, 10, 64)
	pageSize, _ = strconv.ParseInt(pageSizeStr, 10, 64)
	err = nil
	if page < 0 || pageSize < 0 {
		err = errors.New("页码参数有误")
	}
	return
}
