package mysql

import (
	"database/sql"
	"fmt"
	"gin-cli/src/models"

	"go.uber.org/zap"
)

func GetCommunityTagList() (data []*models.Community, err error) {
	var sqlStr = "select community_id, community_name from community"
	if err := db.Select(&data, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in d")
			err = nil
		}
	}
	return
}

func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	data = new(models.CommunityDetail)
	var sqlStr = `select community_id, community_name, introduction, create_time from community where community_id = ?`
	err = db.Get(data, sqlStr, id)
	fmt.Println(err, "eer")
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		fmt.Println(err, "eer")
		return
	}
	return
}
