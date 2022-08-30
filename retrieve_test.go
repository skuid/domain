package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

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
			_, result, err := domain.GetRetrievePlan(auth, tc.givenFilter)
			assert.NoError(t, err)

			data, err := json.Marshal(result)
			t.Log(string(data))
			assert.NoError(t, err)
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

			duration, results, err := domain.ExecuteRetrieval(auth, plans)
			t.Log(duration)
			t.Log(results)
			t.Log(err)
		})
	}
}
