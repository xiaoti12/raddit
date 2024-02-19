package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"raddit/models"
)

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
