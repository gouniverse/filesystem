package filesystem

import "time"

type StorageInterface interface {
	Copy(originFile, targetFile string) error
	DeleteDirectory(dirPath string) error
	DeleteFile(filePaths []string) error
	Directories(dir string) ([]string, error)
	Files(dir string) ([]string, error)
	MakeDirectory(dir string) error
	Move(originFile, targetFile string) error
	Put(filePath string, content []byte) error
	Size(filePath string) (int64, error)
	LastModified(file string) (time.Time, error)
	Url(file string) (string, error)
}
