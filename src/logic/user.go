package logic

import (
	"fmt"
	"gin-cli/src/dao/mysql"
	"gin-cli/src/models"
	"gin-cli/src/pkg/jwt"
	"gin-cli/src/pkg/snowflake"
)

// SignUp 注册
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

func Login(p *models.UserLoginParams) (token string, err error) {
	user := &models.UserModel{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	fmt.Println(user.UserID, "GenToken")
	token, err = jwt.GenToken(user.UserID, user.Username)
	return
}
