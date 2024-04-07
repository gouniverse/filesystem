package filesystem

import "time"

type DirectoryStorageInterface interface {
	// DirectoriesList returns a slice of all the directories within the given directory.
	DirectoriesList(dir string) (dirs []string, err error)

	// FilesList returns a slice of all the files within the given directory.
	FilesList(dir string) (files []string, err error)

	// List returns a slice of all the directories and files within the given directory.
	List(dir string) (files []string, err error)

	// DirectoryCopy copies a file to a new location.
	DirectoryCopy(originDirPath, targetDirPath string) (err error)

	// DirectoryCreate creates a new directory.
	DirectoryCreate(dir string) (err error)

	// DirectoryDelete deletes a directory and all its subdirectories.
	DirectoryDelete(dirPath string) (err error)

	// DirectorySize retrieves the size of the file in bytes.
	DirectorySize(dirPath string) (size int64, err error)

	// DirectoryUrl returns the public url of the file.
	DirectoryUrl(dirPath string) (url string, err error)
}

type FileStorageInterface interface {
	// FileCopy copies a file to a new location.
	FileCopy(originFilePath, targetFilePath string) (err error)

	// FileDelete deletes one or more files.
	FileDelete(filePaths []string) (err error)

	// FileLastModified returns the file's last modification time.
	FileLastModified(file string) (modTime time.Time, err error)

	// FilePut stores a file in the storage.
	FilePut(filePath string, content []byte) (err error)

	// FileRead reads the entire file into memory and returns its content.
	FileRead(filePath string) (content []byte, err error)

	// FileSize retrieves the size of the file in bytes.
	FileSize(filePath string) (size int64, err error)

	// FileUrl returns the public url of the file.
	FileUrl(filePath string) (url string, err error)
}

// StorageInterface represents the interface of a storage
type StorageInterface interface {
	DirectoryStorageInterface

	FileStorageInterface

	// Exists checks if a file or directory exists.
	Exists(filePath string) (exists bool, err error)

	// Move moves a file to a new location.
	Move(originFile, targetFile string) (err error)
}
