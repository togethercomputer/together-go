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
		Messages: []together.ChatCompletionNewParamsMessageUnion{{
			OfChatCompletionNewsMessageChatCompletionSystemMessageParam: &together.ChatCompletionNewParamsMessageChatCompletionSystemMessageParam{
				Content: "content",
				Role:    "system",
				Name:    together.String("name"),
			},
		}},
		Model:                         together.ChatCompletionNewParamsModelMetaLlamaMetaLlama3_1_8BInstructTurbo,
		ContextLengthExceededBehavior: together.ChatCompletionNewParamsContextLengthExceededBehaviorTruncate,
		Echo:                          together.Bool(true),
		FrequencyPenalty:              together.Float(0),
		FunctionCall: together.ChatCompletionNewParamsFunctionCallUnion{
			OfChatCompletionNewsFunctionCallString: together.String("none"),
		},
		LogitBias: map[string]float64{
			"105":  21.4,
			"1024": -10.5,
		},
		Logprobs:          together.Int(0),
		MaxTokens:         together.Int(0),
		MinP:              together.Float(0),
		N:                 together.Int(1),
		PresencePenalty:   together.Float(0),
		ReasoningEffort:   together.ChatCompletionNewParamsReasoningEffortMedium,
		RepetitionPenalty: together.Float(0),
		ResponseFormat: together.ChatCompletionNewParamsResponseFormatUnion{
			OfText: &together.ChatCompletionNewParamsResponseFormatText{},
		},
		SafetyModel: together.String("safety_model_name"),
		Seed:        together.Int(42),
		Stop:        []string{"string"},
		Temperature: together.Float(0),
		ToolChoice: together.ChatCompletionNewParamsToolChoiceUnion{
			OfString: together.String("tool_name"),
		},
		Tools: []together.ToolsParam{{
			Function: together.ToolsFunctionParam{
				Description: together.String("A description of the function."),
				Name:        together.String("function_name"),
				Parameters: map[string]any{
					"foo": "bar",
				},
			},
			Type: together.String("tool_type"),
		}},
		TopK: together.Int(0),
		TopP: together.Float(0),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
