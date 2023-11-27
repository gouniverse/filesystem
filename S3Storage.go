package filesystem

import (
	"context"
	"io/ioutil"
	"net/http"
	"path"

	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gabriel-vasile/mimetype"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/support/file"
	"github.com/gouniverse/utils"
	"github.com/mingrammer/cfmt"
)

// S3Storage implements the StorageInterface for an S3 compliant file storage,
// i.e. AWS S3, DigitalOcean Spaces, Minio, etc
type S3Storage struct {
	disk Disk
}

var _ StorageInterface = (*S3Storage)(nil) // verify it extends the task interface

func (s *S3Storage) client() (*s3.Client, error) {
	customResolver := s3.EndpointResolverFunc(func(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error) {
		return aws.Endpoint{
			// PartitionID:   "aws",
			URL: "https://" + s.disk.Url,
			// SigningRegion: "us-west-2",
			HostnameImmutable: true,
		}, nil
	})

	client := s3.New(s3.Options{
		Region:           s.disk.Region,
		Credentials:      aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(s.disk.Key, s.disk.Secret, "")),
		EndpointResolver: customResolver,
		UsePathStyle:     s.disk.UsePathStyleEndpoint,
	})
	return client, nil
}

func (s *S3Storage) Copy(originFile, targetFile string) error {
	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	_, err = s3Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(s.disk.Bucket),
		CopySource: aws.String(s.disk.Bucket + "/" + originFile),
		Key:        aws.String(targetFile),
	})

	return err
}

func (s *S3Storage) DeleteFile(filePaths []string) error {
	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}

	var objectIdentifiers []types.ObjectIdentifier
	for _, file := range filePaths {
		objectIdentifiers = append(objectIdentifiers, types.ObjectIdentifier{
			Key: aws.String(file),
		})
	}

	quiet := true
	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(s.disk.Bucket),
		Delete: &types.Delete{
			Objects: objectIdentifiers,
			Quiet:   &quiet,
		},
	}
	ctx := context.TODO()
	_, err = s3Client.DeleteObjects(ctx, input)

	return err
}

func (s *S3Storage) DeleteDirectory(directory string) error {
	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}

	if !strings.HasSuffix(directory, "/") {
		directory += "/"
	}

	listObjectsV2Response, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(s.disk.Bucket),
		Prefix: aws.String(directory),
	})
	if err != nil {
		return err
	}
	if len(listObjectsV2Response.Contents) == 0 {
		return nil
	}

	for {
		for _, item := range listObjectsV2Response.Contents {
			_, err = s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
				Bucket: aws.String(s.disk.Bucket),
				Key:    item.Key,
			})
			if err != nil {
				return err
			}
		}

		if *listObjectsV2Response.IsTruncated {
			listObjectsV2Response, err = s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
				Bucket:            aws.String(s.disk.Bucket),
				ContinuationToken: listObjectsV2Response.ContinuationToken,
			})
			if err != nil {
				return err
			}
		} else {
			break
		}
	}

	return nil
}

func (s *S3Storage) Directories(dir string) []string {
	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}
	dirs := []string{}
	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(s.disk.Bucket),
		Prefix:    aws.String(s.toValidS3DirPath(dir)),
		Delimiter: aws.String("/"),
	}
	ctx := context.TODO()
	objects, err := s3Client.ListObjectsV2(ctx, input)
	if err != nil {
		panic(err)
	}
	for _, commonPrefix := range objects.CommonPrefixes {
		dirs = append(dirs, *commonPrefix.Prefix)
	}
	return dirs
}

func (s *S3Storage) Files(dir string) []string {
	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}
	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(s.disk.Bucket),
		Prefix:    aws.String(s.toValidS3DirPath(dir)),
		Delimiter: aws.String("/"),
	}
	ctx := context.TODO()
	objects, err := s3Client.ListObjectsV2(ctx, input)
	if err != nil {
		panic(err)
	}

	files := []string{}
	for _, object := range objects.Contents {
		if s.toValidS3DirPath(dir) == *object.Key {
			continue
		}
		files = append(files, *object.Key)
	}
	return files
}

func (s *S3Storage) Exists(file string) bool {
	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s.disk.Bucket),
		Key:    aws.String(file),
	}
	ctx := context.TODO()

	_, err = s3Client.HeadObject(ctx, input)

	return err == nil
}

// func (r *S3) Get(file string) (string, error) {
// 	resp, err := r.instance.GetObject(r.ctx, &s3.GetObjectInput{
// 		Bucket: aws.String(r.bucket),
// 		Key:    aws.String(file),
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	data, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	if err := resp.Body.Close(); err != nil {
// 		return "", err
// 	}

// 	return string(data), nil
// }

func (s *S3Storage) MakeDirectory(directory string) error {
	if !strings.HasSuffix(directory, "/") {
		directory += "/"
	}

	return s.Put(directory, []byte(""))
}

func (s *S3Storage) Missing(file string) bool {
	return !s.Exists(file)
}

func (s *S3Storage) Move(oldFile, newFile string) error {
	if err := s.Copy(oldFile, newFile); err != nil {
		return err
	}

	return s.DeleteFile([]string{oldFile})
}

func (s *S3Storage) Put(filePath string, content []byte) error {
	mimeType := mimetype.Detect(content)

	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}
	cfmt.Successln("File upload: ", filePath)
	cfmt.Successln("Mimetype: ", mimeType)
	size := int64(len(content))
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.disk.Bucket),
		Key:    aws.String(filePath),
		Body:   strings.NewReader(string(content)),
		// ContentLength:      int64(len(content)),
		// ContentType:        aws.String(mtype.String()),
		// Body:               bytes.NewReader(buffer),
		ContentLength:      &size,
		ContentType:        aws.String(http.DetectContentType(content)),
		ContentDisposition: aws.String("attachment"),
		ACL:                types.ObjectCannedACLPublicRead,
		// ACL:                aws.String("public-read"),
	}

	_, err = s3Client.PutObject(context.TODO(), input)

	return err
}

func (s *S3Storage) PutFile(filePath string, source filesystem.File) (string, error) {
	return s.PutFileAs(filePath, source, utils.StrRandom(40))
}

func (s *S3Storage) PutFileAs(filePath string, source filesystem.File, name string) (string, error) {
	fullPath, err := fullPathOfFile(filePath, source, name)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(source.File())
	if err != nil {
		return "", err
	}

	if err := s.Put(fullPath, data); err != nil {
		return "", err
	}

	return fullPath, nil
}

func (s *S3Storage) Size(file string) int64 {
	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	resp, err := s3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.disk.Bucket),
		Key:    aws.String(file),
	})
	if err != nil {
		return -1
	}

	return *resp.ContentLength
}

func (s *S3Storage) LastModified(file string) time.Time {
	s3Client, err := s.client()
	if err != nil {
		panic(err)
	}
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s.disk.Bucket),
		Key:    aws.String(file),
	}
	ctx := context.TODO()
	resp, err := s3Client.HeadObject(ctx, input)
	if err != nil {
		return time.Time{}
	}

	l, err := time.LoadLocation("Europe/London")
	if err != nil {
		return time.Time{}
	}

	return aws.ToTime(resp.LastModified).In(l)
}

func (s *S3Storage) Url(file string) string {
	// cfmt.Errorln("Disk URL: ", s.disk.Url)
	// cfmt.Errorln("Disk Key: ", s.disk.Key)
	// cfmt.Errorln("Config Media Endpoint: ", config.MediaEndpoint)
	// cfmt.Errorln("Config Media Key: ", config.MediaKey)
	return strings.TrimSuffix(s.disk.Url, "/") + "/" + strings.TrimPrefix(file, "/")
}

// toValidS3DirPath trims "./" and "/" prefixes/suffixes from a given path and
// returns the resulting string. If the resulting string is not empty and
// doesn't end with "/", it appends "/" to the end.
//
// path string - The path to be sanitized.
// string - The sanitized path.
func (s *S3Storage) toValidS3DirPath(path string) string {
	realPath := strings.TrimPrefix(path, "./")
	realPath = strings.TrimPrefix(realPath, "/")
	realPath = strings.TrimPrefix(realPath, ".")
	if realPath != "" && !strings.HasSuffix(realPath, "/") {
		realPath += "/"
	}

	return realPath
}

func fullPathOfFile(filePath string, source filesystem.File, name string) (string, error) {
	extension := path.Ext(name)
	if extension == "" {
		var err error
		extension, err = file.Extension(source.File(), true)
		if err != nil {
			return "", err
		}

		return strings.TrimSuffix(filePath, "/") + "/" + strings.TrimSuffix(strings.TrimPrefix(path.Base(name), "/"), "/") + "." + extension, nil
	} else {
		return strings.TrimSuffix(filePath, "/") + "/" + strings.TrimPrefix(path.Base(name), "/"), nil
	}
}
