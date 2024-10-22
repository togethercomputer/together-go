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
		Documents: together.F[together.RerankParamsDocumentsUnion](together.RerankParamsDocumentsArray([]map[string]interface{}{{
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
		}})),
		Model:           together.F("Salesforce/Llama-Rank-V1"),
		Query:           together.F("What animals can I find near Peru?"),
		RankFields:      together.F([]string{"title", "text"}),
		ReturnDocuments: together.F(true),
		TopN:            together.F(int64(2)),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
