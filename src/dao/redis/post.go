package redis

import "gin-cli/src/models"

func GetPostListByTypeOfIds(p *models.ParamsPostList) ([]string, error) {
	// 从redis里面取出id列表
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.确定查询的索引起始点
	start := (p.Page - 1) * p.PageSize
	end := start + p.PageSize - 1

	// 3.ZREVRANGE查询
	return rdb.ZRevRange(key, start, end).Result()
}
