package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Alfred-Onuada/go-dropbox/internals/helpers"
	"github.com/Alfred-Onuada/go-dropbox/internals/types"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN environment variable not set")
	}

	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
}

func TestAuth(w http.ResponseWriter, r *http.Request){
	username := r.Context().Value("username").(string)

	w.Write([]byte(username))
	return
}
// LoginHandler is a handler for the login route
var users []types.User
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// get the request body

	decoder := json.NewDecoder(r.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		helpers.JSONResponse(w,false,"An Error Occured",http.StatusInternalServerError,nil)
		return
	}
   
	username := data["username"]
	password := data["password"]
      
	var user types.User
	if err := DB.Where("username = ?", username).First(&user); err == nil {
		helpers.JSONResponse(w,false,"Invalid Credentials",http.StatusUnauthorized,nil)
		return

	}
	validpassword := checkPasswordHash(password,user.Password)
	
	if !validpassword  {
		helpers.JSONResponse(w,false,"Invalid Credentials",http.StatusUnauthorized,nil)
		return

	} 
	 token,err := GenerateJWT(user.Username)
	 if err !=nil {
		helpers.JSONResponse(w,false ,"Something Went Wrong , ask your ex",http.StatusInternalServerError,nil)
		return
	 }
		jsonresp := map[string]interface{}{
			"item": user.Item,
			"tokentype" : "Bearer",
			"token": token,
		}
		helpers.JSONResponse(w,true,"Login Successfull",http.StatusOK,jsonresp)
		return
		
	}
	


// RegisterHandler is a handler for the register route
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// get the request body
	decoder := json.NewDecoder(r.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		helpers.JSONResponse(w,false,"There was a problem processing the JSON",http.StatusInternalServerError,nil)
		return

	}
	
	// get the username and password
	username := data["username"]
	password := data["password"]
	 var existinguser types.User
	if err := DB.Where("username = ?" , username).First(&existinguser).Error; err==nil {
		helpers.JSONResponse(w,false,"Username Is Taken",http.StatusBadRequest,nil)
		return

	}
	 
	 curDir,err := os.Getwd() // change later to s3 
	 if err != nil {
		helpers.JSONResponse(w,false,"Unable to get the directory",http.StatusInternalServerError,nil)
		return
	}
	 userDir := curDir + "/users/" + username
	 hashpassword,err := hashPassword(password)
	 if err != nil {
		helpers.JSONResponse(w,false,"An Unknown error occured",http.StatusInternalServerError,nil)
		return

	 }
	 user := types.User{ 
		Username: username,
		Password: hashpassword,
		Item: nil,// change to either db or s3path
	 } 
	 item := types.Item{
		Name: username,
		Path: userDir,
		ItemType: types.Folder,
		Items: nil,
	 }
	 user.Item = &item
	 if err := DB.Create(&user).Error; err != nil {
		helpers.JSONResponse(w,false,"There was a problem processing your request",http.StatusInternalServerError,nil)
		return

	}
	helpers.JSONResponse(w,true,"Login Successfull",http.StatusOK,nil)
	return

	
}


func IsDirExists(path string) (bool, error) {
    info, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            // The directory does not exist
            return false, nil
        }
        // Some other error occurred, like a permission issue
        return false, err
    }
    // Return true if the path is a directory
    return info.IsDir(), nil
}


func hashPassword(password string) (string, error) {
	// helper function to hash passwords
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	// helper function to check passwords hash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
