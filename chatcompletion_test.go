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

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Chat.Completions.New(context.TODO(), together.ChatCompletionNewParams{
		Messages: together.F([]together.ChatCompletionNewParamsMessage{{
			Content: together.F("content"),
			Role:    together.F(together.ChatCompletionNewParamsMessagesRoleSystem),
		}}),
		Model:            together.F(together.ChatCompletionNewParamsModelQwenQwen2_5_72BInstructTurbo),
		Echo:             together.F(true),
		FrequencyPenalty: together.F(0.000000),
		FunctionCall:     together.F[together.ChatCompletionNewParamsFunctionCallUnion](together.ChatCompletionNewParamsFunctionCallString(together.ChatCompletionNewParamsFunctionCallStringNone)),
		LogitBias: together.F(map[string]float64{
			"105":  21.400000,
			"1024": -10.500000,
		}),
		Logprobs:          together.F(int64(0)),
		MaxTokens:         together.F(int64(0)),
		MinP:              together.F(0.000000),
		N:                 together.F(int64(1)),
		PresencePenalty:   together.F(0.000000),
		RepetitionPenalty: together.F(0.000000),
		ResponseFormat: together.F(together.ChatCompletionNewParamsResponseFormat{
			Schema: together.F(map[string]string{
				"foo": "string",
			}),
			Type: together.F("json"),
		}),
		SafetyModel: together.F("safety_model_name"),
		Seed:        together.F(int64(42)),
		Stop:        together.F([]string{"string"}),
		Temperature: together.F(0.000000),
		ToolChoice:  together.F[together.ChatCompletionNewParamsToolChoiceUnion](shared.UnionString("tool_name")),
		Tools: together.F([]together.ToolsParam{{
			Function: together.F(together.ToolsFunctionParam{
				Description: together.F("A description of the function."),
				Name:        together.F("function_name"),
				Parameters: together.F(map[string]interface{}{
					"foo": "bar",
				}),
			}),
			Type: together.F("tool_type"),
		}}),
		TopK: together.F(int64(0)),
		TopP: together.F(0.000000),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
