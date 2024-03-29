package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// This is the result of getting the plan from the pliny
// deployment retrieval plan endpoint
type NlxPlanPayload struct {
	// Cloud Data Service is WARDEN
	CloudDataService *NlxPlan `json:"skuidCloudDataService"`
	// Metadata Service is PLINY
	MetadataService *NlxPlan `json:"skuidMetadataService"`
}

type NlxPlan struct {
	Host     string      `json:"host"`
	Port     string      `json:"port"`
	Endpoint string      `json:"url"`
	Type     string      `json:"type"`
	Metadata NlxMetadata `json:"metadata"`
	Warnings []string    `json:"warnings"`
}

// Serialize this and provide it with the
// request for retrieval
type NlxPlanFilter struct {
	AppName   string   `json:"appName,omitempty"`
	PageNames []string `json:"pages,omitempty"`
}

func GetDeployPlan(auth *Authorization) (duration time.Duration, result NlxPlanPayload, err error) {
	planStart := time.Now()
	defer func() { duration = time.Since(planStart) }()

	if result, err = JsonBodyRequest[NlxPlanPayload](
		fmt.Sprintf("%s/api/%v/metadata/deploy/plan", auth.Host, DEFAULT_API_VERSION),
		http.MethodPost,
		[]byte{},
		RequestHeaders{
			HeaderContentType:   DEFAULT_CONTENT_TYPE,
			HeaderAuthorization: fmt.Sprintf("Bearer %v", auth.AuthorizationToken),
		},
	); err != nil {
		return
	}

	return
}

// NewRetrievalRequestBody marshals the NlxMetadata into json and returns
// the body. This is the payload expected for the Retrieval Request
func NewRetrievalRequestBody(metadata NlxMetadata) (body []byte) {
	// there should be no issue marshalling this thing.
	body, _ = json.Marshal(struct {
		Metadata NlxMetadata `json:"metadata"`
	}{
		Metadata: metadata,
	})
	return
}
