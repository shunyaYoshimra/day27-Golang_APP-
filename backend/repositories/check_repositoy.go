package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type CheckRepository struct {
	Conn *gorm.DB
}

func NewCheckRepository() CheckRepository {
	return CheckRepository{Conn: database.GetDB().Table("checks")}
}

func (cr *CheckRepository) ChecksOfArticle(id int) (checks []entity.Check) {
	cr.Conn.Where("article_id = ?", id).Find(&checks)
	return
}

func (cr *CheckRepository) ChecksOfUser(id int) (checks []entity.Check) {
	cr.Conn.Where("user_id = ?", id).Find(&checks)
	return
}

func (cr *CheckRepository) CheckUnique(articleID, userID int) (check entity.Check, err error) {
	err = cr.Conn.Where("article_id = ? AND user_id = ?", articleID, userID).First(&check).Error
	return
}

func (cr *CheckRepository) Create(article *entity.Article) (err error) {
	err = cr.Conn.Create(article).Error
	return
}

func (cr *CheckRepository) Delete(article entity.Article) (err error) {
	err = cr.Conn.Delete(&article).Error
	return
}
