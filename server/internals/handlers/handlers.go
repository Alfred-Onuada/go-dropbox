package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	filename := r.URL.Query().Get("q") //this gets the filename from the query string should update to body

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "There was a problem with your request", http.StatusBadRequest)
		return
	}

	///error handling in go too useless
	/*
	  so this is the part where we are supposed to update the file, but we are not updating the file, we are just creating a new file, so what is the point of this function?
	*/

	/// address pointing doesn't make sense sha , while can't i just send you the file why must i use & to get addres then * to get the value go too useless
	var file types.File
	index, err := getFileIndex(filename) //this is supposed to be getFile(filename) but i am using getFileIndex because i am returning the index of the file

	if err != nil {
		http.Error(w, "The File You are requesting for was not found", http.StatusNotFound)
	}
	err = json.Unmarshal(body, &file) //this is supposed to be json.NewDecoder(r.Body).Decode(&file) but i am using Unmarshal because i am reading the body as a byte slice
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "There was a problem with your request", http.StatusBadRequest)
		return
	}
	fileexists := FileExists(file.Name)
	if fileexists {
		http.Error(w, "Duplicate file names not allowed", http.StatusBadRequest)
		return
	}
	files[index] = file
	jsonresp, err := json.Marshal(files[index])
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "There was a problem processing the JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonresp)
	w.WriteHeader(http.StatusOK)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("q")
	index, err := getFileIndex(filename)
	if err != nil {
		http.Error(w, "The File You are requesting for was not found", http.StatusNotFound)
	}
	files = append(files[:index], files[index+1:]...)
	jsonresp := []byte(`{"status":true,"message":"File Deleted Successfully"}`) // i dey use status if you no like am drink hypo 😂
	w.Header().Set("content-type", "application/json")
	w.Write(jsonresp)
	w.WriteHeader(http.StatusOK)
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
func getFileIndex(filename string) (int, error) {
	//helper function to get the index of the file
	for i, file := range files {
		if file.Name == filename {
			return i, nil
		}
	}
	return 0, errors.New("File not found")
}
