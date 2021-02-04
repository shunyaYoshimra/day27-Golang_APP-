package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository() UserRepository {
	return UserRepository{Conn: database.GetDB().Table("users")}
}

func (ur *UserRepository) RetrieveUsers() []entity.User {
	var users []entity.User
	ur.Conn.Find(&users)
	return users
}

func (ur *UserRepository) FindByID(id int) (user entity.User, err error) {
	err = ur.Conn.First(&user, id).Error
	return
}

func (ur *UserRepository) FindByEmail(email string) (user entity.User, err error) {
	err = ur.Conn.Where("email = ?", email).First(&user).Error
	fmt.Println("-------")
	fmt.Println(err)
	fmt.Println("-------")
	return
}

func (ur *UserRepository) Create(user *entity.User) (err error) {
	err = ur.Conn.Create(user).Error
	return
}
