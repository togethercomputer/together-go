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

func TestEndpointDeploymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Deployments.New(
		context.TODO(),
		"endpointId",
		together.EndpointDeploymentNewParams{
			ProjectID: together.String("projectId"),
			Autoscaling: together.DeploymentAutoscalingParam{
				MaxReplicas: together.Int(0),
				MinReplicas: together.Int(0),
				ScaleDown: together.ScalingRulesParam{
					Policies: []together.ScalingPolicyParam{{
						PeriodSeconds: 60,
						Type:          together.ScalingPolicyTypeScalingPolicyTypePercent,
						Value:         25,
					}, {
						PeriodSeconds: 60,
						Type:          together.ScalingPolicyTypeScalingPolicyTypePods,
						Value:         10,
					}},
					SelectPolicy: together.ScalingRulesSelectPolicyScalingPolicySelectMin,
				},
				ScaleDownWindow:   together.String("-160513s"),
				ScaleToZeroWindow: together.String("-160513s"),
				ScaleUp: together.ScalingRulesParam{
					Policies: []together.ScalingPolicyParam{{
						PeriodSeconds: 60,
						Type:          together.ScalingPolicyTypeScalingPolicyTypePercent,
						Value:         25,
					}, {
						PeriodSeconds: 60,
						Type:          together.ScalingPolicyTypeScalingPolicyTypePods,
						Value:         10,
					}},
					SelectPolicy: together.ScalingRulesSelectPolicyScalingPolicySelectMin,
				},
				ScaleUpWindow: together.String("-160513s"),
				ScalingMetrics: []together.ScalingMetricParam{{
					Name:       together.ScalingMetricNameActiveSessions,
					Target:     0,
					Type:       together.ScalingMetricTypeMetricTargetTypeValue,
					Percentile: together.String("percentile"),
				}},
			},
			Name:                            "name",
			ValidateOnly:                    together.Bool(true),
			Config:                          together.String("config"),
			ConfigID:                        together.String("configId"),
			InactiveTimeout:                 together.Int(0),
			MaxConcurrentRequestsPerReplica: together.String("maxConcurrentRequestsPerReplica"),
			Model:                           together.String("model"),
			ModelID:                         together.String("modelId"),
			ModelRevisionID:                 together.String("modelRevisionId"),
			Placement: together.EndpointDeploymentNewParamsPlacementUnion{
				OfInline: &together.EndpointDeploymentNewParamsPlacementInline{
					Inline: together.DeploymentPlacementConfigParam{
						CompliancePolicy: together.DeploymentPlacementConfigCompliancePolicyParam{
							Hipaa: together.Bool(true),
						},
						Constraint: together.DeploymentPlacementConfigConstraintEnforcementRequired,
						Regions:    []string{"string"},
					},
				},
			},
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

func TestEndpointDeploymentGet(t *testing.T) {
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
	_, err := client.Endpoints.Deployments.Get(
		context.TODO(),
		"id",
		together.EndpointDeploymentGetParams{
			ProjectID:  together.String("projectId"),
			EndpointID: "endpointId",
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

func TestEndpointDeploymentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Deployments.Update(
		context.TODO(),
		"id",
		together.EndpointDeploymentUpdateParams{
			ProjectID:  together.String("projectId"),
			EndpointID: "endpointId",
			UpdateMask: together.String("updateMask"),
			Autoscaling: together.DeploymentAutoscalingParam{
				MaxReplicas: together.Int(0),
				MinReplicas: together.Int(0),
				ScaleDown: together.ScalingRulesParam{
					Policies: []together.ScalingPolicyParam{{
						PeriodSeconds: 60,
						Type:          together.ScalingPolicyTypeScalingPolicyTypePercent,
						Value:         25,
					}, {
						PeriodSeconds: 60,
						Type:          together.ScalingPolicyTypeScalingPolicyTypePods,
						Value:         10,
					}},
					SelectPolicy: together.ScalingRulesSelectPolicyScalingPolicySelectMin,
				},
				ScaleDownWindow:   together.String("-160513s"),
				ScaleToZeroWindow: together.String("-160513s"),
				ScaleUp: together.ScalingRulesParam{
					Policies: []together.ScalingPolicyParam{{
						PeriodSeconds: 60,
						Type:          together.ScalingPolicyTypeScalingPolicyTypePercent,
						Value:         25,
					}, {
						PeriodSeconds: 60,
						Type:          together.ScalingPolicyTypeScalingPolicyTypePods,
						Value:         10,
					}},
					SelectPolicy: together.ScalingRulesSelectPolicyScalingPolicySelectMin,
				},
				ScaleUpWindow: together.String("-160513s"),
				ScalingMetrics: []together.ScalingMetricParam{{
					Name:       together.ScalingMetricNameActiveSessions,
					Target:     0,
					Type:       together.ScalingMetricTypeMetricTargetTypeValue,
					Percentile: together.String("percentile"),
				}},
			},
			Etag:                            together.String("etag"),
			InactiveTimeout:                 together.Int(0),
			MaxConcurrentRequestsPerReplica: together.String("maxConcurrentRequestsPerReplica"),
			Name:                            together.String("name"),
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

func TestEndpointDeploymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Deployments.List(
		context.TODO(),
		"endpointId",
		together.EndpointDeploymentListParams{
			ProjectID: together.String("projectId"),
			After:     together.String("after"),
			Filter:    together.String("filter"),
			Limit:     together.Int(0),
			OrderBy:   together.String("orderBy"),
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

func TestEndpointDeploymentDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Endpoints.Deployments.Delete(
		context.TODO(),
		"id",
		together.EndpointDeploymentDeleteParams{
			ProjectID:  together.String("projectId"),
			EndpointID: "endpointId",
			Etag:       together.String("etag"),
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
