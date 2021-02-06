package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type ContactRepository struct {
	Conn *gorm.DB
}

func NewContactRepository() ContactRepository {
	return ContactRepository{Conn: database.GetDB().Table("contacts")}
}

func (cr *ContactRepository) FindByUser(id int) (contact entity.Contact, err error) {
	err = cr.Conn.Where("user_id = ?", id).First(&contact).Error
	return
}

func (cr *ContactRepository) Create(contact *entity.Contact) (err error) {
	err = cr.Conn.Create(contact).Error
	return
}

func (cr *ContactRepository) Update(contact entity.Contact, media, account string) (err error) {
	err = cr.Conn.Model(&contact).Update(entity.Contact{Media: media, Account: account}).Error
	return
}
