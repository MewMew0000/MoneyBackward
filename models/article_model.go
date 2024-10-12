package models

import (
	"MoneyBackward/models/ctype"
)

type ESIndexInterFace interface {
	Index() string
	Mapping() string
}

type ArticleModel struct {
	ID        string `json:"id" structs:"id"` // ElasticsearchのID
	CreatedAt string `json:"created_at" structs:"created_at"`
	UpdatedAt string `json:"updated_at" structs:"updated_at"`

	Title    string `json:"title" structs:"title"`
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"`
	Abstract string `json:"abstract" structs:"abstract"`
	Content  string `json:"content,omit(list)" structs:"content"`

	LookCount     int `json:"look_count" structs:"look_count"`         // 閲覧数
	CommentCount  int `json:"comment_count" structs:"comment_count"`   // コメント数
	DiggCount     int `json:"digg_count" structs:"digg_count"`         // いいね数
	CollectsCount int `json:"collects_count" structs:"collects_count"` // お気に入り数

	UserID       uint   `json:"user_id" structs:"user_id"`
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"`
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`

	Category string `json:"category" structs:"category"` // 分類
	Source   string `json:"source" structs:"source"`     // 出所
	Link     string `json:"link" structs:"link"`         // 原文リンク

	BannerID  uint   `json:"banner_id" structs:"banner_id"`   // カバーID
	BannerUrl string `json:"banner_url" structs:"banner_url"` // カバー

	Tags ctype.Array `json:"tags" structs:"tags"` // タグ
}
