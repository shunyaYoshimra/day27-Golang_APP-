package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type ProfileRepository struct {
	Conn *gorm.DB
}

func NewProfileRepository() ProfileRepository {
	return ProfileRepository{Conn: database.GetDB().Table("profiles")}
}

func (pr *ProfileRepository) RetrieveProfiles() (profiles []entity.Profile) {
	pr.Conn.Find(&profiles)
	return
}

func (pr *ProfileRepository) FindByUser(id int) (profile entity.Profile, err error) {
	err = pr.Conn.Where("user_id = ?", id).First(&profile).Error
	return
}

func (pr *ProfileRepository) Create(profile *entity.Profile) (err error) {
	err = pr.Conn.Create(profile).Error
	return
}

func (pr *ProfileRepository) Update(profile entity.Profile, description, subject string) (err error) {
	err = pr.Conn.Model(&profile).Update(entity.Profile{Description: description, Subject: subject}).Error
	return
}
