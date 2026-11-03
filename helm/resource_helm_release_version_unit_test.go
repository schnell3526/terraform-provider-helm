// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package helm

import "testing"

// Regression test for hashicorp/terraform-provider-helm#1694.
func TestVersionsEqual(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"1.2.3", "1.2.3", true},
		{"v1.2.3", "1.2.3", true},
		{"1.2.3", "v1.2.3", true},
		{"v0.3.0", "0.3.0+abc123", true},
		{"0.3.0+abc", "0.3.0+def", true},
		{"1.2.3", "1.2.4", false},
		{"", "2025.6.1", false},
		{"1.2.3-rc1", "1.2.3", false},
	}
	for _, c := range cases {
		if got := versionsEqual(c.a, c.b); got != c.want {
			t.Errorf("versionsEqual(%q, %q) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
