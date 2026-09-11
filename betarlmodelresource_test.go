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

func TestBetaRlModelResourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.ModelResources.New(context.TODO(), together.BetaRlModelResourceNewParams{
		BaseModel: "Qwen/Qwen3-0.6B",
		ComputeConfig: together.BetaRlModelResourceNewParamsComputeConfig{
			GPUType:              "B200-SXM",
			NumGeneratorReplicas: together.Int(2),
		},
		LoraEnabled: together.Bool(true),
		OptimizerConfig: together.OptimizerConfigParam{
			Adam: together.OptimizerConfigAdamParam{},
			Muon: together.OptimizerConfigMuonParam{
				ScalingStrategy: together.MuonScalingStrategyMuonScalingStrategyOriginal,
			},
		},
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlModelResourceGet(t *testing.T) {
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
	_, err := client.Beta.Rl.ModelResources.Get(context.TODO(), "model_resources_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlModelResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.ModelResources.List(context.TODO(), together.BetaRlModelResourceListParams{
		After:     together.String("after"),
		CreatedBy: together.String("created_by"),
		Limit:     together.Int(0),
		Status:    []string{"MODEL_RESOURCES_STATUS_PENDING"},
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlModelResourceEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.ModelResources.EstimateCost(context.TODO(), together.BetaRlModelResourceEstimateCostParams{
		BaseModel: "Qwen/Qwen3-0.6B",
		ComputeConfig: together.BetaRlModelResourceEstimateCostParamsComputeConfig{
			GPUType:              "B200-SXM",
			NumGeneratorReplicas: together.Int(2),
		},
		LoraEnabled: together.Bool(true),
		OptimizerConfig: together.OptimizerConfigParam{
			Adam: together.OptimizerConfigAdamParam{},
			Muon: together.OptimizerConfigMuonParam{
				ScalingStrategy: together.MuonScalingStrategyMuonScalingStrategyOriginal,
			},
		},
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlModelResourceStopWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.ModelResources.Stop(
		context.TODO(),
		"model_resources_id",
		together.BetaRlModelResourceStopParams{
			Force: together.Bool(true),
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
