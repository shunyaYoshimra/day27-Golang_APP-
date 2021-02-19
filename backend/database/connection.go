package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

var conn *gorm.DB
var err error

func Connection() {
	provider := os.Getenv("DB_PROVIDER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	authDB := user + ":" + pass + "@/" + name + "?parseTime=true"
	conn, err = gorm.Open(provider, authDB)
	if err != nil {
		panic(err)
	}
}

func AppConnection() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}
	Connection()
}

func TestConnection() {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}
	Connection()
}

func GetDB() *gorm.DB {
	return conn
}

func DropAllTable() {
	conn.DropTable(
		entity.User{},
		entity.Profile{},
		entity.Contact{},
		entity.Abroad{},
		entity.Occupation{},
		entity.Question{},
		entity.Answer{},
		entity.Post{},
		entity.Image{},
		entity.Favorite{},
		entity.Article{},
		entity.ArticleLine{},
		entity.BoldNumber{},
		entity.LinkNumber{},
		entity.Check{},
	)
}
