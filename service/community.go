package service

import (
	"raddit/dao/mysql"
	"raddit/models"
)

func GetCommunityList() ([]*models.CommunityBasic, error) {
	return mysql.GetCommunityList()
}
