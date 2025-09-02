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

func TestEvaluationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Evaluation.New(context.TODO(), together.EvaluationNewParams{
		Parameters: together.F[together.EvaluationNewParamsParametersUnion](together.EvaluationNewParamsParametersEvaluationClassifyParameters{
			InputDataFilePath: together.F("file-abcd-1234"),
			Judge: together.F(together.EvaluationJudgeModelConfigParam{
				ModelName:      together.F("meta-llama/Llama-3-70B-Instruct-Turbo"),
				SystemTemplate: together.F("You are a helpful assistant that classifies text."),
			}),
			Labels:          together.F([]string{"Toxic", "Non-Toxic"}),
			PassLabels:      together.F([]string{"Non-Toxic"}),
			ModelToEvaluate: together.F[together.EvaluationNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion](shared.UnionString("output_column")),
		}),
		Type: together.F(together.EvaluationNewParamsTypeClassify),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvaluationGet(t *testing.T) {
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
	_, err := client.Evaluation.Get(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvaluationGetStatus(t *testing.T) {
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
	_, err := client.Evaluation.GetStatus(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvaluationUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Evaluation.UpdateStatus(
		context.TODO(),
		"id",
		together.EvaluationUpdateStatusParams{
			Status: together.F(together.EvaluationUpdateStatusParamsStatusCompleted),
			Error:  together.F("error"),
			Results: together.F[together.EvaluationUpdateStatusParamsResultsUnion](together.EvaluationUpdateStatusParamsResultsEvaluationClassifyResults{
				GenerationFailCount: together.F(0.000000),
				InvalidLabelCount:   together.F(0.000000),
				JudgeFailCount:      together.F(0.000000),
				LabelCounts:         together.F(`{"yes": 10, "no": 0}`),
				PassPercentage:      together.F(10.000000),
				ResultFileID:        together.F("file-1234-aefd"),
			}),
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
