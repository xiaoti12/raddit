package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"raddit/models"
)

var salt = "xiaoti"

func InsertUser(user *models.User) error {
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	encPwd := encryptPassword(user.Password)
	_, err := db.Exec(sqlStr, user.UserID, user.Username, encPwd)
	if err != nil {
		return err
	}
	return nil

}
func CheckUserExsits(username string) error {
	var count int
	sqlStr := "select count(*) from user where username = ?"
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user already exsits")
	}
	return nil
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
