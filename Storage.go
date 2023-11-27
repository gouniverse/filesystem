package filesystem

import (
	// "project/config"

	"github.com/samber/lo"
)

func NewStorage(disk Disk) StorageInterface {

	// var disks = map[string]Disk{
	// 	DEFAULT: {
	// 		DiskName: DEFAULT,
	// 		Driver:   DRIVER_S3,
	// 		// Url:                  "https://" + config.MediaEndpoint,
	// 		Url:                  config.MediaUrl,
	// 		Region:               config.MediaRegion,
	// 		Key:                  config.MediaKey,
	// 		Secret:               config.MediaSecret,
	// 		UsePathStyleEndpoint: true,
	// 	},
	// 	CDN: {
	// 		DiskName: CDN,
	// 		Driver:   DRIVER_STATIC,
	// 		Url:      config.CDN_MEDIA_URL,
	// 	},
	// }

	// disk := lo.ValueOr(disks, diskName, Disk{})

	if lo.IsEmpty(disk) {
		return nil
	}

	if disk.Driver == DRIVER_S3 {
		storage := &S3Storage{
			disk: disk,
		}
		return storage
	}

	if disk.Driver == DRIVER_STATIC {
		storage := &StaticStorage{
			disk: disk,
		}
		return storage
	}

	//storage := &OsStorage{disk: diskName}
	return nil

}
