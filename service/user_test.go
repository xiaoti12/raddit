package service

import (
	"raddit/models"
	"testing"
)

func TestLogin(t *testing.T) {
	// 需要初始化配置和mysql…
	p := &models.LoginParams{
		Username: "",
		Password: "",
	}
	token, err := Login(p)
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("token is empty")
	}
	t.Log(token)
}
