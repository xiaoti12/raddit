package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"raddit/models"
)

var ErrorInvalidID = errors.New("invalid query id")

func GetCommunityList() ([]*models.CommunityBasic, error) {
	var communityList []*models.CommunityBasic
	sqlStr := `select community_id,community_name from community`
	err := db.Select(&communityList, sqlStr)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("no data in community table")
		err = nil
	}
	return communityList, err
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	var commDetail = new(models.CommunityDetail)
	// cannot use `select *` here
	sqlStr := `select 
    	community_id, community_name, introduction, create_time, update_time
		from community where community_id = ?`
	err := db.Get(commDetail, sqlStr, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		err = ErrorInvalidID
	}
	return commDetail, err
}

func GetCommunityBasic(id int64) (*models.CommunityBasic, error) {
	var communityBasic = new(models.CommunityBasic)
	sqlStr := `select community_id, community_name
		from community
		where community_id = ?`
	err := db.Get(communityBasic, sqlStr, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		err = ErrorInvalidID
	}
	return communityBasic, err
}
