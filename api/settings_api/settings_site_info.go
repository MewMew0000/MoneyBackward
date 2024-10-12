package settings_api

import (
	"MoneyBackward/models/res"

	"github.com/gin-gonic/gin"

	"MoneyBackward/global"
)

func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	res.OkWithData(global.Conf.SiteInfo, c)
}
