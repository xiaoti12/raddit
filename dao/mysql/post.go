package mysql

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"raddit/models"
	"strings"
)

func InsertPost(post *models.Post) error {
	sqlStr := `insert into post
	(post_id, author_id, community_id, title, content)
	values (?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.ID, post.AuthorID, post.CommunityID, post.Title, post.Content)
	if err != nil {
		return err
	}
	return nil
}

func GetPostByID(id int64) (*models.Post, error) {
	p := new(models.Post)
	sqlStr := `select post_id, author_id, community_id, title, content, status, create_time
	from post
	where post_id = ?`
	err := db.Get(p, sqlStr, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		err = ErrorInvalidID
	}
	return p, err
}

func GetPostList(page, size int) ([]*models.Post, error) {
	sqlStr := `select post_id, author_id, community_id, title, content, status, create_time
	from post
	limit ?,?`
	var posts []*models.Post
	err := db.Select(&posts, sqlStr, (page-1)*size, size)
	return posts, err
}

func GetPostListByIDs(ids []string) ([]*models.Post, error) {
	sqlStr := `select post_id, author_id, community_id, title, content, status, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id,?)
	`
	query, args, err := sqlx.In(sqlStr, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	var posts []*models.Post
	err = db.Select(&posts, query, args...)
	return posts, err
}
