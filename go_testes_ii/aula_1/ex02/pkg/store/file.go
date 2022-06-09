package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	ReadWasCalled() bool
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName, false}
	}
	return nil
}

type FileStore struct {
	FileName string
	readWasCalled bool
}

func (fs *FileStore) Read(data interface{}) error {
	fs.readWasCalled = true
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStore) ReadWasCalled() bool {
	return fs.readWasCalled
}
