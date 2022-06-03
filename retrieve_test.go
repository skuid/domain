package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/skuid/domain"
	"github.com/skuid/domain/util"
)

func TestRetrievePlan(t *testing.T) {
	util.SkipIntegrationTest(t)

	auth, err := domain.Authorize(authHost, authUser, authPass)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		description string

		givenFilter *domain.NlxPlanFilter

		expectedError string
	}{
		{},
	} {
		t.Run(tc.description, func(t *testing.T) {
			duration, result, err := domain.GetRetrievePlan(auth, tc.givenFilter)
			t.Log(duration)
			t.Log(err)

			data, err := json.Marshal(result)
			t.Log(string(data))
			t.Log(err)

			for _, plan := range []domain.NlxPlan{
				result.CloudDataService, result.MetadataService,
			} {
				t.Logf("PLAN (%v):", plan.Type)
				b, _ := json.MarshalIndent(plan, "", " ")
				t.Log(string(b))
			}
		})
	}
}

func TestExecuteRetrieval(t *testing.T) {
	util.SkipIntegrationTest(t)
	auth, err := domain.Authorize(authHost, authUser, authPass)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		description string

		expectedError string
	}{
		{},
	} {
		t.Run(tc.description, func(t *testing.T) {
			duration, plans, err := domain.GetRetrievePlan(auth, nil)
			t.Log(duration)
			t.Log(plans)
			t.Log(err)

			duration, results, err := domain.ExecuteRetrieval(auth, plans, false)
			t.Log(duration)
			t.Log(results)
			t.Log(err)
		})
	}
}
