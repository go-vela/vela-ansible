// SPDX-License-Identifier: Apache-2.0

package lint

import (
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		failure bool
		lint    *Linter
	}{
		{
			failure: true,
			lint: &Linter{
				LintPlaybook: "",
			},
		},
	}

	for _, test := range tests {
		err := Exec(t.Context(), test.lint)

		if test.failure {
			if err == nil {
				t.Errorf("Flags should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Flags returned err: %v", err)
		}
	}
}
