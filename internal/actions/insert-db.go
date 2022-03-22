package actions

import (
	"TestProj/internal/model/book"
	"github.com/jinzhu/gorm"
)

func InsertBooksDB(db *gorm.DB)  {
	var books = []book.Book{{Name: "TitleBook_1",Annotation : "SomeAnnotation_1",Authors: "Somebody_1", PublishingYear: 1949, PublishingName: "NameOfPublish", Copies: 1, Address: "A1", Status: false},
		{Name: "TitleBook_2",Annotation : "SomeAnnotation_2",Authors: "Somebody_2", PublishingYear: 1949, PublishingName: "NameOfPublish", Copies: 1, Address: "A1", Status: false},
		{Name: "TitleBook_3",Annotation : "SomeAnnotation_3",Authors: "Somebody_3", PublishingYear: 1949, PublishingName: "NameOfPublish", Copies: 1, Address: "A1", Status: false},
		{Name: "TitleBook_4",Annotation : "SomeAnnotation_4",Authors: "Somebody_4", PublishingYear: 1949, PublishingName: "NameOfPublish", Copies: 1, Address: "A1", Status: false},
		{Name: "TitleBook_5",Annotation : "SomeAnnotation_5",Authors: "Somebody_5", PublishingYear: 1949, PublishingName: "NameOfPublish", Copies: 1, Address: "A1", Status: false},
		{Name: "TitleBook_6",Annotation : "SomeAnnotation_6",Authors: "Somebody_6", PublishingYear: 1949, PublishingName: "NameOfPublish", Copies: 1, Address: "A1", Status: false},
		{Name: "TitleBook_7",Annotation : "SomeAnnotation_7",Authors: "Somebody_7", PublishingYear: 1949, PublishingName: "NameOfPublish", Copies: 1, Address: "A1", Status: false},
		{Name: "TitleBook_8",Annotation : "SomeAnnotation_8",Authors: "Somebody_8", PublishingYear: 1949, PublishingName: "NameOfPublish", Copies: 1, Address: "A1", Status: false}}

	for i, _ := range books {
		db.Create(&books[i])
	}
}
