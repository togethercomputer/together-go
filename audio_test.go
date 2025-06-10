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

func TestAudioNewWithOptionalParams(t *testing.T) {
	t.Skip("AttributeError: BinaryAPIResponse object has no attribute response")
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
	resp, err := client.Audio.New(context.TODO(), together.AudioNewParams{
		Input:            together.F("input"),
		Model:            together.F(together.AudioNewParamsModelCartesiaSonic),
		Voice:            together.F(together.AudioNewParamsVoiceLaidbackWoman),
		Language:         together.F(together.AudioNewParamsLanguageEn),
		ResponseEncoding: together.F(together.AudioNewParamsResponseEncodingPcmF32le),
		ResponseFormat:   together.F(together.AudioNewParamsResponseFormatMP3),
		SampleRate:       together.F(0.000000),
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
