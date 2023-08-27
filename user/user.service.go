package user

import (
	"encoding/json"
	"goproject/db"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User

	res := db.DB.Find(&users)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(r)
	res := db.DB.First(&user, params["id"])
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userDto createUserDto
	json.NewDecoder(r.Body).Decode(&userDto)
	user := User{
		FisrtName: userDto.FisrtName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
	}
	res := db.DB.Create(&user)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user User
	res := db.DB.First(&user, params["id"])
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	var updateDto updateUserDto
	json.NewDecoder(r.Body).Decode(&updateDto)

	if updateDto.FisrtName != "" {
		user.FisrtName = updateDto.FisrtName
	}
	if updateDto.LastName != "" {
		user.LastName = updateDto.LastName
	}
	if updateDto.Email != "" {
		user.Email = updateDto.Email
	}

	json.NewDecoder(r.Body).Decode(&user)
	resSave := db.DB.Save(&user)
	if resSave.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resSave.Error.Error())
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(r)
	res := db.DB.Delete(&user, params["id"])
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}
	json.NewEncoder(w).Encode("The user is succesfully deleted")
}
