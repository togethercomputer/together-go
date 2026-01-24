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

func TestBetaJigGet(t *testing.T) {
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
	_, err := client.Beta.Jig.Get(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaJigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Jig.Update(
		context.TODO(),
		"id",
		together.BetaJigUpdateParams{
			Args: []string{"string"},
			Autoscaling: map[string]string{
				"foo": "string",
			},
			Command:     []string{"string"},
			CPU:         together.Float(0.1),
			Description: together.String("description"),
			EnvironmentVariables: []together.BetaJigUpdateParamsEnvironmentVariable{{
				Name:            "name",
				Value:           together.String("value"),
				ValueFromSecret: together.String("value_from_secret"),
			}},
			GPUCount:                      together.Int(0),
			GPUType:                       together.BetaJigUpdateParamsGPUTypeH100_80gb,
			HealthCheckPath:               together.String("health_check_path"),
			Image:                         together.String("image"),
			MaxReplicas:                   together.Int(0),
			Memory:                        together.Float(0.1),
			MinReplicas:                   together.Int(0),
			Name:                          together.String("x"),
			Port:                          together.Int(0),
			Storage:                       together.Int(0),
			TerminationGracePeriodSeconds: together.Int(0),
			Volumes: []together.BetaJigUpdateParamsVolume{{
				MountPath: "mount_path",
				Name:      "name",
			}},
		},
	)
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaJigDeployWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Jig.Deploy(context.TODO(), together.BetaJigDeployParams{
		GPUType: together.BetaJigDeployParamsGPUTypeH100_80gb,
		Image:   "image",
		Name:    "x",
		Args:    []string{"string"},
		Autoscaling: map[string]string{
			"foo": "string",
		},
		Command:     []string{"string"},
		CPU:         together.Float(0.1),
		Description: together.String("description"),
		EnvironmentVariables: []together.BetaJigDeployParamsEnvironmentVariable{{
			Name:            "name",
			Value:           together.String("value"),
			ValueFromSecret: together.String("value_from_secret"),
		}},
		GPUCount:                      together.Int(0),
		HealthCheckPath:               together.String("health_check_path"),
		MaxReplicas:                   together.Int(0),
		Memory:                        together.Float(0.1),
		MinReplicas:                   together.Int(0),
		Port:                          together.Int(0),
		Storage:                       together.Int(0),
		TerminationGracePeriodSeconds: together.Int(0),
		Volumes: []together.BetaJigDeployParamsVolume{{
			MountPath: "mount_path",
			Name:      "name",
		}},
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaJigDestroy(t *testing.T) {
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
	_, err := client.Beta.Jig.Destroy(context.TODO(), "id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
