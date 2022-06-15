package domain_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/skuid/domain"
	"github.com/skuid/domain/logging"
	"github.com/skuid/domain/util"
)

func TestZip(t *testing.T) {
	util.SkipIntegrationTest(t)
	cd, _ := os.Getwd()
	relpath := filepath.Join(cd, "..", "..", "_deploy")
	bb, err := domain.Archive(relpath, nil)
	if err != nil {
		logging.Get().Fatal(err)
	}
	logging.Get().Info(len(bb))

}
