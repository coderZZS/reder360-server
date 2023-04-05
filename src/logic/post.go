package logic

import (
	"fmt"
	"gin-cli/src/dao/mysql"
	"gin-cli/src/dao/redis"
	"gin-cli/src/models"
	"gin-cli/src/pkg/snowflake"

	"go.uber.org/zap"
)

func PostAdd(p *models.Post) (err error) {
	p.ID = int64(snowflake.GenID())
	err = mysql.PostAdd(p)
	if err != nil {
		return
	}
	return redis.VoteCreatePost(p.ID)
}

func PostDetail(id int64) (data *models.PostDetail, err error) {
	post, err := mysql.GetPostDetail(id)
	if err != nil {
		zap.L().Error("post detail failed", zap.Error(err))
		return
	}
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("post detail failed", zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("post detail failed", zap.Error(err))
		return
	}
	data = &models.PostDetail{
		AuthorName:      user.Username,
		CommunityDetail: community,
		Post:            post,
	}
	return
}

func PostList(page, pageSize int64) (data []*models.PostDetail, err error) {
	posts, err := mysql.GetPostList(page, pageSize)
	if err != nil {
		return nil, err
	}
	data = make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("post detail failed", zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			fmt.Println(post.CommunityID, "post.CommunityID")
			zap.L().Error("post detail failed", zap.Error(err))
			return nil, err
		}
		postDetail := &models.PostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}

	return data, nil
}

func PostListByType(p *models.ParamsPostList) (data []*models.PostDetail, err error) {
	ids, err := redis.GetPostListByTypeOfIds(p)
	if err != nil || len(ids) == 0 {
		return nil, err
	}
	// 根据ids切片查询详细数据
	posts, err := mysql.GetPostListByTypeOfDetail(ids)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("post detail failed", zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			fmt.Println(post.CommunityID, "post.CommunityID")
			zap.L().Error("post detail failed", zap.Error(err))
			return nil, err
		}
		postDetail := &models.PostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
