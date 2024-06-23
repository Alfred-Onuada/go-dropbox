package types

// capital letter means the field is accesible from outside this package
type File struct {
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	Mimetype  string `json:"mimetype"`
}
