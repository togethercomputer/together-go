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

func TestImageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Images.New(context.TODO(), together.ImageNewParams{
		Model:    together.F(together.ImageNewParamsModelBlackForestLabsFlux1SchnellFree),
		Prompt:   together.F("cat floating in space, cinematic"),
		Guidance: together.F(0.000000),
		Height:   together.F(int64(0)),
		ImageLoras: together.F([]together.ImageNewParamsImageLora{{
			Path:  together.F("path"),
			Scale: together.F(0.000000),
		}}),
		ImageURL:       together.F("image_url"),
		N:              together.F(int64(0)),
		NegativePrompt: together.F("negative_prompt"),
		OutputFormat:   together.F(together.ImageNewParamsOutputFormatJpeg),
		ResponseFormat: together.F(together.ImageNewParamsResponseFormatBase64),
		Seed:           together.F(int64(0)),
		Steps:          together.F(int64(0)),
		Width:          together.F(int64(0)),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
