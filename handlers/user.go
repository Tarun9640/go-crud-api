package handlers

import (
	"crud-go/models"
	"crud-go/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	// Dummy data 
	users := storage.Users

	// Tell client we are sending JSON
	w.Header().Set("Content-Type", "application/json")

	// Always set status code
	w.WriteHeader(http.StatusOK)

	// Convert Go → JSON and send
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w,"Method Not allowed",http.StatusMethodNotAllowed)
	}

	var user models.User

	// Decode JSON request → struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,"Unable to reab json",http.StatusBadRequest)
		return
	}

	//fake id 
	user.ID = len(storage.Users) + 1

	storage.Users = append(storage.Users, user)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w,"method not correct",http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts) - 1]

	id,err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w,"invalid user id",http.StatusBadRequest)
		return
	}

	for _,user := range storage.Users {
		if id == user.ID {
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(user)

			return
		}
	}
	//If loop completes → user not found
	http.Error(w, "user not found", http.StatusNotFound)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}

	var updatedUser models.User

	//// Extract ID
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts) - 1]

	id,err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w,"invalid user id",http.StatusBadRequest)
		return
	}

	//Decode() returns ONLY error
	err = json.NewDecoder(r.Body).Decode(&updatedUser)

	if err != nil {
		http.Error(w,"Unable to read json",http.StatusBadRequest)
		return
	}

	//if you find id update with input json
	//If you want to MODIFY slice → use INDEX.because range return copy..if u use value u change only copy not orginal
	for i := range storage.Users {
		if storage.Users[i].ID == id {
			storage.Users[i].Name = updatedUser.Name
			storage.Users[i].Email = updatedUser.Email

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(storage.Users[i])
			return
		}
	}

	http.Error(w,"user not found",http.StatusNotFound)

}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}

	//var deleteUser models.User

	// Extract ID
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts) - 1]

	id,err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w,"invalid user id",http.StatusBadRequest)
		return
	}

	// Find and delete
	for i := range storage.Users {
		if storage.Users[i].ID == id {
			//Delete item from slice
			//Slice deletion means: Reconnect memory : A → B → C to A → C
			storage.Users = append(storage.Users[:i],storage.Users[i+1:]...)

			w.WriteHeader(http.StatusNoContent)

			return
		}
	}

	http.Error(w,"user not found to delete",http.StatusNotFound)

}
