package storage

import (
	"github.com/StackExchange/wmi"
)

type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
}

type StorageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
}

func GetStorageInfo() (loaclStorages []Storage, err error) {
	var storageinfo []StorageInfo

	err = wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
	if err != nil {
		return
	}
	for _, storage := range storageinfo {
		info := Storage{
			Name:       storage.Name,
			FileSystem: storage.FileSystem,
			Total:      storage.Size,
			Free:       storage.FreeSpace,
		}
		loaclStorages = append(loaclStorages, info)
	}
	return
}
