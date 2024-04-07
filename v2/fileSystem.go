package filesystem

import "time"

func FileSystem(filesystemAdapter FilesystemAdapter) *fileSystem {
	return &fileSystem{
		adapter: filesystemAdapter,
	}
}

type fileSystem struct {
	adapter FilesystemAdapter
}

func (f *fileSystem) Copy(originPath string, targetPath string) error {
	return f.adapter.Copy(originPath, targetPath)
}

func (f *fileSystem) Delete(path string) error {
	return f.adapter.Delete(path)
}

func (f *fileSystem) DirectoryCreate(path string, config map[string]string) error {
	return f.adapter.DirectoryCreate(path, config)
}

func (f *fileSystem) DirectoryDelete(path string) error {
	return f.adapter.DirectoryDelete(path)
}

func (f *fileSystem) DirectoryExists(path string) bool {
	return f.adapter.DirectoryExists(path)
}

func (f *fileSystem) DirectoriesList(path string) ([]string, error) {
	return f.adapter.DirectoriesList(path)
}

// func (f *fileSystem) DirectoriesDelete(path []string) error {
// 	return f.adapter.DirectoriesDelete(path)
// }

func (f *fileSystem) Exists(path string) bool {
	return f.adapter.Exists(path)
}

func (f *fileSystem) FileDelete(path string) error {
	return f.adapter.FileDelete(path)
}

// func (f *fileSystem) FileExtension(path string) string {
// 	return f.adapter.FileExtension(path)
// }

func (f *fileSystem) FileExists(path string) bool {
	return f.adapter.FileExists(path)
}

func (f *fileSystem) FilesList(path string) ([]string, error) {
	return f.adapter.FilesList(path)
}

func (f *fileSystem) FileRead(path string) (string, error) {
	return f.adapter.FileRead(path)
}

func (f *fileSystem) FileReadStream(path string) ([]byte, error) {
	return f.adapter.FileReadStream(path)
}

func (f *fileSystem) FileSize(path string) (int64, error) {
	return f.adapter.FileSize(path)
}

func (f *fileSystem) FileWrite(path string, contents []byte, config map[string]string) error {
	return f.adapter.FileWrite(path, contents, config)
}

func (f *fileSystem) FileWriteStream(path string, contents []byte, config map[string]string) error {
	return f.adapter.FileWriteStream(path, contents, config)
}

func (f *fileSystem) Has(path string) bool {
	return f.adapter.Exists(path)
}

func (f *fileSystem) LastModified(path string) (time.Time, error) {
	return f.adapter.LastModified(path)
}

func (f *fileSystem) List(path string) ([]string, error) {
	return f.adapter.List(path)
}

func (f *fileSystem) FileMimeType(path string) (string, error) {
	return f.adapter.FileMimeType(path)
}

func (f *fileSystem) Move(originPath string, targetPath string) error {
	return f.adapter.Move(originPath, targetPath)
}

func (f *fileSystem) Missing(path string) bool {
	return !f.adapter.Exists(path)
}

func (f *fileSystem) Url(path string) (string, error) {
	return f.adapter.Url(path)
}

func (f *fileSystem) Visibility(path string) (string, error) {
	return f.adapter.Visibility(path)
}

func (f *fileSystem) SetVisibility(path string, visibility string) {
	f.adapter.SetVisibility(path, visibility)
}

// func (f *fileSystem) PublicUrl(path string, config map[string]string) (string, error) {
// 	return f.adapter.PublicUrl(path, config)
// }

// func (f *fileSystem) TemporaryUrl(path string, expiresAt time.Time, config map[string]string) (string, error) {
// 	return f.adapter.TemporaryUrl(path, expiresAt, config)
// }

//     public function checksum(string $path, array $config = []): string
//     {
//         $config = $this->config->extend($config);

//         if ( ! $this->adapter instanceof ChecksumProvider) {
//             return $this->calculateChecksumFromStream($path, $config);
//         }

//         try {
//             return $this->adapter->checksum($path, $config);
//         } catch (ChecksumAlgoIsNotSupported) {
//             return $this->calculateChecksumFromStream($path, $config);
//         }
//     }
