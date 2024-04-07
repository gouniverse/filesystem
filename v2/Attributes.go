package filesystem

import "time"

func NewAttributes() *Attributes {
	return &Attributes{}
}

type Attributes struct {
	isDirectory  bool
	isFile       bool
	path         string
	lastModified time.Time
	mimeType     string
	extension    string
	size         int64
	visibility   string
}

func (a *Attributes) IsDirectory() bool {
	return a.isDirectory
}

func (a *Attributes) SetIsDirectory(isDirectory bool) *Attributes {
	a.isDirectory = isDirectory
	return a
}

func (a *Attributes) IsFile() bool {
	return a.isFile
}

func (a *Attributes) SetIsFile(isFile bool) *Attributes {
	a.isFile = isFile
	return a
}

func (a *Attributes) Extension() string {
	return a.extension
}

func (a *Attributes) SetExtension(extension string) *Attributes {
	a.extension = extension
	return a
}

func (a *Attributes) LastModifiedTime() time.Time {
	return a.lastModified
}

func (a *Attributes) SetLastModifiedTime(lastModified time.Time) *Attributes {
	a.lastModified = lastModified
	return a
}

func (a *Attributes) MimeType() string {
	return a.mimeType
}

func (a *Attributes) SetMimeType(mimeType string) *Attributes {
	a.mimeType = mimeType
	return a
}

func (a *Attributes) Path() string {
	return a.path
}

func (a *Attributes) SetPath(path string) *Attributes {
	a.path = path
	return a
}

func (a *Attributes) Size() int64 {
	return a.size
}

func (a *Attributes) SetSize(size int64) *Attributes {
	a.size = size
	return a
}

func (a *Attributes) Visibility() string {
	return a.visibility
}

func (a *Attributes) SetVisibility(visibility string) *Attributes {
	a.visibility = visibility
	return a
}
