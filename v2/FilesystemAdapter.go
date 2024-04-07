package filesystem

import (
	"time"
)

// filesystemAdapterCommonInterface represents the interface for directory and file
type filesystemAdapterCommonInterface interface {
	Attributes(path string) (*Attributes, error)
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
}

type filesystemAdapterFileInterface interface {
	FileDelete(filePath string) error
	// MAYBE: FilesDelete(filePaths []string) error
	FileExists(path string) bool // new
	FileRead(path string) (string, error)
	FileReadStream(path string) ([]byte, error)
	FileSize(filePath string) (int64, error)
	FileWrite(path string, content []byte, config map[string]string) error
	FileWriteStream(path string, content []byte, config map[string]string) error
}

type filesystemAdapterDirectoryInterface interface {
	DirectoryCreate(dirPath string, config map[string]string) error
	DirectoryDelete(dirPath string) error
	// MAYBE: DirectoriesDelete(dirPaths []string) error
	DirectoriesList(dir string) ([]string, error)
	DirectoryExists(path string) bool // new
	FilesList(dir string) ([]string, error)
}

type FilesystemAdapter interface {
	filesystemAdapterCommonInterface
	filesystemAdapterDirectoryInterface
	filesystemAdapterFileInterface
}
