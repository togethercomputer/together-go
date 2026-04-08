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
		Model: "model",
		Fps:   together.Int(0),
		FrameImages: []together.VideoNewParamsFrameImage{{
			InputImage: "input_image",
			Frame: together.VideoNewParamsFrameImageFrameUnion{
				OfFloat: together.Float(0),
			},
		}},
		GenerateAudio: together.Bool(true),
		GuidanceScale: together.Int(0),
		Height:        together.Int(0),
		Media: together.VideoNewParamsMedia{
			AudioInputs: []together.VideoNewParamsMediaAudioInputUnion{{
				OfString: together.String("string"),
			}},
			FrameImages: []together.VideoNewParamsMediaFrameImage{{
				InputImage: "input_image",
				Frame: together.VideoNewParamsMediaFrameImageFrameUnion{
					OfFloat: together.Float(0),
				},
			}},
			FrameVideos: []together.VideoNewParamsMediaFrameVideo{{
				Video: "video",
			}},
			ReferenceImages: []string{"string"},
			ReferenceVideos: []together.VideoNewParamsMediaReferenceVideo{{
				Video: "video",
			}},
			SourceVideo: together.VideoNewParamsMediaSourceVideoUnion{
				OfString: together.String("string"),
			},
		},
		NegativePrompt:  together.String("negative_prompt"),
		OutputFormat:    together.VideoNewParamsOutputFormatMP4,
		OutputQuality:   together.Int(0),
		Prompt:          together.String("x"),
		Ratio:           together.String("ratio"),
		ReferenceImages: []string{"string"},
		Resolution:      together.String("resolution"),
		Seconds:         together.String("seconds"),
		Seed:            together.Int(0),
		Steps:           together.Int(10),
		Width:           together.Int(0),
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
