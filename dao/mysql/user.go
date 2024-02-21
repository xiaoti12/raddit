package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"raddit/models"
)

var salt = "xiaoti"

var (
	ErrorUserExist       = errors.New("user already exists")
	ErrorUserNotExist    = errors.New("user not exists")
	ErrorInvalidPassword = errors.New("invalid password")
)

func InsertUser(user *models.User) error {
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	encPwd := encryptPassword(user.Password)
	_, err := db.Exec(sqlStr, user.UserID, user.Username, encPwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUsernameByID(id int64) (string, error) {
	var username string
	sqlStr := "select username from user where user_id = ?"
	err := db.Get(&username, sqlStr, id)
	if err != nil {
		return "", err
	}
	return username, nil
}

func CheckUserExists(username string) error {
	var count int
	sqlStr := "select count(*) from user where username = ?"
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func CheckUserNotExists(username string) error {
	var count int
	sqlStr := "select count(*) from user where username = ?"
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrorUserNotExist
	}
	return nil
}

func CheckUserLogin(user *models.User) error {
	// 原始密码加密
	encPwd := encryptPassword(user.Password)
	sqlStr := "select user_id,username,password from user where username = ?"
	err := db.Get(user, sqlStr, user.Username)
	if err != nil {
		return err
	}
	// 将原始加密密码和数据库加密密码比较
	if encPwd != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
