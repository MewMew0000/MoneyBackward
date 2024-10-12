package models

type CommentModel struct {
	MODEL              `json:",select(c)"`
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments,select(c)"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentID    *uint           `gorm:"comment:親コメントID" json:"parent_comment_id,select(c)"`
	Content            string          `gorm:"size:256;comment:コメント内容" json:"content,select(c)"`
	DiggCount          int             `gorm:"size:8;default:0;comment:いいね数" json:"digg_count,select(c)"`
	CommentCount       int             `gorm:"size:8;default:0;comment:子コメント数" json:"comment_count,select(c)"`
	ArticleID          string          `gorm:"size:32;comment:文章ID" json:"article_id,select(c)"`
	User               UserModel       `gorm:"comment:関連ユーザー" json:"user,select(c)"`
	UserID             uint            `gorm:"comment:未関連ユーザー" json:"user_id,select(c)"`
}
