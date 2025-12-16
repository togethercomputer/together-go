// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/togethercomputer/together-go"
	"github.com/togethercomputer/together-go/internal/testutil"
	"github.com/togethercomputer/together-go/option"
)

func TestAudioTranscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Audio.Transcriptions.New(context.TODO(), together.AudioTranscriptionNewParams{
		File: together.AudioTranscriptionNewParamsFileUnion{
			OfFile: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		},
		Diarize:        together.Bool(true),
		Language:       together.String("en"),
		MaxSpeakers:    together.Int(0),
		MinSpeakers:    together.Int(0),
		Model:          together.AudioTranscriptionNewParamsModelOpenAIWhisperLargeV3,
		Prompt:         together.String("prompt"),
		ResponseFormat: together.AudioTranscriptionNewParamsResponseFormatJson,
		Temperature:    together.Float(0),
		TimestampGranularities: together.AudioTranscriptionNewParamsTimestampGranularitiesUnion{
			OfAudioTranscriptionNewsTimestampGranularitiesArrayItemArray: []string{"word", "segment"},
		},
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
