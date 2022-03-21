package handler

import (
	"TestProj/internal/model/book"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var books []book.Book
	db.Find(&books)
	respondJSON(w, http.StatusOK, books)
}

func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var books []book.Book

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&books); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&books).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, books)
}

func GetBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	books := book.Book{}
	if err := db.First(&books).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}
	respondJSON(w, http.StatusOK, books)
}

func GetBookByPage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	books := book.Book{}
	if err := db.Find(&books, book.Book{ID : id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}
	respondJSON(w, http.StatusOK, books)
}

func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	books := book.Book{}
	if err := db.Find(&books, book.Book{ID : id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&books); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&books).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, books)
}

func DeleteBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	books := book.Book{}
	if err := db.Find(&books, book.Book{ID : id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}
	if err := db.Delete(&books).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	books := book.Book{}
	if err := db.Find(&books, book.Book{ID : id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}
	books.Disable()
	if err := db.Save(&books).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, books)
}

func EnableBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	books := book.Book{}
	if err := db.Find(&books, book.Book{ID : id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}
	books.Enable()
	if err := db.Save(&books).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, books)
}

// getBookOr404 gets a books instance if exists, or respond the 404 error otherwise
func getBookOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *book.Book {
	books := book.Book{}
	if err := db.First(&books, book.Book{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &books
}

