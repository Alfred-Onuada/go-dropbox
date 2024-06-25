package types

import "gorm.io/gorm"

// capital letter means the field is accesible from outside this package
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

type BaseResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data  interface{} `json:"data,omitempty"`
}