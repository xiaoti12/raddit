package service

import (
	"raddit/dao/mysql"
	"raddit/models"
	"raddit/pkg/snowflake"
)

func Register(p *models.RegisterParams) error {
	// check if user already exsits
	err := mysql.CheckUserExsits(p.Username)
	if err != nil {
		return err
	}
	// generate user id
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// insert user info to database
	err = mysql.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}
