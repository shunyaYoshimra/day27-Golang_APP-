package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type AnswerRepository struct {
	Conn *gorm.DB
}

func NewAnswerRepository() AnswerRepository {
	return AnswerRepository{Conn: database.GetDB().Table("answers")}
}

func (ar *AnswerRepository) RetrieveAnswers(id int) (answers []entity.Answer) {
	ar.Conn.Where("question_id = ?", id).Order("created_at desc").Find(&answers)
	return
}

func (ar *AnswerRepository) FindByID(id int) (answer entity.Answer, err error) {
	err = ar.Conn.First(&answer, id).Error
	return
}

func (ar *AnswerRepository) Create(answer *entity.Answer) (err error) {
	err = ar.Conn.Create(answer).Error
	return
}

func (ar *AnswerRepository) Update(answer entity.Answer, content string) (err error) {
	err = ar.Conn.Model(&answer).Update(entity.Answer{Content: content}).Error
	return
}
