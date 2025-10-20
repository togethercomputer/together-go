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

func TestVideoNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Videos.New(context.TODO(), together.VideoNewParams{
		Model: together.F("model"),
		Fps:   together.F(int64(0)),
		FrameImages: together.F([]together.VideoNewParamsFrameImage{{
			Frame:      together.F[together.VideoNewParamsFrameImagesFrameUnion](shared.UnionFloat(0.000000)),
			InputImage: together.F("input_image"),
		}}),
		GuidanceScale:   together.F(int64(0)),
		Height:          together.F(int64(0)),
		NegativePrompt:  together.F("negative_prompt"),
		OutputFormat:    together.F(together.VideoNewParamsOutputFormatMP4),
		OutputQuality:   together.F(int64(0)),
		Prompt:          together.F("x"),
		ReferenceImages: together.F([]interface{}{map[string]interface{}{}}),
		Seconds:         together.F("seconds"),
		Seed:            together.F(int64(0)),
		Steps:           together.F(int64(10)),
		Width:           together.F(int64(0)),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVideoGet(t *testing.T) {
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
	_, err := client.Videos.Get(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
