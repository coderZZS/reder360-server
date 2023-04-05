package mysql

import (
	"gin-cli/src/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func PostAdd(p *models.Post) (err error) {
	sqlStr := `insert into post( post_id, title, content, author_id, community_id, status)
				values(?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID, 0)
	return
}

func GetPostDetail(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id=?`
	err = db.Get(post, sqlStr, id)
	return
}

func GetPostList(page, pageSize int64) (data []*models.Post, err error) {
	data = make([]*models.Post, 0, pageSize)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
				from post
				limit ?,?`
	err = db.Select(&data, sqlStr, (page-1)*pageSize, pageSize)
	return
}

func GetPostListByTypeOfDetail(ids []string) (data []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
		from post
		where post_id in (?)
		order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&data, query, args...)
	return
}
