package models

import "time"

type UserCollectModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"comment:ユーザーID" json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32;comment:文章のESID"`
	CreatedAt time.Time `gorm:"comment:お気に入りにした時点"`
}
