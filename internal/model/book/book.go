package book

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
	Annotation string `json:"annotation"`
	Authors string `json:"authors"`
	PublishingYear int `json:"publishing_year"`
	PublishingName string `json:"publishing_name"`
	Copies int `json:"copies"`
	Address string `json:"address"`
	Status bool `json:"status"`
}

func (b *Book) Disable() {
	b.Status = false
}

func (b *Book) Enable() {
	b.Status = true
}
