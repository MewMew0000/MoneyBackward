package models

type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32;comment:広告のタイトル" json:"title"`
	Href   string `gorm:"comment:リンクのジャンプ" json:"href"`
	Images string `gorm:"comment:画像" json:"images"`
	IsShow bool   `gorm:"comment:表示する" json:"is_show"`
}
