package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	gorm.Model
	Name     string   `json:"name"`
	Path     string   `json:"path"` // AWS S3 path
	ItemType ItemType `json:"itemType"`
	ParentID *uint    `json:"-"` // Use ParentID to reference the parent Item
	Parent   *Item    `json:"-"` // Self-referencing foreign key
	Items    []*Item  `json:"items,omitempty" gorm:"foreignKey:ParentID"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	ItemID   *uint  `json:"-"` // Use ItemID to reference the root Item
	Item     *Item  `json:"item,omitempty" gorm:"foreignKey:ItemID"`
}


func HandleMigration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
