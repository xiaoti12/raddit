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
