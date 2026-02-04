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

func TestRlTrainingSessionNewWithOptionalParams(t *testing.T) {
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
	err := client.Rl.TrainingSessions.New(context.TODO(), together.RlTrainingSessionNewParams{
		Body: together.RlTrainingSessionNewParamsBody{
			BaseModel:    "meta-llama/Meta-Llama-3-8B-Instruct",
			CheckpointID: together.String("checkpoint-123"),
			LoraConfig: together.RlTrainingSessionNewParamsBodyLoraConfig{
				Alpha:   together.Int(0),
				Dropout: together.Float(0),
				Rank:    together.Int(0),
			},
			LrSchedulerConfig: together.RlTrainingSessionNewParamsBodyLrSchedulerConfig{
				Linear: together.RlTrainingSessionNewParamsBodyLrSchedulerConfigLinear{
					Params: together.RlTrainingSessionNewParamsBodyLrSchedulerConfigLinearParams{
						LrMin:       together.Float(0),
						WarmupSteps: together.Int(0),
					},
				},
			},
			OptimizerConfig: together.RlTrainingSessionNewParamsBodyOptimizerConfig{
				Adamw: together.RlTrainingSessionNewParamsBodyOptimizerConfigAdamw{
					Params: together.RlTrainingSessionNewParamsBodyOptimizerConfigAdamwParams{
						Beta1:       together.Float(0),
						Beta2:       together.Float(0),
						Eps:         together.Float(0),
						Lr:          together.Float(0),
						WeightDecay: together.Float(0),
					},
				},
				MaxGradNorm: together.Float(0),
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

func TestRlTrainingSessionGet(t *testing.T) {
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
	err := client.Rl.TrainingSessions.Get(context.TODO(), "session_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRlTrainingSessionListWithOptionalParams(t *testing.T) {
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
	err := client.Rl.TrainingSessions.List(context.TODO(), together.RlTrainingSessionListParams{
		Limit:  together.String("limit"),
		Offset: together.String("offset"),
		Status: together.String("status"),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRlTrainingSessionForwardBackward(t *testing.T) {
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
	err := client.Rl.TrainingSessions.ForwardBackward(
		context.TODO(),
		"session_id",
		together.RlTrainingSessionForwardBackwardParams{
			Body: together.RlTrainingSessionForwardBackwardParamsBody{
				LossFn: "LOSS_FN_GRPO",
				Samples: []together.RlTrainingSessionForwardBackwardParamsBodySample{{
					LossFnInputs: together.RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputs{
						TargetTokens: together.RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsTargetTokens{
							Data:  []float64{123, 456, 789},
							Dtype: "D_TYPE_INT64",
						},
						Weights: together.RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsWeights{
							Data:  []float64{0.1, 0.2, 0.3},
							Dtype: "D_TYPE_FLOAT32",
						},
					},
					ModelInput: together.RlTrainingSessionForwardBackwardParamsBodySampleModelInput{
						Chunks: []together.RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunk{{
							EncodedText: together.RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunkEncodedText{
								Tokens: []int64{123, 456, 789},
							},
						}},
					},
				}},
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

func TestRlTrainingSessionOptimStep(t *testing.T) {
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
	err := client.Rl.TrainingSessions.OptimStep(
		context.TODO(),
		"session_id",
		together.RlTrainingSessionOptimStepParams{
			Body: map[string]any{},
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

func TestRlTrainingSessionStop(t *testing.T) {
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
	err := client.Rl.TrainingSessions.Stop(context.TODO(), "session_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
