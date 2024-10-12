package models

import "MoneyBackward/models/ctype"

type LoginDataModel struct {
	MODEL
	UserID    uint             `gorm:"comment:ユーザーID" json:"user_id"`
	UserModel UserModel        `gorm:"foreignKey:UserID" json:"-"`
	IP        string           `gorm:"size:20;comment:ip" json:"ip"`
	NickName  string           `gorm:"size:42;comment:ニックネーム" json:"nick_name"`
	Token     string           `gorm:"size:256;comment:token" json:"token"`
	Device    string           `gorm:"size:256;comment:ログイン失敗" json:"device"`
	Addr      string           `gorm:"size:64;comment:Address" json:"addr"`
	LoginType ctype.SignStatus `gorm:"size:type=smallint(6);comment:ログイン方法" json:"login_type"`
}
