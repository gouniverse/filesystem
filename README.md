# File System

<a href="https://gitpod.io/#https://github.com/gouniverse/filesystem" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

![tests](https://github.com/gouniverse/filesystem/workflows/tests/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/filesystem)](https://goreportcard.com/report/github.com/gouniverse/filesystem)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/filesystem)](https://pkg.go.dev/github.com/gouniverse/filesystem)


## Usage

```go
storage, err = filesystem.NewStorage(filesystem.Disk{
  DiskName:             "S3",
  Driver:               filesystem.DRIVER_S3,
  Url:                  config.MediaUrl,
  Region:               config.MediaRegion,
  Key:                  config.MediaKey,
  Secret:               config.MediaSecret,
  Bucket:               config.MediaBucket,
  UsePathStyleEndpoint: true,
})

if err != nil {
  return err.Error()
}
```
