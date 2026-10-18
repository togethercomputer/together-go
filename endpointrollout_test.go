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

func TestEndpointRolloutNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.New(
		context.TODO(),
		"endpointId",
		together.EndpointRolloutNewParams{
			ProjectID:          together.String("projectId"),
			SourceDeploymentID: "dep_source123",
			TargetDeploymentID: "dep_target456",
			BlueGreen:          together.BlueGreenConfigParam{},
			Canary: together.CanaryConfigParam{
				StepInterval: together.String("300s"),
				Steps: []together.RolloutStepParam{{
					Traffic:  25,
					Replicas: together.Int(0),
				}, {
					Traffic:  50,
					Replicas: together.Int(0),
				}, {
					Traffic:  100,
					Replicas: together.Int(0),
				}},
			},
			FinalSourceReplicas: together.Int(0),
			FinalTargetReplicas: together.Int(0),
			Metrics: []together.MetricRuleParam{{
				Name:       together.MetricRuleNameRouterLatency,
				Percentile: together.Int(95),
				RegressionCheck: together.RegressionCheckParam{
					Direction:            together.RegressionCheckDirectionRegressionDirectionHigherIsWorse,
					MaxRegressionPercent: together.Float(0),
				},
				Stat: together.MetricRuleStatMetricStatTypePercentile,
				ThresholdCheck: together.ThresholdCheckParam{
					Operator: together.ThresholdCheckOperatorThresholdOperatorLt,
					Value:    together.Float(30000),
				},
				Window: together.String("300s"),
			}},
			Rolling: together.RollingConfigParam{},
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

func TestEndpointRolloutGet(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.Get(
		context.TODO(),
		"id",
		together.EndpointRolloutGetParams{
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

func TestEndpointRolloutListWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.List(
		context.TODO(),
		"endpointId",
		together.EndpointRolloutListParams{
			ProjectID: together.String("projectId"),
			After:     together.String("after"),
			Filter:    together.EndpointRolloutListParamsFilterRolloutFilterActive,
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

func TestEndpointRolloutDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.Delete(
		context.TODO(),
		"id",
		together.EndpointRolloutDeleteParams{
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

func TestEndpointRolloutCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.Cancel(
		context.TODO(),
		"id",
		together.EndpointRolloutCancelParams{
			ProjectID:   together.String("projectId"),
			EndpointID:  "endpointId",
			Reason:      "reason",
			Disposition: together.EndpointRolloutCancelParamsDispositionCancelDispositionFreeze,
			Etag:        together.String("etag"),
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

func TestEndpointRolloutPauseWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.Pause(
		context.TODO(),
		"id",
		together.EndpointRolloutPauseParams{
			ProjectID:  together.String("projectId"),
			EndpointID: "endpointId",
			Etag:       together.String("etag"),
			Reason:     together.String("reason"),
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

func TestEndpointRolloutPreviewDefaultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.PreviewDefaults(
		context.TODO(),
		"endpointId",
		together.EndpointRolloutPreviewDefaultsParams{
			ProjectID:          together.String("projectId"),
			SourceDeploymentID: "dep_source123",
			TargetDeploymentID: "dep_target456",
			BlueGreen:          together.BlueGreenConfigParam{},
			Canary: together.CanaryConfigParam{
				StepInterval: together.String("300s"),
				Steps: []together.RolloutStepParam{{
					Traffic:  25,
					Replicas: together.Int(0),
				}, {
					Traffic:  50,
					Replicas: together.Int(0),
				}, {
					Traffic:  100,
					Replicas: together.Int(0),
				}},
			},
			FinalSourceReplicas: together.Int(0),
			FinalTargetReplicas: together.Int(0),
			Metrics: []together.MetricRuleParam{{
				Name:       together.MetricRuleNameRouterLatency,
				Percentile: together.Int(95),
				RegressionCheck: together.RegressionCheckParam{
					Direction:            together.RegressionCheckDirectionRegressionDirectionHigherIsWorse,
					MaxRegressionPercent: together.Float(0),
				},
				Stat: together.MetricRuleStatMetricStatTypePercentile,
				ThresholdCheck: together.ThresholdCheckParam{
					Operator: together.ThresholdCheckOperatorThresholdOperatorLt,
					Value:    together.Float(30000),
				},
				Window: together.String("300s"),
			}},
			Rolling: together.RollingConfigParam{},
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

func TestEndpointRolloutPromoteWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.Promote(
		context.TODO(),
		"id",
		together.EndpointRolloutPromoteParams{
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

func TestEndpointRolloutResumeWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.Resume(
		context.TODO(),
		"id",
		together.EndpointRolloutResumeParams{
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

func TestEndpointRolloutStart(t *testing.T) {
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
	_, err := client.Endpoints.Rollouts.Start(
		context.TODO(),
		"id",
		together.EndpointRolloutStartParams{
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
