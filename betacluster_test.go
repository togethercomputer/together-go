// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/togethercomputer/together-go"
	"github.com/togethercomputer/together-go/internal/testutil"
	"github.com/togethercomputer/together-go/option"
)

func TestBetaClusterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Clusters.New(context.TODO(), together.BetaClusterNewParams{
		BillingType:         together.BetaClusterNewParamsBillingTypeReserved,
		ClusterName:         "cluster_name",
		CudaVersion:         "cuda_version",
		DurationDays:        0,
		GPUType:             together.BetaClusterNewParamsGPUTypeH100Sxm,
		NumGPUs:             0,
		NvidiaDriverVersion: "nvidia_driver_version",
		Region:              "region",
		AcceptanceTestsParams: together.BetaClusterNewParamsAcceptanceTestsParams{
			DcgmDiagLevel:         "DCGM_DIAG_LEVEL_SHORT",
			DcgmDiagSkipped:       together.Bool(true),
			Enabled:               together.Bool(true),
			GPUBurnDuration:       together.Int(0),
			GPUBurnSkipped:        together.Bool(true),
			NcclMultiNodeSkipped:  together.Bool(true),
			NcclSingleNodeSkipped: together.Bool(true),
		},
		AddOns: []together.BetaClusterNewParamsAddOn{{
			AddOnType: "add_on_type",
			Name:      "name",
			Config: together.BetaClusterNewParamsAddOnConfig{
				Dashboard: together.BetaClusterNewParamsAddOnConfigDashboard{
					Enabled: together.Bool(true),
				},
				Ingress: together.BetaClusterNewParamsAddOnConfigIngress{
					Enabled: together.Bool(true),
				},
			},
		}},
		AutoScale:        together.Bool(true),
		AutoScaleMaxGPUs: together.Int(0),
		AutoScaled:       together.Bool(true),
		CapacityPoolID:   together.String("capacity_pool_id"),
		ClusterConfig: together.BetaClusterNewParamsClusterConfig{
			LoadBalancer:       "NONE",
			GPUOperatorVersion: together.String("gpu_operator_version"),
			Ingress: together.BetaClusterNewParamsClusterConfigIngress{
				Enabled: together.Bool(true),
			},
			JumphostEnabled:            together.Bool(true),
			KubernetesDashboardEnabled: together.Bool(true),
			Observability: together.BetaClusterNewParamsClusterConfigObservability{
				Enabled: together.Bool(true),
			},
			SlurmStartupScripts: together.BetaClusterNewParamsClusterConfigSlurmStartupScripts{
				ControllerEpilog:  together.String("controller_epilog"),
				ControllerProlog:  together.String("controller_prolog"),
				ExtraSlurmConf:    together.String("extra_slurm_conf"),
				LoginInitScript:   together.String("login_init_script"),
				NodesetInitScript: together.String("nodeset_init_script"),
				WorkerEpilog:      together.String("worker_epilog"),
				WorkerProlog:      together.String("worker_prolog"),
			},
		},
		ClusterType:            together.BetaClusterNewParamsClusterTypeKubernetes,
		GPUNodeFailoverEnabled: together.Bool(true),
		InstallTraefik:         together.Bool(true),
		NumCapacityPoolGPUs:    together.Int(0),
		NumPreemptibleGPUs:     together.Int(0),
		NumReservedGPUs:        together.Int(0),
		OidcConfig: together.BetaClusterNewParamsOidcConfig{
			ClientID:       "client_id",
			GroupClaim:     "group_claim",
			GroupPrefix:    "group_prefix",
			IssuerURL:      "issuer_url",
			UsernameClaim:  "username_claim",
			UsernamePrefix: "username_prefix",
			CaCert:         together.String("ca_cert"),
		},
		ProjectID:            together.String("project_id"),
		ReservationEndTime:   together.Time(time.Now()),
		ReservationStartTime: together.Time(time.Now()),
		SharedVolume: together.BetaClusterNewParamsSharedVolume{
			Region:                 "region",
			SizeTib:                0,
			VolumeName:             "volume_name",
			IsLifecycleIndependent: together.Bool(true),
		},
		SlurmImage:      together.String("slurm_image"),
		SlurmShmSizeGib: together.Int(0),
		VolumeID:        together.String("volume_id"),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaClusterGet(t *testing.T) {
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
	_, err := client.Beta.Clusters.Get(context.TODO(), "cluster_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaClusterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Clusters.Update(
		context.TODO(),
		"cluster_id",
		together.BetaClusterUpdateParams{
			AddOns: []together.BetaClusterUpdateParamsAddOn{{
				Name: "name",
				Config: together.BetaClusterUpdateParamsAddOnConfig{
					Dashboard: together.BetaClusterUpdateParamsAddOnConfigDashboard{
						Enabled: together.Bool(true),
					},
					Ingress: together.BetaClusterUpdateParamsAddOnConfigIngress{
						Enabled: together.Bool(true),
					},
				},
			}},
			ClusterConfig: together.BetaClusterUpdateParamsClusterConfig{
				LoadBalancer:       "NONE",
				GPUOperatorVersion: together.String("gpu_operator_version"),
				Ingress: together.BetaClusterUpdateParamsClusterConfigIngress{
					Enabled: together.Bool(true),
				},
				JumphostEnabled:            together.Bool(true),
				KubernetesDashboardEnabled: together.Bool(true),
				Observability: together.BetaClusterUpdateParamsClusterConfigObservability{
					Enabled: together.Bool(true),
				},
				SlurmStartupScripts: together.BetaClusterUpdateParamsClusterConfigSlurmStartupScripts{
					ControllerEpilog:  together.String("controller_epilog"),
					ControllerProlog:  together.String("controller_prolog"),
					ExtraSlurmConf:    together.String("extra_slurm_conf"),
					LoginInitScript:   together.String("login_init_script"),
					NodesetInitScript: together.String("nodeset_init_script"),
					WorkerEpilog:      together.String("worker_epilog"),
					WorkerProlog:      together.String("worker_prolog"),
				},
			},
			ClusterType:        together.BetaClusterUpdateParamsClusterTypeKubernetes,
			NumGPUs:            together.Int(0),
			NumPreemptibleGPUs: together.Int(0),
			NumReservedGPUs:    together.Int(0),
			ReservationEndTime: together.Time(time.Now()),
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

func TestBetaClusterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Clusters.List(context.TODO(), together.BetaClusterListParams{
		ProjectID: together.String("project_id"),
	})
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaClusterDelete(t *testing.T) {
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
	_, err := client.Beta.Clusters.Delete(context.TODO(), "cluster_id")
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaClusterListRegions(t *testing.T) {
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
	_, err := client.Beta.Clusters.ListRegions(context.TODO())
	if err != nil {
		var apierr *together.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
