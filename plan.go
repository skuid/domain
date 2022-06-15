package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// This is the result of getting the plan from the pliny
// deployment retrieval plan endpoint
type NlxPlanPayload map[string]NlxPlan

const (
	WardenKey = "skuidCloudDataService"
	PlinyKey  = "skuidMetadataService"
)

// struct {
// 	// Cloud Data Service is WARDEN
// 	CloudDataService NlxPlan `json:"skuidCloudDataService"`
// 	// Metadata Service is PLINY
// 	MetadataService NlxPlan `json:"skuidMetadataService"`
// }

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
	AppName   string   `json:"appName"`
	PageNames []string `json:"pages"`
}

func GetDeployPlan(auth *Authorization) (duration time.Duration, result NlxPlanPayload, err error) {
	planStart := time.Now()
	defer func() { duration = time.Since(planStart) }()

	if result, err = FastJsonBodyRequest[NlxPlanPayload](
		fmt.Sprintf("%s/api/%v/metadata/deploy/plan", auth.Host, DEFAULT_API_VERSION),
		fasthttp.MethodPost,
		[]byte{},
		RequestHeaders{
			fasthttp.HeaderContentType:   DEFAULT_CONTENT_TYPE,
			fasthttp.HeaderAuthorization: fmt.Sprintf("Bearer %v", auth.AuthorizationToken),
		},
	); err != nil {
		return
	}

	return
}

func SplitPlans(plans NlxPlanPayload) (splitPlans NlxPlanPayload) {
	splitPlans = make(NlxPlanPayload)
	for k, p := range plans {
		key := func(item, keyType string) string {
			return fmt.Sprintf("%v-%v-%v", k, keyType, item)
		}
		m := p.Metadata
		for _, item := range m.Apps {
			v := p
			v.Metadata = NlxMetadata{Apps: []string{item}}
			splitPlans[key(item, "app")] = v
		}
		for _, item := range m.AuthProviders {
			v := p
			v.Metadata = NlxMetadata{AuthProviders: []string{item}}
			splitPlans[key(item, "authprovider")] = v
		}
		for _, item := range m.ComponentPacks {
			v := p
			v.Metadata = NlxMetadata{ComponentPacks: []string{item}}
			splitPlans[key(item, "componentpack")] = v
		}
		for _, item := range m.DataServices {
			v := p
			v.Metadata = NlxMetadata{DataServices: []string{item}}
			splitPlans[key(item, "dataservice")] = v
		}
		for _, item := range m.DataSources {
			v := p
			v.Metadata = NlxMetadata{DataSources: []string{item}}
			splitPlans[key(item, "datasource")] = v
		}
		for _, item := range m.DesignSystems {
			v := p
			v.Metadata = NlxMetadata{DesignSystems: []string{item}}
			splitPlans[key(item, "designsystem")] = v
		}
		for _, item := range m.Variables {
			v := p
			v.Metadata = NlxMetadata{Variables: []string{item}}
			splitPlans[key(item, "variable")] = v
		}
		for _, item := range m.Files {
			v := p
			v.Metadata = NlxMetadata{Files: []string{item}}
			splitPlans[key(item, "file")] = v
		}
		for _, item := range m.Pages {
			v := p
			v.Metadata = NlxMetadata{Pages: []string{item}}
			splitPlans[key(item, "page")] = v
		}
		for _, item := range m.PermissionSets {
			v := p
			v.Metadata = NlxMetadata{PermissionSets: []string{item}}
			splitPlans[key(item, "permissionset")] = v
		}
		for _, item := range m.Profiles {
			v := p
			v.Metadata = NlxMetadata{Profiles: []string{item}}
			splitPlans[key(item, "profile")] = v
		}
		for _, item := range m.Site {
			v := p
			v.Metadata = NlxMetadata{Site: []string{item}}
			splitPlans[key(item, "site")] = v
		}
		for _, item := range m.Themes {
			v := p
			v.Metadata = NlxMetadata{Themes: []string{item}}
			splitPlans[key(item, "theme")] = v
		}
	}
	return
}

// NewRetrievalRequestBody marshals the NlxMetadata into json and returns
// the body. This is the payload expected for the Retrieval Request
func NewRetrievalRequestBody(metadata NlxMetadata, zip bool) (body []byte) {
	// there should be no issue marshalling this thing.
	body, _ = json.Marshal(struct {
		Metadata NlxMetadata `json:"metadata"`
		DoZip    bool        `json:"zip"`
	}{
		Metadata: metadata,
		DoZip:    zip,
	})
	return
}
