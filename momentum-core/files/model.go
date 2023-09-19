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
	Id   string
	Name string
	Body string
}

type Dir struct {
	Id      string
	Name    string
	SubDirs []*Dir
	Files   []*File
}

type Overwrite struct {
	OriginFileId      string
	OriginFileLine    int
	OverwriteFileId   string
	OverwriteFileLine int
}

func NewFile(id string, name string, encodedBody string) *File {

	f := new(File)

	f.Id = id
	f.Name = name
	f.Body = encodedBody

	return f
}
