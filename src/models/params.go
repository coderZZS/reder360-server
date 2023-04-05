package models

// UserSignUpParams 注册接口参数
type UserSignUpParams struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// UserLoginParams 登录接口参数
type UserLoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// VotePostParams 帖子投票参数
type VotePostParams struct {
	//UserID // 从token解析即可
	PostID    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction" binding:"oneof=1 0 -1"`
}

type ParamsPostList struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"page_size" form:"page_size"`
	Order    string `json:"order" form:"order"`
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)
