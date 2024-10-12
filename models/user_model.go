package models

import "MoneyBackward/models/ctype"

type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36;comment:ニックネーム" json:"nick_name,select(c|info)"`
	UserName   string           `gorm:"size:36;comment:ユーザー名" json:"user_name"`
	Password   string           `gorm:"size:128;comment:パスワード" json:"-"`
	Avatar     string           `gorm:"size:256;comment:プロフィール画像" json:"avatar,select(c)"`
	Email      string           `gorm:"size:128;comment:メールアドレス" json:"email,select(info)"`
	Tel        string           `gorm:"size:18;comment:電話番号" json:"tel"`
	Addr       string           `gorm:"size:64;comment:Address" json:"addr,select(c|info)"`
	Token      string           `gorm:"size:64;comment:他のプラットフォームのユニークID" json:"token"`
	IP         string           `gorm:"size:20;comment:ip" json:"ip,select(c)"`
	Role       ctype.Role       `gorm:"size:4;default:1;comment:ロール" json:"role,select(info)"`
	SignStatus ctype.SignStatus `gorm:"type=smallint(6);comment:登録元" json:"sign_status,select(info)"`
	Integral   int              `gorm:"default:0;comment:ポイント" json:"integral,select(info)"`
	Scope      int              `gorm:"default:0;comment:ポイント" json:"scope,select(info)"`
	Sign       string           `gorm:"size:128;comment:サイン" json:"sign,select(info)"`
	Link       string           `gorm:"size:128;comment:link" json:"link,select(info)"`
}
