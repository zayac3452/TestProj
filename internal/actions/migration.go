package actions

import (
	"TestProj/internal/model/book"
	"TestProj/internal/model/user"
	"github.com/jinzhu/gorm"
)

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&book.Book{}, &user.User{})
	return db
}

