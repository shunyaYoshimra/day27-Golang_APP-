package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type AbroadRepository struct {
	Conn *gorm.DB
}

func NewAbroadRepository() AbroadRepository {
	return AbroadRepository{Conn: database.GetDB().Table("abroads")}
}

func (ar *AbroadRepository) RetrieveAbroads() (abroads []entity.Abroad) {
	ar.Conn.Find(&abroads)
	return
}

func (ar *AbroadRepository) FindByUser(id int) (abroad entity.Abroad, err error) {
	err = ar.Conn.Where("user_id = ?", id).First(&abroad).Error
	return
}

func (ar *AbroadRepository) Create(abroad *entity.Abroad) (err error) {
	err = ar.Conn.Create(abroad).Error
	return
}

func (ar *AbroadRepository) Update(abroad entity.Abroad, country, college, description string) (err error) {
	err = ar.Conn.Model(&abroad).Update(entity.Abroad{Country: country, College: college, Description: description}).Error
	return
}
