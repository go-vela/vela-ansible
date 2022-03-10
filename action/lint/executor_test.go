// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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
		err := Exec(test.lint)

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
