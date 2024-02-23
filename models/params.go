package models

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

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
	UserID   int64 `json:"user_id,string"`
	Attitude int   `json:"attitude" binding:"oneof=1 0 -1"`
}

type PostListParams struct {
	Page      int    `json:"page" form:"page"`
	Size      int    `json:"size" form:"size"`
	OrderType string `json:"order" form:"order"`
}
