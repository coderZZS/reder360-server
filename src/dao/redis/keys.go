package redis

// redis key 命名空间，方便查询

const (
	KeyPrefix          = "gincli:"
	KeyPostTimeZSet    = "post:time"   // zset;帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;帖子及投票分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;参数是post id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
