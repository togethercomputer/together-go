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

func TestBetaRlOperationNewInferenceCheckpoint(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.NewInferenceCheckpoint(context.TODO(), "session_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlOperationNewTrainingCheckpoint(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.NewTrainingCheckpoint(context.TODO(), "session_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRlOperationCustomForwardBackward(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.CustomForwardBackward(
		context.TODO(),
		"session_id",
		together.BetaRlOperationCustomForwardBackwardParams{
			Gradients: []together.BetaRlOperationCustomForwardBackwardParamsGradient{{
				Data:  []float64{-0.1, 0.05, -0.08, 0.12, -0.03},
				Dtype: together.DTypeDTypeFloat32,
			}},
			Samples: []together.BetaRlOperationCustomForwardBackwardParamsSample{{
				LossFnInputs: map[string]together.TensorDataParam{
					"foo": {
						Data:              []float64{1, 2, 3},
						Dtype:             together.TensorDataDtypeInt64,
						Shape:             []int64{3},
						SparseColIndices:  []int64{0, 2},
						SparseCrowIndices: []int64{0, 2},
					},
				},
				ModelInput: together.ModelInputParam{
					Chunks: []together.ModelInputChunkParam{{
						EncodedText: together.EncodedTextChunkParam{
							Tokens: []together.EncodedTextChunkTokenUnionParam{{
								OfInt: together.Int(123),
							}, {
								OfInt: together.Int(456),
							}, {
								OfInt: together.Int(789),
							}},
						},
					}},
				},
				RoutedExperts: together.RoutedExpertsParam{
					Shape: []together.RoutedExpertsShapeUnionParam{{
						OfString: together.String("512"),
					}, {
						OfString: together.String("64"),
					}, {
						OfString: together.String("8"),
					}},
					Data:      together.String("U3RhaW5sZXNzIHJvY2tz"),
					ObjectUri: together.String("https://example.com"),
				},
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

func TestBetaRlOperationForward(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.Forward(
		context.TODO(),
		"session_id",
		together.BetaRlOperationForwardParams{
			Samples: []together.BetaRlOperationForwardParamsSample{{
				LossFnInputs: map[string]together.TensorDataParam{
					"foo": {
						Data:              []float64{1, 2, 3},
						Dtype:             together.TensorDataDtypeInt64,
						Shape:             []int64{3},
						SparseColIndices:  []int64{0, 2},
						SparseCrowIndices: []int64{0, 2},
					},
				},
				ModelInput: together.ModelInputParam{
					Chunks: []together.ModelInputChunkParam{{
						EncodedText: together.EncodedTextChunkParam{
							Tokens: []together.EncodedTextChunkTokenUnionParam{{
								OfInt: together.Int(123),
							}, {
								OfInt: together.Int(456),
							}, {
								OfInt: together.Int(789),
							}},
						},
					}},
				},
				RoutedExperts: together.RoutedExpertsParam{
					Shape: []together.RoutedExpertsShapeUnionParam{{
						OfString: together.String("512"),
					}, {
						OfString: together.String("64"),
					}, {
						OfString: together.String("8"),
					}},
					Data:      together.String("U3RhaW5sZXNzIHJvY2tz"),
					ObjectUri: together.String("https://example.com"),
				},
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

func TestBetaRlOperationForwardBackwardWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.ForwardBackward(
		context.TODO(),
		"session_id",
		together.BetaRlOperationForwardBackwardParams{
			Loss: together.LossConfigParam{
				Type: together.LossTypeLossTypeGrpo,
				CispoParams: together.CispoLossParams{
					ClipHighThreshold: together.Float(4),
					ClipLowThreshold:  together.Float(0),
				},
				CrossEntropyParams: together.CrossEntropyLossParams{},
				DroParams: together.DroLossParams{
					Beta: 0.05,
				},
				GrpoParams: together.GrpoLossParams{
					AggType:           together.GrpoLossAggregationTypeGrpoLossAggregationTypeFixedHorizon,
					Beta:              together.Float(0.1),
					ClipHighThreshold: together.Float(1.2),
					ClipLowThreshold:  together.Float(0.8),
					RatioType:         together.GrpoLossRatioTypeGrpoLossRatioTypeToken,
				},
				PpoParams: together.PpoLossParams{
					ClipHighThreshold: together.Float(1.2),
					ClipLowThreshold:  together.Float(0.8),
				},
			},
			Samples: []together.BetaRlOperationForwardBackwardParamsSample{{
				LossFnInputs: map[string]together.TensorDataParam{
					"foo": {
						Data:              []float64{1, 2, 3},
						Dtype:             together.TensorDataDtypeInt64,
						Shape:             []int64{3},
						SparseColIndices:  []int64{0, 2},
						SparseCrowIndices: []int64{0, 2},
					},
				},
				ModelInput: together.ModelInputParam{
					Chunks: []together.ModelInputChunkParam{{
						EncodedText: together.EncodedTextChunkParam{
							Tokens: []together.EncodedTextChunkTokenUnionParam{{
								OfInt: together.Int(123),
							}, {
								OfInt: together.Int(456),
							}, {
								OfInt: together.Int(789),
							}},
						},
					}},
				},
				RoutedExperts: together.RoutedExpertsParam{
					Shape: []together.RoutedExpertsShapeUnionParam{{
						OfString: together.String("512"),
					}, {
						OfString: together.String("64"),
					}, {
						OfString: together.String("8"),
					}},
					Data:      together.String("U3RhaW5sZXNzIHJvY2tz"),
					ObjectUri: together.String("https://example.com"),
				},
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

func TestBetaRlOperationOptimStepWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.OptimStep(
		context.TODO(),
		"session_id",
		together.BetaRlOperationOptimStepParams{
			AdamParams: together.AdamParams{
				Beta1:        together.Float(0.9),
				Beta2:        together.Float(0.95),
				Eps:          together.Float(1e-8),
				GradClipNorm: together.Float(10),
				LearningRate: together.Float(0.0001),
				WeightDecay:  together.Float(0.1),
			},
			MuonParams: together.MuonParams{
				Adam: together.AdamParams{
					Beta1:        together.Float(0.9),
					Beta2:        together.Float(0.95),
					Eps:          together.Float(1e-8),
					GradClipNorm: together.Float(10),
					LearningRate: together.Float(0.0001),
					WeightDecay:  together.Float(0.1),
				},
				GradClipNorm:      together.Float(10),
				LearningRate:      together.Float(0.02),
				Momentum:          together.Float(0.95),
				NewtonSchulzSteps: together.Int(5),
				WeightDecay:       together.Float(0),
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

func TestBetaRlOperationGetCustomForwardBackward(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.GetCustomForwardBackward(
		context.TODO(),
		"operation_id",
		together.BetaRlOperationGetCustomForwardBackwardParams{
			SessionID: "session_id",
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

func TestBetaRlOperationGetForward(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.GetForward(
		context.TODO(),
		"operation_id",
		together.BetaRlOperationGetForwardParams{
			SessionID: "session_id",
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

func TestBetaRlOperationGetForwardBackward(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.GetForwardBackward(
		context.TODO(),
		"operation_id",
		together.BetaRlOperationGetForwardBackwardParams{
			SessionID: "session_id",
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

func TestBetaRlOperationGetInferenceCheckpoint(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.GetInferenceCheckpoint(
		context.TODO(),
		"operation_id",
		together.BetaRlOperationGetInferenceCheckpointParams{
			SessionID: "session_id",
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

func TestBetaRlOperationGetOptimStep(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.GetOptimStep(
		context.TODO(),
		"operation_id",
		together.BetaRlOperationGetOptimStepParams{
			SessionID: "session_id",
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

func TestBetaRlOperationGetSample(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.GetSample(
		context.TODO(),
		"operation_id",
		together.BetaRlOperationGetSampleParams{
			SessionID: "session_id",
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

func TestBetaRlOperationGetTrainingCheckpoint(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.GetTrainingCheckpoint(
		context.TODO(),
		"operation_id",
		together.BetaRlOperationGetTrainingCheckpointParams{
			SessionID: "session_id",
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

func TestBetaRlOperationGetWeightsSync(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.GetWeightsSync(
		context.TODO(),
		"operation_id",
		together.BetaRlOperationGetWeightsSyncParams{
			SessionID: "session_id",
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

func TestBetaRlOperationSampleWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.Sample(
		context.TODO(),
		"session_id",
		together.BetaRlOperationSampleParams{
			ModelInputs: []together.ModelInputParam{{
				Chunks: []together.ModelInputChunkParam{{
					EncodedText: together.EncodedTextChunkParam{
						Tokens: []together.EncodedTextChunkTokenUnionParam{{
							OfInt: together.Int(123),
						}, {
							OfInt: together.Int(456),
						}, {
							OfInt: together.Int(789),
						}},
					},
				}},
			}},
			NumSamples:                   together.Int(1),
			PromptLogprobs:               together.Bool(false),
			ReturnRoutedExperts:          together.Bool(false),
			ReturnRoutedExpertsObjectUri: together.Bool(false),
			SamplingParams: together.SamplingParams{
				MaxTokens: together.Int(512),
				Seed: together.SamplingParamsSeedUnion{
					OfString: together.String("42"),
				},
				Stop:        []string{"\n", "END"},
				Temperature: together.Float(1),
				TopK:        together.Int(-1),
				TopP:        together.Float(1),
			},
			TopkPromptLogprobs: together.Int(0),
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

func TestBetaRlOperationWeightsSync(t *testing.T) {
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
	_, err := client.Beta.Rl.Operations.WeightsSync(
		context.TODO(),
		"session_id",
		together.BetaRlOperationWeightsSyncParams{
			WeightSyncType: together.WeightSyncTypeWeightSyncTypeSynchronous,
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
