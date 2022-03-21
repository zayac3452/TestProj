package user

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string	`gorm:"unique" json:"name"`
	Surname string `json:"surname"`
}
