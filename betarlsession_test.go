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

func TestBetaRlSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.Sessions.New(context.TODO(), together.BetaRlSessionNewParams{
		ModelResourcesID: "123e4567-e89b-12d3-a456-426614174000",
		DisplayName:      together.String("gsm8k-experiment-2"),
		LoraConfig: together.LoraConfigParam{
			Alpha:   together.Int(64),
			Dropout: together.Float(0),
			Rank:    together.Int(32),
		},
		Metadata: together.BetaRlSessionNewParamsMetadata{
			Wandb: together.BetaRlSessionNewParamsMetadataWandb{
				Entity:  together.String("example-org"),
				Group:   together.String("gsm8k-35b-sweep"),
				Project: together.String("grpo-gsm8k"),
				RunID:   together.String("abc123"),
				RunName: together.String("exp2-thinking-4k-ctx"),
				URL:     together.String("https://wandb.ai/example-org/example-project/runs/run-id"),
			},
		},
		ResumeFromCheckpointID: together.String("123e4567-e89b-12d3-a456-426614174000"),
		ResumeFromHfCheckpoint: together.String("your-org/llama-3-8b-finetuned"),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlSessionGet(t *testing.T) {
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
	_, err := client.Beta.Rl.Sessions.Get(context.TODO(), "session_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlSessionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.Sessions.List(context.TODO(), together.BetaRlSessionListParams{
		After:            together.String("after"),
		CreatedBy:        together.String("created_by"),
		Limit:            together.Int(0),
		ModelResourcesID: together.String("model_resources_id"),
		Status:           []string{"TRAINING_SESSION_STATUS_CREATING"},
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlSessionStop(t *testing.T) {
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
	_, err := client.Beta.Rl.Sessions.Stop(context.TODO(), "session_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
