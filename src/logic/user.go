package logic

import (
	"gin-cli/src/dao/mysql"
	"gin-cli/src/models"
	"gin-cli/src/pkg/snowflake"
)

func SignUp(p *models.UserSignUpParams) (err error) {
	// 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成UID
	userID := snowflake.GenID()
	// 创建user实例
	user := models.UserModel{
		Username: p.Username,
		Password: p.Password,
		UserID:   userID,
	}
	// 执行插入逻辑
	err = mysql.InsertUser(&user)
	return err
}
