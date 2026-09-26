//go:build linux

package adaptive

import (
	"fmt"
	"math"
	"syscall"
)

func GetDiskInfo(directory string) (DiskInfo, error) {
	var stat syscall.Statfs_t

	if err := syscall.Statfs(directory, &stat); err != nil {
		return DiskInfo{}, fmt.Errorf("failed to get filesystem statistics: %w", err)
	}

	totalBytes := stat.Blocks * uint64(stat.Bsize)
	freeBytes := stat.Bavail * uint64(stat.Bsize)

	if totalBytes == 0 {
		return DiskInfo{}, fmt.Errorf("filesystem total size is zero")
	}

	freePercent := (float64(freeBytes) / float64(totalBytes)) * 100

	return DiskInfo{
		FreeBytes:   freeBytes,
		TotalBytes:  totalBytes,
		FreePercent: math.Round(freePercent*100) / 100,
	}, nil
}
