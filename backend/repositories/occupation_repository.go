package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type OccupationRepository struct {
	Conn *gorm.DB
}

func NewOccupationRepository() OccupationRepository {
	return OccupationRepository{Conn: database.GetDB().Table("occupations")}
}

func (or *OccupationRepository) RetrieveOccuaptions() (occuaptions []entity.Occupation) {
	or.Conn.Find(&occuaptions)
	return
}

func (or *OccupationRepository) FindByUser(id int) (occupation entity.Occupation, err error) {
	err = or.Conn.Where("user_id = ?", id).First(&occupation).Error
	return
}

func (or *OccupationRepository) Create(occupation *entity.Occupation) (err error) {
	err = or.Conn.Create(occupation).Error
	return
}

func (or *OccupationRepository) Update(occupation entity.Occupation, kind, company, description string) (err error) {
	err = or.Conn.Model(&occupation).Update(entity.Occupation{Kind: kind, Company: company, Description: description}).Error
	return
}
