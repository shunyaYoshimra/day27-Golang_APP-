package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type QuestionRepository struct {
	Conn *gorm.DB
}

func NewQuestionRepository() QuestionRepository {
	return QuestionRepository{Conn: database.GetDB().Table("questions")}
}

func (qr *QuestionRepository) RetrieveQuestions() (questions []entity.Question) {
	qr.Conn.Order("created_at desc").Find(&questions)
	return
}

func (qr *QuestionRepository) FindByID(id int) (question entity.Question, err error) {
	err = qr.Conn.First(&question, id).Error
	return
}

func (qr *QuestionRepository) FindByUser(id int) (questions []entity.Question) {
	qr.Conn.Where("user_id = ?", id).Order("created_at desc").Find(&questions)
	return
}

func (qr *QuestionRepository) FindByAbout(about string) (questions []entity.Question) {
	qr.Conn.Where("about = ?", about).Order("created_at desc").Find(&questions)
	return
}

func (qr *QuestionRepository) Create(question *entity.Question) (err error) {
	err = qr.Conn.Create(question).Error
	return
}

func (qr *QuestionRepository) Update(question entity.Question, title, content string) (err error) {
	err = qr.Conn.Model(&question).Update(entity.Question{Title: title, Content: content}).Error
	return
}

func (qr *QuestionRepository) ChangeStatus(question entity.Question) (err error) {
	err = qr.Conn.Model(&question).Update(entity.Question{Status: true}).Error
	return
}

func (qr *QuestionRepository) Delete(question entity.Question) (err error) {
	err = qr.Conn.Delete(&question).Error
	return
}
