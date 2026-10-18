// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/togethercomputer/together-go"
	"github.com/togethercomputer/together-go/internal/testutil"
	"github.com/togethercomputer/together-go/option"
)

func TestEndpointAbExperimentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.AbExperiments.New(
		context.TODO(),
		"endpointId",
		together.EndpointAbExperimentNewParams{
			ProjectID: together.String("projectId"),
			Members: []together.AbMemberParam{{
				DeploymentID: "deploymentId",
				Percent:      0,
				Role:         together.AbMemberRoleAbExperimentMemberRoleControl,
			}},
			Name:        "name",
			Description: together.String("description"),
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

func TestEndpointAbExperimentGet(t *testing.T) {
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
	_, err := client.Endpoints.AbExperiments.Get(
		context.TODO(),
		"id",
		together.EndpointAbExperimentGetParams{
			ProjectID:  together.String("projectId"),
			EndpointID: "endpointId",
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

func TestEndpointAbExperimentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.AbExperiments.Update(
		context.TODO(),
		"id",
		together.EndpointAbExperimentUpdateParams{
			ProjectID:   together.String("projectId"),
			EndpointID:  "endpointId",
			UpdateMask:  together.String("updateMask"),
			Description: together.String("description"),
			Etag:        together.String("etag"),
			Members: []together.AbMemberParam{{
				DeploymentID: "deploymentId",
				Percent:      0,
				Role:         together.AbMemberRoleAbExperimentMemberRoleControl,
			}},
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

func TestEndpointAbExperimentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.AbExperiments.List(
		context.TODO(),
		"endpointId",
		together.EndpointAbExperimentListParams{
			ProjectID: together.String("projectId"),
			After:     together.String("after"),
			Limit:     together.Int(0),
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

func TestEndpointAbExperimentDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.AbExperiments.Delete(
		context.TODO(),
		"id",
		together.EndpointAbExperimentDeleteParams{
			ProjectID:  together.String("projectId"),
			EndpointID: "endpointId",
			Etag:       together.String("etag"),
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
