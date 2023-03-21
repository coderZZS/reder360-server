package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gin-cli/src/models"
)

const SECRET = "HELLO WORLD"

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

func InsertUser(user *models.UserModel) (err error) {
	// 加密
	user.Password = encryPassword(user.Password)
	// 插入
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return nil
}

func encryPassword(pwd string) string {
	h := md5.New()
	h.Write([]byte(SECRET))
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}
