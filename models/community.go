package models

import "time"

type CommunityBasic struct {
	ID   int64  `json:"community_id" db:"community_id"`
	Name string `json:"community_name" db:"community_name"`
}

type CommunityDetail struct {
	*CommunityBasic
	Introduction string    `json:"introduction" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
}
