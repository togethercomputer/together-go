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

func TestFineTuningNewWithOptionalParams(t *testing.T) {
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
	_, err := client.FineTuning.New(context.TODO(), together.FineTuningNewParams{
		Model:        "model",
		TrainingFile: "training_file",
		BatchSize: together.FineTuningNewParamsBatchSizeUnion{
			OfInt: together.Int(0),
		},
		FromCheckpoint:   together.String("from_checkpoint"),
		FromHfModel:      together.String("from_hf_model"),
		HfAPIToken:       together.String("hf_api_token"),
		HfModelRevision:  together.String("hf_model_revision"),
		HfOutputRepoName: together.String("hf_output_repo_name"),
		LearningRate:     together.Float(0),
		LrScheduler: together.LrSchedulerParam{
			LrSchedulerType: together.LrSchedulerLrSchedulerTypeLinear,
			LrSchedulerArgs: together.LrSchedulerLrSchedulerArgsUnionParam{
				OfLinearLrSchedulerArgs: &together.LinearLrSchedulerArgsParam{
					MinLrRatio: together.Float(0),
				},
			},
		},
		MaxGradNorm:  together.Float(0),
		NCheckpoints: together.Int(0),
		NEpochs:      together.Int(0),
		NEvals:       together.Int(0),
		Suffix:       together.String("suffix"),
		TrainOnInputs: together.FineTuningNewParamsTrainOnInputsUnion{
			OfBool: together.Bool(true),
		},
		TrainingMethod: together.FineTuningNewParamsTrainingMethodUnion{
			OfTrainingMethodSft: &together.TrainingMethodSftParam{
				Method: together.TrainingMethodSftMethodSft,
				TrainOnInputs: together.TrainingMethodSftTrainOnInputsUnionParam{
					OfBool: together.Bool(true),
				},
			},
		},
		TrainingType: together.FineTuningNewParamsTrainingTypeUnion{
			OfFullTrainingType: &together.FullTrainingTypeParam{
				Type: together.FullTrainingTypeTypeFull,
			},
		},
		ValidationFile:   together.String("validation_file"),
		WandbAPIKey:      together.String("wandb_api_key"),
		WandbBaseURL:     together.String("wandb_base_url"),
		WandbName:        together.String("wandb_name"),
		WandbProjectName: together.String("wandb_project_name"),
		WarmupRatio:      together.Float(0),
		WeightDecay:      together.Float(0),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuningGet(t *testing.T) {
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
	_, err := client.FineTuning.Get(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuningList(t *testing.T) {
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
	_, err := client.FineTuning.List(context.TODO())
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuningDelete(t *testing.T) {
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
	_, err := client.FineTuning.Delete(
		context.TODO(),
		"id",
		together.FineTuningDeleteParams{
			Force: true,
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

func TestFineTuningCancel(t *testing.T) {
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
	_, err := client.FineTuning.Cancel(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuningDownloadWithOptionalParams(t *testing.T) {
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
	_, err := client.FineTuning.Download(context.TODO(), together.FineTuningDownloadParams{
		FtID:           "ft_id",
		Checkpoint:     together.FineTuningDownloadParamsCheckpointMerged,
		CheckpointStep: together.Int(0),
		Output:         together.String("output"),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuningListCheckpoints(t *testing.T) {
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
	_, err := client.FineTuning.ListCheckpoints(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuningListEvents(t *testing.T) {
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
	_, err := client.FineTuning.ListEvents(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
