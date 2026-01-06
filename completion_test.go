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

func TestCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Completions.New(context.TODO(), together.CompletionNewParams{
		Model:            together.CompletionNewParamsModel("mistralai/Mixtral-8x7B-Instruct-v0.1"),
		Prompt:           "<s>[INST] What is the capital of France? [/INST]",
		Echo:             together.Bool(true),
		FrequencyPenalty: together.Float(0),
		LogitBias: map[string]float64{
			"105":  21.4,
			"1024": -10.5,
		},
		Logprobs:          together.Int(0),
		MaxTokens:         together.Int(0),
		MinP:              together.Float(0),
		N:                 together.Int(1),
		PresencePenalty:   together.Float(0),
		RepetitionPenalty: together.Float(0),
		SafetyModel:       together.CompletionNewParamsSafetyModel("safety_model_name"),
		Seed:              together.Int(42),
		Stop:              []string{"string"},
		Temperature:       together.Float(0),
		TopK:              together.Int(0),
		TopP:              together.Float(0),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
