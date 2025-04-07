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
	"github.com/togethercomputer/together-go/shared"
)

func TestFineTuneNewWithOptionalParams(t *testing.T) {
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
	_, err := client.FineTune.New(context.TODO(), together.FineTuneNewParams{
		Model:          together.F("model"),
		TrainingFile:   together.F("training_file"),
		BatchSize:      together.F[together.FineTuneNewParamsBatchSizeUnion](shared.UnionInt(int64(0))),
		FromCheckpoint: together.F("from_checkpoint"),
		LearningRate:   together.F(0.000000),
		LrScheduler: together.F(together.FineTuneNewParamsLrScheduler{
			LrSchedulerType: together.F(together.FineTuneNewParamsLrSchedulerLrSchedulerTypeLinear),
			LrSchedulerArgs: together.F[together.FineTuneNewParamsLrSchedulerLrSchedulerArgsUnion](together.FineTuneNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs{
				MinLrRatio: together.F(0.000000),
			}),
		}),
		MaxGradNorm:   together.F(0.000000),
		NCheckpoints:  together.F(int64(0)),
		NEpochs:       together.F(int64(0)),
		NEvals:        together.F(int64(0)),
		Suffix:        together.F("suffix"),
		TrainOnInputs: together.F[together.FineTuneNewParamsTrainOnInputsUnion](shared.UnionBool(true)),
		TrainingMethod: together.F[together.FineTuneNewParamsTrainingMethodUnion](together.FineTuneNewParamsTrainingMethodTrainingMethodSft{
			Method: together.F(together.FineTuneNewParamsTrainingMethodTrainingMethodSftMethodSft),
		}),
		TrainingType: together.F[together.FineTuneNewParamsTrainingTypeUnion](together.FineTuneNewParamsTrainingTypeFullTrainingType{
			Type: together.F(together.FineTuneNewParamsTrainingTypeFullTrainingTypeTypeFull),
		}),
		ValidationFile:   together.F("validation_file"),
		WandbAPIKey:      together.F("wandb_api_key"),
		WandbBaseURL:     together.F("wandb_base_url"),
		WandbName:        together.F("wandb_name"),
		WandbProjectName: together.F("wandb_project_name"),
		WarmupRatio:      together.F(0.000000),
		WeightDecay:      together.F(0.000000),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuneGet(t *testing.T) {
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
	_, err := client.FineTune.Get(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuneList(t *testing.T) {
	t.Skip("invalid oneOf in required props")
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
	_, err := client.FineTune.List(context.TODO())
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuneCancel(t *testing.T) {
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
	_, err := client.FineTune.Cancel(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuneDownloadWithOptionalParams(t *testing.T) {
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
	_, err := client.FineTune.Download(context.TODO(), together.FineTuneDownloadParams{
		FtID:           together.F("ft_id"),
		Checkpoint:     together.F(together.FineTuneDownloadParamsCheckpointMerged),
		CheckpointStep: together.F(int64(0)),
		Output:         together.F("output"),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuneListEvents(t *testing.T) {
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
	_, err := client.FineTune.ListEvents(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
