package service

import (
	"raddit/dao/mysql"
	"raddit/models"
	"raddit/pkg/jwt"
	"raddit/pkg/snowflake"
)

func Register(p *models.RegisterParams) error {
	// check if user already exists
	err := mysql.CheckUserExists(p.Username)
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

func Login(p *models.LoginParams) (string, error) {
	// check if user exists
	err := mysql.CheckUserNotExists(p.Username)
	if err != nil {
		return "", err
	}
	// check if password is correct
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.CheckUserLogin(user)
	if err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}
