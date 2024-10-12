package models

type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"primaryKey;comment:送信者ID" json:"send_user_id"`
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"size:42;comment:送信者のニックネーム" json:"send_user_nick_name"`
	SendUserAvatar   string    `gorm:"comment:送信者のプロフィール" json:"send_user_avatar"`

	RevUserID       uint      `gorm:"primaryKey;comment:受信者ID" json:"rev_user_id"`
	RevUserModel    UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevUserNickName string    `gorm:"size:42;comment:受信者のニックネーム" json:"rev_user_nick_name"`
	RevUserAvatar   string    `gorm:"comment:受信者のプロフィール" json:"rev_user_avatar"`
	IsRead          bool      `gorm:"default:false;comment:確認した" json:"is_read"`
	Content         string    `gorm:"comment:メッセージ内容" json:"content"`
}
