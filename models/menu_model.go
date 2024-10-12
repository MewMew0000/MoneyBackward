package models

import "MoneyBackward/models/ctype"

type MenuModel struct {
	MODEL
	Title        string        `gorm:"size:32;comment:タイトル" json:"title"`
	Path         string        `gorm:"size:32;comment:パス" json:"path"`
	Slogan       string        `gorm:"size:64;comment:slogan" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string;comment:概要" json:"abstract"`
	AbstractTime int           `gorm:"comment:概要切り替え時間" json:"abstract_time"`
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"`
	BannerTime   int           `gorm:"comment:banner画像の切り替え時間" json:"banner_time"`
	Sort         int           `gorm:"size:10;comment:番号" json:"sort"`
}
