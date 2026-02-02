// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package helm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// TestStateUpgraderV0_OldType_WithResources tests that the version 0 oldType
// can unmarshal state that contains a resources attribute.
// This is a regression test for https://github.com/schnell3526/terraform-provider-helm/issues/18
func TestStateUpgraderV0_OldType_WithResources(t *testing.T) {
	// oldType definition from version 0 StateUpgrader
	oldType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"metadata": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name":        tftypes.String,
						"namespace":   tftypes.String,
						"revision":    tftypes.Number,
						"version":     tftypes.String,
						"chart":       tftypes.String,
						"app_version": tftypes.String,
						"values":      tftypes.String,
					},
				},
			},
			"postrender": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"binary_path": tftypes.String,
						"args":        tftypes.List{ElementType: tftypes.String},
					},
				},
			},
			"set": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name":  tftypes.String,
						"value": tftypes.String,
						"type":  tftypes.String,
					},
				},
			},
			"set_sensitive": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name":  tftypes.String,
						"value": tftypes.String,
						"type":  tftypes.String,
					},
				},
			},
			"atomic":                     tftypes.Bool,
			"chart":                      tftypes.String,
			"cleanup_on_fail":            tftypes.Bool,
			"create_namespace":           tftypes.Bool,
			"dependency_update":          tftypes.Bool,
			"description":                tftypes.String,
			"devel":                      tftypes.Bool,
			"disable_crd_hooks":          tftypes.Bool,
			"disable_openapi_validation": tftypes.Bool,
			"disable_webhooks":           tftypes.Bool,
			"force_update":               tftypes.Bool,
			"id":                         tftypes.String,
			"keyring":                    tftypes.String,
			"lint":                       tftypes.Bool,
			"manifest":                   tftypes.String,
			"max_history":                tftypes.Number,
			"name":                       tftypes.String,
			"namespace":                  tftypes.String,
			"recreate_pods":              tftypes.Bool,
			"render_subchart_notes":      tftypes.Bool,
			"replace":                    tftypes.Bool,
			"repository":                 tftypes.String,
			"repository_ca_file":         tftypes.String,
			"repository_cert_file":       tftypes.String,
			"repository_key_file":        tftypes.String,
			"repository_password":        tftypes.String,
			"repository_username":        tftypes.String,
			"reset_values":               tftypes.Bool,
			"resources":                  tftypes.Map{ElementType: tftypes.String},
			"reuse_values":               tftypes.Bool,
			"skip_crds":                  tftypes.Bool,
			"status":                     tftypes.String,
			"timeout":                    tftypes.Number,
			"values":                     tftypes.List{ElementType: tftypes.String},
			"verify":                     tftypes.Bool,
			"version":                    tftypes.String,
			"wait":                       tftypes.Bool,
			"wait_for_jobs":              tftypes.Bool,
		},
	}

	// Test case 1: State WITH resources attribute (from schnell3526/helm provider)
	t.Run("state with resources attribute", func(t *testing.T) {
		stateJSON := `{
			"metadata": [{
				"name": "test",
				"namespace": "default",
				"revision": 1,
				"version": "1.0.0",
				"chart": "nginx",
				"app_version": "1.0.0",
				"values": "{}"
			}],
			"postrender": [],
			"set": [],
			"set_sensitive": [],
			"atomic": false,
			"chart": "nginx",
			"cleanup_on_fail": false,
			"create_namespace": false,
			"dependency_update": false,
			"description": "",
			"devel": false,
			"disable_crd_hooks": false,
			"disable_openapi_validation": false,
			"disable_webhooks": false,
			"force_update": false,
			"id": "test",
			"keyring": "",
			"lint": false,
			"manifest": "",
			"max_history": 0,
			"name": "test",
			"namespace": "default",
			"recreate_pods": false,
			"render_subchart_notes": true,
			"replace": false,
			"repository": "",
			"repository_ca_file": "",
			"repository_cert_file": "",
			"repository_key_file": "",
			"repository_password": "",
			"repository_username": "",
			"reset_values": false,
			"resources": {"Deployment/test": "present"},
			"reuse_values": false,
			"skip_crds": false,
			"status": "deployed",
			"timeout": 300,
			"values": [],
			"verify": false,
			"version": "1.0.0",
			"wait": true,
			"wait_for_jobs": false
		}`

		_, err := tftypes.ValueFromJSON([]byte(stateJSON), oldType)
		if err != nil {
			t.Fatalf("Failed to unmarshal state with resources attribute: %v", err)
		}
	})

	// Test case 2: State WITHOUT resources attribute (from hashicorp/helm provider)
	t.Run("state without resources attribute", func(t *testing.T) {
		stateJSON := `{
			"metadata": [{
				"name": "test",
				"namespace": "default",
				"revision": 1,
				"version": "1.0.0",
				"chart": "nginx",
				"app_version": "1.0.0",
				"values": "{}"
			}],
			"postrender": [],
			"set": [],
			"set_sensitive": [],
			"atomic": false,
			"chart": "nginx",
			"cleanup_on_fail": false,
			"create_namespace": false,
			"dependency_update": false,
			"description": "",
			"devel": false,
			"disable_crd_hooks": false,
			"disable_openapi_validation": false,
			"disable_webhooks": false,
			"force_update": false,
			"id": "test",
			"keyring": "",
			"lint": false,
			"manifest": "",
			"max_history": 0,
			"name": "test",
			"namespace": "default",
			"recreate_pods": false,
			"render_subchart_notes": true,
			"replace": false,
			"repository": "",
			"repository_ca_file": "",
			"repository_cert_file": "",
			"repository_key_file": "",
			"repository_password": "",
			"repository_username": "",
			"reset_values": false,
			"reuse_values": false,
			"skip_crds": false,
			"status": "deployed",
			"timeout": 300,
			"values": [],
			"verify": false,
			"version": "1.0.0",
			"wait": true,
			"wait_for_jobs": false
		}`

		_, err := tftypes.ValueFromJSON([]byte(stateJSON), oldType)
		if err != nil {
			t.Fatalf("Failed to unmarshal state without resources attribute: %v", err)
		}
	})
}

// TestStateUpgraderV1_OldType_WithResources tests that the version 1 oldType
// can unmarshal state that contains a resources attribute.
// This is a regression test for https://github.com/schnell3526/terraform-provider-helm/issues/18
func TestStateUpgraderV1_OldType_WithResources(t *testing.T) {
	// oldType definition from version 1 StateUpgrader
	oldType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"metadata": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name":           tftypes.String,
						"namespace":      tftypes.String,
						"revision":       tftypes.Number,
						"version":        tftypes.String,
						"chart":          tftypes.String,
						"app_version":    tftypes.String,
						"values":         tftypes.String,
						"first_deployed": tftypes.Number,
						"last_deployed":  tftypes.Number,
						"notes":          tftypes.String,
					},
				},
			},
			"postrender": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"binary_path": tftypes.String,
						"args":        tftypes.List{ElementType: tftypes.String},
					},
				},
			},
			"set": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name":  tftypes.String,
						"value": tftypes.String,
						"type":  tftypes.String,
					},
				},
			},
			"set_list": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name": tftypes.String,
						"value": tftypes.List{
							ElementType: tftypes.String,
						},
					},
				},
			},
			"set_sensitive": tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name":  tftypes.String,
						"value": tftypes.String,
						"type":  tftypes.String,
					},
				},
			},
			"atomic":                     tftypes.Bool,
			"chart":                      tftypes.String,
			"cleanup_on_fail":            tftypes.Bool,
			"create_namespace":           tftypes.Bool,
			"dependency_update":          tftypes.Bool,
			"description":                tftypes.String,
			"devel":                      tftypes.Bool,
			"disable_crd_hooks":          tftypes.Bool,
			"disable_openapi_validation": tftypes.Bool,
			"disable_webhooks":           tftypes.Bool,
			"force_update":               tftypes.Bool,
			"id":                         tftypes.String,
			"keyring":                    tftypes.String,
			"lint":                       tftypes.Bool,
			"manifest":                   tftypes.String,
			"max_history":                tftypes.Number,
			"name":                       tftypes.String,
			"namespace":                  tftypes.String,
			"pass_credentials":           tftypes.Bool,
			"recreate_pods":              tftypes.Bool,
			"render_subchart_notes":      tftypes.Bool,
			"replace":                    tftypes.Bool,
			"repository":                 tftypes.String,
			"repository_ca_file":         tftypes.String,
			"repository_cert_file":       tftypes.String,
			"repository_key_file":        tftypes.String,
			"repository_password":        tftypes.String,
			"repository_username":        tftypes.String,
			"reset_values":               tftypes.Bool,
			"resources":                  tftypes.Map{ElementType: tftypes.String},
			"reuse_values":               tftypes.Bool,
			"skip_crds":                  tftypes.Bool,
			"status":                     tftypes.String,
			"timeout":                    tftypes.Number,
			"upgrade_install":            tftypes.Bool,
			"values":                     tftypes.List{ElementType: tftypes.String},
			"verify":                     tftypes.Bool,
			"version":                    tftypes.String,
			"wait":                       tftypes.Bool,
			"wait_for_jobs":              tftypes.Bool,
		},
	}

	// Test case 1: State WITH resources attribute (from schnell3526/helm provider)
	t.Run("state with resources attribute", func(t *testing.T) {
		stateJSON := `{
			"metadata": [{
				"name": "test",
				"namespace": "default",
				"revision": 1,
				"version": "1.0.0",
				"chart": "nginx",
				"app_version": "1.0.0",
				"values": "{}",
				"first_deployed": 1234567890,
				"last_deployed": 1234567890,
				"notes": "test notes"
			}],
			"postrender": [],
			"set": [],
			"set_list": [],
			"set_sensitive": [],
			"atomic": false,
			"chart": "nginx",
			"cleanup_on_fail": false,
			"create_namespace": false,
			"dependency_update": false,
			"description": "",
			"devel": false,
			"disable_crd_hooks": false,
			"disable_openapi_validation": false,
			"disable_webhooks": false,
			"force_update": false,
			"id": "test",
			"keyring": "",
			"lint": false,
			"manifest": "",
			"max_history": 0,
			"name": "test",
			"namespace": "default",
			"pass_credentials": false,
			"recreate_pods": false,
			"render_subchart_notes": true,
			"replace": false,
			"repository": "",
			"repository_ca_file": "",
			"repository_cert_file": "",
			"repository_key_file": "",
			"repository_password": "",
			"repository_username": "",
			"reset_values": false,
			"resources": {"Deployment/test": "present"},
			"reuse_values": false,
			"skip_crds": false,
			"status": "deployed",
			"timeout": 300,
			"upgrade_install": false,
			"values": [],
			"verify": false,
			"version": "1.0.0",
			"wait": true,
			"wait_for_jobs": false
		}`

		_, err := tftypes.ValueFromJSON([]byte(stateJSON), oldType)
		if err != nil {
			t.Fatalf("Failed to unmarshal state with resources attribute: %v", err)
		}
	})

	// Test case 2: State WITHOUT resources attribute (from hashicorp/helm v2.17.0)
	t.Run("state without resources attribute", func(t *testing.T) {
		stateJSON := `{
			"metadata": [{
				"name": "test",
				"namespace": "default",
				"revision": 1,
				"version": "1.0.0",
				"chart": "nginx",
				"app_version": "1.0.0",
				"values": "{}",
				"first_deployed": 1234567890,
				"last_deployed": 1234567890,
				"notes": "test notes"
			}],
			"postrender": [],
			"set": [],
			"set_list": [],
			"set_sensitive": [],
			"atomic": false,
			"chart": "nginx",
			"cleanup_on_fail": false,
			"create_namespace": false,
			"dependency_update": false,
			"description": "",
			"devel": false,
			"disable_crd_hooks": false,
			"disable_openapi_validation": false,
			"disable_webhooks": false,
			"force_update": false,
			"id": "test",
			"keyring": "",
			"lint": false,
			"manifest": "",
			"max_history": 0,
			"name": "test",
			"namespace": "default",
			"pass_credentials": false,
			"recreate_pods": false,
			"render_subchart_notes": true,
			"replace": false,
			"repository": "",
			"repository_ca_file": "",
			"repository_cert_file": "",
			"repository_key_file": "",
			"repository_password": "",
			"repository_username": "",
			"reset_values": false,
			"reuse_values": false,
			"skip_crds": false,
			"status": "deployed",
			"timeout": 300,
			"upgrade_install": false,
			"values": [],
			"verify": false,
			"version": "1.0.0",
			"wait": true,
			"wait_for_jobs": false
		}`

		_, err := tftypes.ValueFromJSON([]byte(stateJSON), oldType)
		if err != nil {
			t.Fatalf("Failed to unmarshal state without resources attribute: %v", err)
		}
	})
}

// TestStateUpgraderOldType_WithoutResourcesFails demonstrates the issue fixed
// by adding resources to oldType. Without the resources attribute in oldType,
// unmarshaling a state that contains resources would fail.
func TestStateUpgraderOldType_WithoutResourcesFails(t *testing.T) {
	// oldType WITHOUT resources attribute (this would cause the error)
	oldTypeWithoutResources := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"name":      tftypes.String,
			"namespace": tftypes.String,
		},
	}

	// State WITH resources attribute
	stateJSON := `{
		"name": "test",
		"namespace": "default",
		"resources": {"Deployment/test": "present"}
	}`

	_, err := tftypes.ValueFromJSON([]byte(stateJSON), oldTypeWithoutResources)
	if err == nil {
		t.Fatal("Expected error when unmarshaling state with resources attribute using oldType without resources, but got none")
	}

	// Verify the error message contains "unsupported attribute"
	expectedError := "unsupported attribute"
	if err.Error() == "" || !contains(err.Error(), expectedError) {
		t.Fatalf("Expected error containing %q, but got: %v", expectedError, err)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
