package models

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primaryKey;comment:id" json:"id,select($any)" structs:"-"`
	CreatedAt time.Time `gorm:"comment:createTime" json:"created_at,select($any)" structs:"-"`
	UpdatedAt time.Time `gorm:"comment:updateTime" json:"-" structs:"-"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}
type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type Options[T any] struct {
	Label string `json:"label"`
	Value T      `json:"value"`
}

const (
	AdminRole   = 1 // admin
	UserRole    = 2 // user
	TouristRole = 3 // tourist
)
