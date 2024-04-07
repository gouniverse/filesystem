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

func (s *StaticStorage) DirectoryDelete(dirPath string) error {
	return errors.New("not implemented")
}

func (s *StaticStorage) DirectoriesList(dirPath string) ([]string, error) {
	return []string{}, errors.New("not implemented")
}

func (s *StaticStorage) DirectoryCreate(dirPath string) error {
	return errors.New("not implemented")
}

func (s *StaticStorage) DirectoryCopy(originFile, targetFile string) error {
	return errors.New("not implemented")
}

func (s *StaticStorage) DirectoryLastModified(dirPath string) (time.Time, error) {
	return time.Time{}, errors.New("not implemented")
}

func (s *StaticStorage) DirectorySize(dirPath string) (int64, error) {
	return 0, errors.New("not implemented")
}

func (s *StaticStorage) DirectoryUrl(dirPath string) (string, error) {
	return strings.TrimRight(s.disk.Url, "/") + "/" + strings.TrimLeft(dirPath, "/"), errors.New("not implemented")
}

func (s *StaticStorage) FilesList(dirPath string) ([]string, error) {
	return []string{}, errors.New("not implemented")
}

func (s *StaticStorage) FileCopy(originFile, targetFile string) error {
	return errors.New("not supported")
}

func (s *StaticStorage) FileDelete(filePaths []string) error {
	return errors.New("not implemented")
}

func (s *StaticStorage) FileLastModified(filePath string) (time.Time, error) {
	return time.Time{}, errors.New("not implemented")
}

func (s *StaticStorage) FilePut(filePath string, content []byte) error {
	return errors.New("not implemented")
}

func (s *StaticStorage) FileRead(filePath string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (s *StaticStorage) FileSize(filePath string) (int64, error) {
	return -1, errors.New("not implemented")
}

func (s *StaticStorage) FileUrl(filePath string) (string, error) {
	return strings.TrimRight(s.disk.Url, "/") + "/" + strings.TrimLeft(filePath, "/"), errors.New("not implemented")
}

func (s *StaticStorage) Exists(filePath string) (bool, error) {
	return false, errors.New("not implemented")
}

func (s *StaticStorage) List(dirPath string) ([]string, error) {
	return []string{}, errors.New("not implemented")
}

func (s *StaticStorage) Move(originFile, targetFile string) error {
	return errors.New("not implemented")
}
