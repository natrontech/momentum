package files

type IFile interface {
	Id() string
	Name() string
	Body() string // Base64 encoded
}

type IDir interface {
	Id() string
	Name() string
	SubDirs() []IDir
	Files() []IFile
}

type IOverwrite interface {
	OriginFileId() string
	OriginFileLine() int
	OverwriteFileId() string
	OverwriteFileLine() int
}

type File struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Body string `json:"body"`
}

type CreateFileRequest struct {
	ParentId string `json:"parentId"`
	Name     string `json:"name"`
	Body     string `json:"body"`
}

type Dir struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	SubDirs []*Dir  `json:"subDirs"`
	Files   []*File `json:"files"`
}

type Overwrite struct {
	OriginFileId      string `json:"originFileId"`
	OriginFileLine    int    `json:"originFileLine"`
	OverwriteFileId   string `json:"overwriteFileId"`
	OverwriteFileLine int    `json:"overwriteFileLine"`
}

func NewFile(id string, name string, encodedBody string) *File {

	f := new(File)

	f.Id = id
	f.Name = name
	f.Body = encodedBody

	return f
}
