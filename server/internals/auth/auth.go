package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Alfred-Onuada/go-dropbox/internals/types"
)

// LoginHandler is a handler for the login route
var users []types.User
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// get the request body
	decoder := json.NewDecoder(r.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "There was a problem processing the JSON", http.StatusInternalServerError)
		return
	}

	// get the username and password
	username := data["username"]
	password := data["password"]
      
	// check if the username and password are correct
	 for _, user := range users {
		if user.Username == username && user.Password == password {		//when db add password hashing and verification
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(interface{}(user.Item).([]byte)))
		return
		}
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	
}

// RegisterHandler is a handler for the register route
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// get the request body
	decoder := json.NewDecoder(r.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "There was a problem processing the JSON", http.StatusInternalServerError)
		return
	}

	// get the username and password
	username := data["username"]
	password := data["password"]

	// check if the username and password are correct
	 curDir,err := os.Getwd()
	 if err != nil {
		http.Error(w, "There was a problem getting the current directory", http.StatusInternalServerError)
		return
	}
	userDir := curDir + "/" + username
    err = os.Mkdir(userDir, 0755) 
    if err != nil {
        http.Error(w, "Failed to create user directory", http.StatusInternalServerError)
        return
    }
	 user := types.User{ 
		Username: username,
		Password: password,
		Item: nil,// change to either db or s3path
	 } 
	 item := types.Item{
		Name: username,
		Path: userDir,
		ItemType: types.Folder,
		Items: nil,
	 }
	 user.Item = &item
	  users = append(users, user)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Registration successful"}`))
	
}

