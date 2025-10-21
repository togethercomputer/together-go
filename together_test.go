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

func TestTogetherRerankWithOptionalParams(t *testing.T) {
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
	_, err := client.Rerank(context.TODO(), together.RerankParams{
		Documents: together.RerankParamsDocumentsUnion{
			OfMapOfAnyMap: []map[string]any{{
				"title": "bar",
				"text":  "bar",
			}, {
				"title": "bar",
				"text":  "bar",
			}, {
				"title": "bar",
				"text":  "bar",
			}, {
				"title": "bar",
				"text":  "bar",
			}},
		},
		Model:           together.RerankParamsModelSalesforceLlamaRankV1,
		Query:           "What animals can I find near Peru?",
		RankFields:      []string{"title", "text"},
		ReturnDocuments: together.Bool(true),
		TopN:            together.Int(2),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
