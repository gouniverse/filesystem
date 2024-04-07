package filesystem

import (
	"time"
)

// filesystemAdapterCommonInterface represents the interface for directory and file
type filesystemAdapterCommonInterface interface {
	// Attributes returns the attributes of the file or directory
	Attributes(path string) (*Attributes, error)

	// Copy copies a file to a new location
	Copy(originPath, targetPath string) error

	// Delete deletes a file or directory
	Delete(path string) error

	// Exists checks if a file or directory exists
	Exists(path string) bool

	// LastModified returns the file or directory last modification time
	LastModified(file string) (time.Time, error)

	// List returns a slice of all the directories and files within the given directory
	List(path string) ([]string, error)

	// Move moves a file or directory to a new location
	Move(originFile, targetFile string) error

	// Size returns the size of the file or directory
	Visibility(filePath string) (string, error)

	// Url returns the public url of the file or directory
	Url(file string) (string, error)

	// SetVisibility sets the visibility of the file or directory
	SetVisibility(filePath string, visibility string)
}

type filesystemAdapterFileInterface interface {
	FileDelete(filePath string) error
	// MAYBE: FilesDelete(filePaths []string) error
	FileExists(path string) bool // new
	FileExtension(filePath string) (string, error)

	// FileMimeType returns the mime type of the file
	FileMimeType(filePath string) (string, error)

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
