package domain

import (
	"os"
	"path/filepath"

	"github.com/gookit/color"

	"github.com/skuid/domain/logging"
)

func ClearDirectories(targetDir string) {
	for _, dirName := range GetMetadataTypeDirNames() {
		dirPath := filepath.Join(targetDir, dirName)
		logging.Get().Debugf("%v: %v", color.Yellow.Sprint("Removing"), dirPath)
		os.RemoveAll(dirPath)
	}
}
