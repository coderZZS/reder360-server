package logic

import (
	"gin-cli/src/dao/mysql"
	"gin-cli/src/models"
)

func GetCommunityTagList() (data []*models.Community, err error) {
	return mysql.GetCommunityTagList()
}

func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetail(id)
}
