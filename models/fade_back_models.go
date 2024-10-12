package models

type FadeBackModel struct {
	MODEL
	Email        string `gorm:"size:64;comment:メールアドレス" json:"email"`
	Content      string `gorm:"size:128;comment:内容" json:"content"`
	ApplyContent string `gorm:"size:128;comment:返信内容" json:"apply_content"`
	IsApply      bool   `gorm:"comment:返信した" json:"is_apply"`
}
