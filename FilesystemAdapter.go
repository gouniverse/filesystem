package filesystem

import (
	"time"
)

type FilesystemAdapter interface {
	// Common Methods
	Copy(originPath, targetPath string) error
	Delete(path string) error
	Exists(path string) bool
	Extension(filePath string) (string, error)
	LastModified(file string) (time.Time, error)
	List(filePath string) ([]string, error)
	MimeType(filePath string) (string, error)
	Move(originFile, targetFile string) error
	Visibility(filePath string) (string, error)
	Url(file string) (string, error)
	SetVisibility(filePath string, visibility string)

	Attributes(filePath string) (*Attributes, error)

	// Directory Methods
	DirectoryCreate(dirPath string, config map[string]string) error
	DirectoryDelete(dirPaths string) error
	DirectoriesDelete(dirPaths []string) error
	DirectoryList(dir string) ([]string, error)
	DirectoryExists(path string) bool // new

	// File Methods
	FileDelete(filePath string) error
	FilesDelete(filePaths []string) error
	FileExists(path string) bool // new
	FileList(dir string) ([]string, error)
	FileRead(path string) (string, error)
	FileReadStream(path string) ([]byte, error)
	FileSize(filePath string) (int64, error)
	FileWrite(path string, content []byte, config map[string]string) error
	FileWriteStream(path string, content []byte, config map[string]string) error
}
