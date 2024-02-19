package service

import (
	"raddit/dao/mysql"
	"raddit/models"
)

func GetCommunityList() ([]*models.CommunityBasic, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetail(id)
}
