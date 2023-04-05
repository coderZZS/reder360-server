package logic

import (
	"gin-cli/src/dao/redis"
	"gin-cli/src/models"
	"strconv"
)

func VotePost(userID int64, p *models.VotePostParams) error {
	return redis.VotePost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Direction))
}
