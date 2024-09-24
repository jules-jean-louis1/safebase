package controllers

import (
	"fmt"
	"os"
)

func GetSizeBackup(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("error getting file info: %w", err)
	}
	return fi.Size(), nil
}
