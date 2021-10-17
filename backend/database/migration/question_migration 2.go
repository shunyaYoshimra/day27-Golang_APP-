package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

func CreateQuestion(conn *gorm.DB) {
	conn.AutoMigrate(&entity.Question{})
}
