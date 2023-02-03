package database

/*
//db.go
import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rahul10-pu/CIGO0122/models"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	passowrd := "postgres"
	args := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable password=" + passowrd
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.Users{})
	DB = db

}
func exam() {
	fmt.Print("ddd")
}
*/

/*
//book.go
package database

import (
	"github.com/jinzhu/gorm"
	"github.com/rahul10-pu/CIGO0122/models"
)
func GetBooks(db *gorm.DB) ([]models.Book, error) {

	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func GetBookByID(db *gorm.DB, id string) (*models.Book, error) {
	book := models.Book{}
	err := db.Select("books.*").Group("books.id").Where("books.id= ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
func DeleteBookByID(db *gorm.DB, id string) error {
	var book models.Book
	err := db.Where("id=?", id).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBookByID(db *gorm.DB, book *models.Book) error {

	err := db.Save(book).Error
	if err != nil {
		return err
	}
	return nil
}
func SaveBook(db *gorm.DB, book *models.Book) error {

	err := db.Save(book).Error
	if err != nil {
		return err
	}
	return nil
}
*/
