package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

func CreateArticle(conn *gorm.DB) {
	conn.AutoMigrate(&entity.Article{})
}

func CreateArticleLine(conn *gorm.DB) {
	conn.AutoMigrate(&entity.ArticleLine{})
}

func CreateBoldNumber(conn *gorm.DB) {
	conn.AutoMigrate(&entity.BoldNumber{})
}

func CreateLinkNumber(conn *gorm.DB) {
	conn.AutoMigrate(&entity.LinkNumber{})
}
