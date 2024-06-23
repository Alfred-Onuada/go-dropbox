package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/Alfred-Onuada/go-dropbox/internals/types"
)

var files []types.File

func getAllFiles(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(files)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "There was a problem processing the JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jsonData)
}

func getASingleFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("q")

	found := false

	for _, file := range files {
		if file.Name == fileName {
			jsonData, err := json.Marshal(file)

			if err != nil {
				fmt.Println(err.Error())
				http.Error(w, "There was a problem processing the JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("content-type", "application/json")
			w.Write(jsonData)

			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Couldn't find the specified file", http.StatusNotFound)
		return
	}
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("q")

	if fileName == "" {
		getAllFiles(w, r)
	} else {
		getASingleFile(w, r)
	}
}

func AddFile(w http.ResponseWriter, r *http.Request) {
	var newFile types.File

	// if lost, check the doc for .Decode, lmao this simply reads the JSON from the r.Body and outputs it to the memory address provided
	err := json.NewDecoder(r.Body).Decode(&newFile)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "There was a problem with your request", http.StatusBadRequest)
		return
	}

	// check that all fields are specified properly
	fileReflection := reflect.TypeOf(newFile)
	for i := 0; i < fileReflection.NumField(); i++ {
		// this checks that the field was not initialized to it's zero value, note this will fail if you have a boolean field, or an int that can be 0 etc.
		// in this case it works fine
		field := fileReflection.Field(i)

		// this is similar to Object.keys and using symbols, the .Interface() just outputs the underlying value
		valueProvided := reflect.ValueOf(newFile).FieldByName(field.Name).Interface()

		if valueProvided == reflect.Zero(field.Type).Interface() {
			http.Error(w, fmt.Sprintf("Please specify a value for %v", field.Name), http.StatusBadRequest)
			return
		}
	}

	// check for duplicate file names
	for _, file := range files {
		if file.Name == newFile.Name {
			http.Error(w, "Duplicate file names not allowed", http.StatusBadRequest)
			return
		}
	}

	files = append(files, newFile)

	fmt.Fprintln(w, "New file added successfully")
}

func UpdateFile(w http.ResponseWriter, r *http.Request) {

}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

}
