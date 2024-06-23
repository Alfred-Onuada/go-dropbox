package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/Alfred-Onuada/go-dropbox/internals/types"
	// "github.com/gorilla/mux"
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
	// Check that the request method is PATCH
	if r.Method != http.MethodPatch {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the filename from the URL path
	path := r.URL.Path
	parts := strings.Split(path, "/")

	// Check if we have the correct number of parts
	if len(parts) < 3 || parts[2] == "" {
		http.Error(w, "Filename not specified", http.StatusBadRequest)
		return
	}

	filename := parts[2]

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "There was a problem with your request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	///error handling in go too useless
	/* 
	so this is the part where we are supposed to update the file, but we are not updating the file, we are just creating a new file, so what is the point of this function?
	*/

	// Unmarshal the JSON into a File struct
	var file types.File
	err = json.Unmarshal(body, &file)
	//this is supposed to be json.NewDecoder(r.Body).Decode(&file) but I am using Unmarshal because I am reading the body as a byte slice
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "There was a problem with your request", http.StatusBadRequest)
		return
	}

	// Get the index of the file to update
	index, err := getFileIndex(filename) //this is supposed to be getFile(filename) but I am using getFileIndex because I am returning the index of the file
	if err != nil {
		http.Error(w, "The File You are requesting for was not found", http.StatusNotFound)
		return
	}

	// Update the file fields
	if file.Name != "" {
		if FileExists(file.Name) && file.Name != files[index].Name {
			http.Error(w, "Duplicate file names not allowed", http.StatusBadRequest)
			return
		}
		files[index].Name = file.Name
	}
	if file.Size != 0 {
		files[index].Size = file.Size
	}
	if file.Extension != "" {
		files[index].Extension = file.Extension
	}
	if file.Mimetype != "" {
		files[index].Mimetype = file.Mimetype
	}

	// Marshal the updated file to JSON
	jsonResp, err := json.Marshal(files[index])
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "There was a problem processing the JSON", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	// Check if we have the correct number of parts
	if len(parts) < 3 || parts[2] == "" {
		http.Error(w, "Filename not specified", http.StatusBadRequest)
		return
	}

	filename := parts[2]

	// Get the index of the file to delete
	index, err := getFileIndex(filename)
	if err != nil {
		http.Error(w, "The File You are requesting for was not found", http.StatusNotFound)
		return
	}

	// Delete the file from the slice
	files = append(files[:index], files[index+1:]...)

	// Prepare the response
	jsonResp := map[string]interface{}{
		"status":  true,
		"message": "File Deleted Successfully",
	}

	// Marshal the response to JSON
	response, err := json.Marshal(jsonResp)
	if err != nil {
		http.Error(w, "There was a problem processing the JSON response", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func getFile(filename string) (types.File,int,error) {
	//helper function to get a file to avoid repitition

	for i, file := range files {
	  if file.Name == filename {
		return file,i,nil
	  }
	}
	return types.File{},0,errors.New("File not found")
  }

  func FileExists(filename string) bool {
	//helper function to check if a file exists
	for _, file := range files {
	  if file.Name == filename {
		return true
	  }
	}
	return false
	  }
  func getFileIndex(filename string) (int,error) {
	//helper function to get the index of the file
	for i, file := range files {
	  if file.Name == filename {
		return i,nil
	  }
	}
	return -1,fmt.Errorf("File not found")
	  }

