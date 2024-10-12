package models

import "MoneyBackward/models/ctype"

type BannerModel struct {
	MODEL
	Path        string            `gorm:"comment:パス" json:"path"`
	Hash        string            `gorm:"comment:ハッシュ値" json:"hash"`
	Name        string            `gorm:"size:38;comment:画像の名" json:"name"`
	ImageType   ctype.ImageType   `gorm:"default:1;comment:タイプ" json:"image_type"`
	MenusBanner []MenuBannerModel `gorm:"foreignKey:BannerID" json:"-"`
}
