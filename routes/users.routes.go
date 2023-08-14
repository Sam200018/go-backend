package routes

import (
	"encoding/json"
	"main/db"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	param := mux.Vars(r)["id"]

	db.DB.First(&user, param)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Lo siento, usuario no encontrado"))
		return
	} else {
		json.NewEncoder(w).Encode(&user)
	}

}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user, auxUser models.User
	json.NewDecoder(r.Body).Decode(&user)

	db.DB.Where("email = ?", user.Email).First(&auxUser)

	if auxUser.ID != 0 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("User already created"))
		return
	}

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	param := mux.Vars(r)["id"]
	db.DB.First(&user, param)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Lo siento, usuario no encontrado"))
		return
	} else {
		db.DB.Delete(&user)
	}
	w.WriteHeader(http.StatusOK)

}
