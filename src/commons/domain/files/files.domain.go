package filesDomain

import "io"

type File struct {
	Id          string
	Name        string
	Data        io.Reader
	ContentType string
	size        int
}

type FileRes struct {
	Id   string
	Name string
	Url  string
}
