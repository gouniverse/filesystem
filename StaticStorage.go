package filesystem

import (
	"errors"
	"strings"
	"time"
)

// StaticStorage implements StorageInterface
// it represents a static (read only) storage, i.e. CDN
// the only supported method is Url(filepath)
type StaticStorage struct {
	disk Disk
}

var _ StorageInterface = (*StaticStorage)(nil) // verify it extends the task interface

func (s *StaticStorage) Copy(originFile, targetFile string) error {
	return errors.New("not supported")
}

func (s *StaticStorage) DeleteFile(filePaths []string) error {
	return errors.New("not supported")
}

func (s *StaticStorage) DeleteDirectory(dirPath string) error {
	return errors.New("not supported")
}

func (s *StaticStorage) Directories(dirPath string) []string {
	return []string{}
}

// implement StorageInterface

func (s *StaticStorage) Files(dirPath string) []string {
	return []string{}
}

func (s *StaticStorage) MakeDirectory(dirPath string) error {
	return errors.New("not supported")
}

func (s *StaticStorage) LastModified(filePath string) time.Time {
	return time.Time{}
}

func (s *StaticStorage) Move(originFile, targetFile string) error {
	return errors.New("not supported")
}

func (s *StaticStorage) Size(filePath string) int64 {
	return 0
}

func (s *StaticStorage) Url(filePath string) string {
	return strings.TrimRight(s.disk.Url, "/") + "/" + strings.TrimLeft(filePath, "/")
}

func (s *StaticStorage) Put(filePath string, content []byte) error {
	return errors.New("not supported")
}
