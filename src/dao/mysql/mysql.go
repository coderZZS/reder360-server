package mysql

import (
	"fmt"
	"reader360Server/src/settings"

	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func Init(config *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	fmt.Println("mysql connect success")
	return nil
}

func Close() {
	db.Close()
}
