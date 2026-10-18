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

func TestBetaEndpointShadowExperimentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.ShadowExperiments.New(
		context.TODO(),
		"endpointId",
		together.BetaEndpointShadowExperimentNewParams{
			ProjectID: together.String("projectId"),
			Name:      "name",
			Source: together.ShadowSourceParam{
				Endpoint: together.ShadowEndpointSourceParam{
					Sampling: together.ShadowEndpointSourceSamplingUnionParam{
						OfUniform: &together.ShadowEndpointSourceSamplingUniformParam{
							Uniform: together.ShadowUniformSamplingParam{
								Rate: 0,
							},
						},
					},
				},
			},
			Targets: []together.BetaEndpointShadowExperimentNewParamsTarget{{
				Name:               "name",
				TargetDeploymentID: "targetDeploymentId",
				Description:        together.String("description"),
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

func TestBetaEndpointShadowExperimentGet(t *testing.T) {
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
	_, err := client.Beta.Endpoints.ShadowExperiments.Get(
		context.TODO(),
		"id",
		together.BetaEndpointShadowExperimentGetParams{
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

func TestBetaEndpointShadowExperimentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.ShadowExperiments.Update(
		context.TODO(),
		"id",
		together.BetaEndpointShadowExperimentUpdateParams{
			ProjectID:   together.String("projectId"),
			EndpointID:  "endpointId",
			UpdateMask:  "updateMask",
			Description: together.String("description"),
			Etag:        together.String("etag"),
			Source: together.ShadowSourceParam{
				Endpoint: together.ShadowEndpointSourceParam{
					Sampling: together.ShadowEndpointSourceSamplingUnionParam{
						OfUniform: &together.ShadowEndpointSourceSamplingUniformParam{
							Uniform: together.ShadowUniformSamplingParam{
								Rate: 0,
							},
						},
					},
				},
			},
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

func TestBetaEndpointShadowExperimentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.ShadowExperiments.List(
		context.TODO(),
		"endpointId",
		together.BetaEndpointShadowExperimentListParams{
			ProjectID:      together.String("projectId"),
			After:          together.String("after"),
			IncludeTargets: together.Bool(true),
			Limit:          together.Int(0),
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

func TestBetaEndpointShadowExperimentDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.ShadowExperiments.Delete(
		context.TODO(),
		"id",
		together.BetaEndpointShadowExperimentDeleteParams{
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
