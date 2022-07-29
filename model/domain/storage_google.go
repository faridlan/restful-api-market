package domain

import "mime/multipart"

type Storage struct {
	Name string
	File multipart.File
}
