package domain_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/skuid/domain"
	"github.com/skuid/domain/util"
)

func TestGetDeployPlan(t *testing.T) {
	util.SkipIntegrationTest(t)

	auth, err := domain.Authorize(authHost, authUser, authPass)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	wd, _ := os.Getwd()
	fp := filepath.Join(wd, "..", "..", "_deploy")

	deploymentPlan, err := domain.Archive(fp, nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	duration, plans, err := domain.PrepareDeployment(auth, deploymentPlan, nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(duration)

	duration, _, err = domain.ExecuteDeployPlan(auth, plans, fp)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(duration)

}

func BenchmarkDeploymentPlan(b *testing.B) {
	util.SkipBenchmark(b)
	auth, _ := domain.Authorize(authHost, authUser, authPass)
	wd, _ := os.Getwd()
	fp := filepath.Join(wd, "..", "..", "_deploy")
	deploymentPlan, _ := domain.Archive(fp, nil)
	_, plans, _ := domain.PrepareDeployment(auth, deploymentPlan, nil)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, _ = domain.ExecuteDeployPlan(auth, plans, fp)
	}
}
