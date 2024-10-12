package models

type MenuBannerModel struct {
	MenuID      uint        `gorm:"comment:メニューのID" json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `gorm:"comment:bannerID" json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10;comment:番号" json:"sort"`
}
