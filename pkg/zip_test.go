package pkg_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/skuid/domain/pkg"
	"github.com/skuid/domain/pkg/logging"
	"github.com/skuid/domain/pkg/util"
)

func TestZip(t *testing.T) {
	util.SkipIntegrationTest(t)
	cd, _ := os.Getwd()
	relpath := filepath.Join(cd, "..", "..", "_deploy")
	bb, err := pkg.Archive(relpath, nil)
	if err != nil {
		logging.Logger.Fatal(err)
	}
	logging.Logger.Info(len(bb))

}