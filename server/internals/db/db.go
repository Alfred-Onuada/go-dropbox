package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type File struct {
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	Mimetype  string `json:"mimetype"`
}
type ItemType int
const(
	Folder ItemType = iota
	Files
)
type Item struct {
	Name string `json:"name"`
	Path string `json:"path"` //aws s3 path
	ItemType  ItemType `json:"itemType"`
	Items []*Item `json:"items,omitempty"`
}

type User struct { 

	Username string `json:"username"`
	Password string `json:"password"`
	Item *Item `json:"item,omitempty"`
}

func HandleMigration() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN environment variable not set")
	}

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
		return
	}
	 db.AutoMigrate(&User{})
	 db.AutoMigrate(&Item{})
	 db.AutoMigrate(&File{})
	//  db.AutoMigrate(&User{}) add more types when finished

}
