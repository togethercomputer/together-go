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

func TestBetaEndpointRolloutNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.New(
		context.TODO(),
		"endpointId",
		together.BetaEndpointRolloutNewParams{
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

func TestBetaEndpointRolloutGet(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.Get(
		context.TODO(),
		"id",
		together.BetaEndpointRolloutGetParams{
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

func TestBetaEndpointRolloutListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.List(
		context.TODO(),
		"endpointId",
		together.BetaEndpointRolloutListParams{
			ProjectID: together.String("projectId"),
			After:     together.String("after"),
			Filter:    together.BetaEndpointRolloutListParamsFilterRolloutFilterActive,
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

func TestBetaEndpointRolloutDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.Delete(
		context.TODO(),
		"id",
		together.BetaEndpointRolloutDeleteParams{
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

func TestBetaEndpointRolloutCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.Cancel(
		context.TODO(),
		"id",
		together.BetaEndpointRolloutCancelParams{
			ProjectID:   together.String("projectId"),
			EndpointID:  "endpointId",
			Reason:      "reason",
			Disposition: together.BetaEndpointRolloutCancelParamsDispositionCancelDispositionFreeze,
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

func TestBetaEndpointRolloutPauseWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.Pause(
		context.TODO(),
		"id",
		together.BetaEndpointRolloutPauseParams{
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

func TestBetaEndpointRolloutPreviewDefaultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.PreviewDefaults(
		context.TODO(),
		"endpointId",
		together.BetaEndpointRolloutPreviewDefaultsParams{
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

func TestBetaEndpointRolloutPromoteWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.Promote(
		context.TODO(),
		"id",
		together.BetaEndpointRolloutPromoteParams{
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

func TestBetaEndpointRolloutResumeWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.Resume(
		context.TODO(),
		"id",
		together.BetaEndpointRolloutResumeParams{
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

func TestBetaEndpointRolloutStart(t *testing.T) {
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
	_, err := client.Beta.Endpoints.Rollouts.Start(
		context.TODO(),
		"id",
		together.BetaEndpointRolloutStartParams{
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
