package flags

import (
	"MoneyBackward/global"
	"MoneyBackward/models"
)

func DB() {
	var err error
	global.MysqlDB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.MysqlDB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	err = global.MysqlDB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Logger.Error("sql table generating failed")
		return
	}
	global.Logger.Info("sql table generating success")
}
