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

func TestAudioTranslationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Audio.Translations.New(context.TODO(), together.AudioTranslationNewParams{
		File:                   together.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		Language:               together.F("en"),
		Model:                  together.F(together.AudioTranslationNewParamsModelOpenAIWhisperLargeV3),
		Prompt:                 together.F("prompt"),
		ResponseFormat:         together.F(together.AudioTranslationNewParamsResponseFormatJson),
		Temperature:            together.F(0.000000),
		TimestampGranularities: together.F[together.AudioTranslationNewParamsTimestampGranularitiesUnion](together.AudioTranslationNewParamsTimestampGranularitiesArray([]together.AudioTranslationNewParamsTimestampGranularitiesArrayItem{together.AudioTranslationNewParamsTimestampGranularitiesArrayItemWord, together.AudioTranslationNewParamsTimestampGranularitiesArrayItemSegment})),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
