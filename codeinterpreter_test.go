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

func TestCodeInterpreterExecuteWithOptionalParams(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
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
	_, err := client.CodeInterpreter.Execute(context.TODO(), together.CodeInterpreterExecuteParams{
		Code:     together.F("print('Hello, world!')"),
		Language: together.F(together.CodeInterpreterExecuteParamsLanguagePython),
		Files: together.F([]together.CodeInterpreterExecuteParamsFile{{
			Content:  together.F("content"),
			Encoding: together.F(together.CodeInterpreterExecuteParamsFilesEncodingString),
			Name:     together.F("name"),
		}}),
		SessionID: together.F("ses_abcDEF123"),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
