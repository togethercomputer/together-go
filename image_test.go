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

func TestImageGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Images.Generate(context.TODO(), together.ImageGenerateParams{
		Model:                together.ImageGenerateParamsModelBlackForestLabsFlux1SchnellFree,
		Prompt:               "cat floating in space, cinematic",
		DisableSafetyChecker: together.Bool(true),
		GuidanceScale:        together.Float(0),
		Height:               together.Int(0),
		ImageLoras: []together.ImageGenerateParamsImageLora{{
			Path:  "path",
			Scale: 0,
		}},
		ImageURL:       together.String("image_url"),
		N:              together.Int(0),
		NegativePrompt: together.String("negative_prompt"),
		OutputFormat:   together.ImageGenerateParamsOutputFormatJpeg,
		ResponseFormat: together.ImageGenerateParamsResponseFormatBase64,
		Seed:           together.Int(0),
		Steps:          together.Int(0),
		Width:          together.Int(0),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
