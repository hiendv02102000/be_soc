package db

import (

	// import source file
	"be_soc/internal/pkg/domain/domain_model/entity"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jinzhu/gorm"
	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *gorm.DB
}

func NewDB() (Database, error) {
	dsn := "bd72d9de6c3c1e:9c1f2305@tcp(us-cdbr-east-05.cleardb.net)/heroku_b6698d216dd2cb8?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	return Database{
		DB: db,
	}, err
}
func (db *Database) MigrateDBWithGorm() {
	db.DB.AutoMigrate(entity.Users{})
	db.DB.AutoMigrate(entity.Comments{})
	db.DB.AutoMigrate(entity.Categories{})
	db.DB.AutoMigrate(entity.Novels{})
	db.DB.AutoMigrate(entity.Chapters{})
	db.DB.AutoMigrate(entity.CommentsChapters{})
	db.DB.AutoMigrate(entity.UsersNovels{})
	db.DB.AutoMigrate(entity.NovelsCategories{})

}

func (db *Database) First(condition interface{}, value interface{}) error {
	err := db.DB.First(value, condition).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	return err
}
func (db *Database) Find(condition interface{}, value interface{}) error {
	err := db.DB.Find(value, condition).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}
func (db *Database) Create(value interface{}) error {
	err := db.DB.Create(value).Error
	return err
}
func (db *Database) Delete(value interface{}) error {
	return db.DB.Delete(value).Error
}
func (db *Database) Update(model interface{}, oldVal interface{}, newVal interface{}) error {
	//fmt.Println(err)
	return db.DB.Model(model).Where(oldVal).Updates(newVal).Error
}
