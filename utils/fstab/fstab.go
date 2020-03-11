package fstab

import (
	"fmt"
	gofstab "github.com/deniswernert/go-fstab"
	"syscall"
)

type DiskStatus struct {
	All  uint64
	Free uint64
	Used uint64
}

func DiskUsage(path string) (disk DiskStatus, err error) {
	fs := syscall.statfs_t{}
	err0 := syscall.statfs()
	if err != nil {
		err = err0
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

func FstabMain() (err error) {
	mounts, err0 := gofstab.ParseSystem()
	if err0 != nil {
		err = err0
		fmt.Println("ferr: ", err)
		return
	}
	for _, val := range mounts {
		if val.File == "swap" || val.File == "/dev/shm" || val.File == "/dev/pts" || val.File == "/proc" || val.File == "/sys" {
			continue
		}
		fmt.Println("f: ", val.File)
		disk, err1 := DiskUsage(val.File)
		if err1 != nil {
			err = err1
			fmt.Println("derr", err)
			return
		}
		fmt.Println(disk)
	}
	return
}
