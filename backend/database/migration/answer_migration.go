package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

func CreateAnswer(conn *gorm.DB) {
	conn.AutoMigrate(&entity.Answer{})
}
