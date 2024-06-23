package types

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