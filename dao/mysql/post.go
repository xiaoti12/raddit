package mysql

import "raddit/models"

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
	if err != nil {
		return nil, err
	}
	return p, nil
}
