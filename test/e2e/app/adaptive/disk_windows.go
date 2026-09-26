package adaptive

import (
	"fmt"
	"math"

	"golang.org/x/sys/windows"
)

func GetDiskInfo(directory string) (DiskInfo, error) {
	path, err := windows.UTF16PtrFromString(directory)
	if err != nil {
		return DiskInfo{}, fmt.Errorf("failed to convert path: %w", err)
	}

	var freeBytesAvailable uint64
	var totalBytes uint64
	var totalFreeBytes uint64

	err = windows.GetDiskFreeSpaceEx(
		path,
		&freeBytesAvailable,
		&totalBytes,
		&totalFreeBytes,
	)
	if err != nil {
		return DiskInfo{}, fmt.Errorf("failed to get disk space: %w", err)
	}

	if totalBytes == 0 {
		return DiskInfo{}, fmt.Errorf("disk total size is zero")
	}

	freePercent := (float64(freeBytesAvailable) / float64(totalBytes)) * 100

	return DiskInfo{
		FreeBytes:   freeBytesAvailable,
		TotalBytes:  totalBytes,
		FreePercent: math.Round(freePercent*100) / 100,
	}, nil
}
