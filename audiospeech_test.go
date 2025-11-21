// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/togethercomputer/together-go"
	"github.com/togethercomputer/together-go/option"
)

func TestAudioSpeechNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := together.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Audio.Speech.New(context.TODO(), together.AudioSpeechNewParams{
		Input:            "input",
		Model:            together.AudioSpeechNewParamsModelCartesiaSonic,
		Voice:            "voice",
		Language:         together.AudioSpeechNewParamsLanguageEn,
		ResponseEncoding: together.AudioSpeechNewParamsResponseEncodingPcmF32le,
		ResponseFormat:   together.AudioSpeechNewParamsResponseFormatMP3,
		SampleRate:       together.Int(0),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
