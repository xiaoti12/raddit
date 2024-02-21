package models

type RegisterParams struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type VoteParams struct {
	PostID   int64 `json:"post_id,string" binding:"required"`
	Attitude int   `json:"attitude" binding:"oneof=1 0 -1"`
}
