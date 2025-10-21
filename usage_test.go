// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together_test

import (
	"context"
	"os"
	"testing"

	"github.com/togethercomputer/together-go"
	"github.com/togethercomputer/together-go/internal/testutil"
	"github.com/togethercomputer/together-go/option"
)

func TestUsage(t *testing.T) {
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
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), together.ChatCompletionNewParams{
		Messages: []together.ChatCompletionNewParamsMessageUnion{{
			OfChatCompletionNewsMessageChatCompletionUserMessageParam: &together.ChatCompletionNewParamsMessageChatCompletionUserMessageParam{
				Role: "user",
				Content: together.ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentUnion{
					OfString: together.String("Say this is a test!"),
				},
			},
		}},
		Model: together.ChatCompletionNewParamsModelQwenQwen2_5_72BInstructTurbo,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", chatCompletion.Choices)
}
