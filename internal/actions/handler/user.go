package handler

import (
"TestProj/internal/model/user"
"encoding/json"
"github.com/gorilla/mux"
"github.com/jinzhu/gorm"
"net/http"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var users []user.User
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var users []user.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&users); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&users).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, users)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	users := getUserOr404(db, name, w, r)
	if users == nil {
		return
	}
	respondJSON(w, http.StatusOK, users)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	users := getUserOr404(db, name, w, r)
	if users == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&users); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&users).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, users)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	users := getUserOr404(db, name, w, r)
	if users == nil {
		return
	}
	if err := db.Delete(&users).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getUserOr404 gets a users instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *user.User {
	users := user.User{}
	if err := db.First(&users, user.User{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &users
}
