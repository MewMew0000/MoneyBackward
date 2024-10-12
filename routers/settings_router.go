package routers

import "MoneyBackward/api"

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingsSiteInfoView)
}
