package filesystem

import "time"

type StorageInterface interface {
	// Disk(disk string) StorageInterface
	Copy(originFile, targetFile string) error
	DeleteDirectory(dirPath string) error
	DeleteFile(filePaths []string) error
	Directories(dir string) []string
	Files(dir string) []string
	MakeDirectory(dir string) error
	Move(originFile, targetFile string) error
	Put(filePath string, content []byte) error
	Size(filePath string) int64
	LastModified(file string) time.Time
	Url(file string) string
}
