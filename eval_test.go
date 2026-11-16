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

func TestEvalNewWithOptionalParams(t *testing.T) {
	t.Skip("Skipping test while we work on the api")
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
	_, err := client.Evals.New(context.TODO(), together.EvalNewParams{
		Parameters: together.EvalNewParamsParametersUnion{
			OfEvalNewsParametersEvaluationClassifyParameters: &together.EvalNewParamsParametersEvaluationClassifyParameters{
				InputDataFilePath: "input_data_file_path",
				Judge: together.EvalNewParamsParametersEvaluationClassifyParametersJudge{
					Model:            "model",
					ModelSource:      "serverless",
					SystemTemplate:   "system_template",
					ExternalAPIToken: together.String("external_api_token"),
					ExternalBaseURL:  together.String("external_base_url"),
					MaxTokens:        together.Int(1),
					NumWorkers:       together.Int(1),
					Temperature:      together.Float(0),
				},
				Labels:     []string{"string", "string"},
				PassLabels: []string{"string"},
				ModelToEvaluate: together.EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion{
					OfString: together.String("string"),
				},
			},
		},
		Type: together.EvalNewParamsTypeClassify,
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvalGet(t *testing.T) {
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
	_, err := client.Evals.Get(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvalListWithOptionalParams(t *testing.T) {
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
	_, err := client.Evals.List(context.TODO(), together.EvalListParams{
		Limit:  together.Int(0),
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

func TestEvalStatus(t *testing.T) {
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
	_, err := client.Evals.Status(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
