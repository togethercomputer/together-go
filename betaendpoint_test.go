// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/togethercomputer/together-go"
	"github.com/togethercomputer/together-go/internal/testutil"
	"github.com/togethercomputer/together-go/option"
)

func TestBetaEndpointNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Endpoints.New(context.TODO(), together.BetaEndpointNewParams{
		ProjectID:  together.String("projectId"),
		Name:       "name",
		Visibility: together.BetaEndpointNewParamsVisibilityVisibilityPrivate,
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaEndpointGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Endpoints.Get(
		context.TODO(),
		"id",
		together.BetaEndpointGetParams{
			ProjectID: together.String("projectId"),
		},
	)
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaEndpointUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Endpoints.Update(
		context.TODO(),
		"id",
		together.BetaEndpointUpdateParams{
			ProjectID:  together.String("projectId"),
			UpdateMask: together.String("updateMask"),
			Etag:       together.String("etag"),
			Name:       together.String("name"),
			TrafficSplit: []together.EndpointTrafficSplitEntryParam{{
				DeploymentID: "deploymentId",
				Weight:       0,
			}},
			Visibility: together.BetaEndpointUpdateParamsVisibilityVisibilityPrivate,
		},
	)
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaEndpointListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Endpoints.List(context.TODO(), together.BetaEndpointListParams{
		ProjectID: together.String("projectId"),
		After:     together.String("after"),
		Filter:    together.String("filter"),
		Limit:     together.Int(0),
		OrderBy:   together.String("orderBy"),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaEndpointDeleteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Endpoints.Delete(
		context.TODO(),
		"id",
		together.BetaEndpointDeleteParams{
			ProjectID: together.String("projectId"),
			Etag:      together.String("etag"),
		},
	)
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaEndpointAnalyticsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Endpoints.Analytics(
		context.TODO(),
		"id",
		together.BetaEndpointAnalyticsParams{
			ProjectID:         together.String("projectId"),
			DeploymentID:      together.String("deploymentId"),
			EndTime:           together.Time(time.Now()),
			Granularity:       together.String("granularity"),
			IncludeTimeSeries: together.Bool(true),
			StartTime:         together.Time(time.Now()),
		},
	)
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaEndpointListEventsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Endpoints.ListEvents(
		context.TODO(),
		"id",
		together.BetaEndpointListEventsParams{
			ProjectID:     together.String("projectId"),
			After:         together.String("after"),
			DeploymentIDs: []string{"string"},
			Limit:         together.Int(0),
			MinLevel:      together.BetaEndpointListEventsParamsMinLevelLevelDebug,
			Since:         together.Time(time.Now()),
			SourceKinds:   []string{"SOURCE_KIND_ENDPOINT"},
			SubjectID:     together.String("subjectId"),
			Types:         []string{"string"},
			Until:         together.Time(time.Now()),
		},
	)
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaEndpointListOrgScopedWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Endpoints.ListOrgScoped(
		context.TODO(),
		"organizationId",
		together.BetaEndpointListOrgScopedParams{
			After:   together.String("after"),
			Filter:  together.String("filter"),
			Limit:   together.Int(0),
			OrderBy: together.String("orderBy"),
		},
	)
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
