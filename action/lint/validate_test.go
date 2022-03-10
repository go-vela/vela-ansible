// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package lint

import (
	"errors"
	"testing"
)

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		failure bool
		lint    *Linter
	}{
		{
			failure: false,
			lint: &Linter{
				LintPlaybook: "testdata/main.yml",
			},
		},
		{
			failure: false,
			lint: &Linter{
				LintVersion: true,
			},
		},
	}

	for _, test := range tests {
		if err := Validate(test.lint); err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}

func TestValidateError(t *testing.T) {
	tests := []struct {
		name    string
		lint    *Linter
		wantErr error
	}{
		{
			name: "Empty playbook",
			lint: &Linter{
				LintPlaybook: "testdata/empty.yml",
			},
			wantErr: ErrorEmptyLintPlaybook,
		},
		{
			name: "Playbook path doesn't exist",
			lint: &Linter{
				LintPlaybook: "notfound.yml",
			},
			wantErr: ErrorInvalidLintPlaybook,
		},
		{
			name: "Playbook empty string",
			lint: &Linter{
				LintPlaybook: "",
			},
			wantErr: ErrorMissingLintPlaybook,
		},
		{
			name:    "Playbook not provided",
			lint:    &Linter{},
			wantErr: ErrorMissingLintPlaybook,
		},
	}

	for _, test := range tests {
		err := Validate(test.lint)

		if err == nil {
			t.Errorf("should have returned err")
		}

		if !errors.Is(err, test.wantErr) {
			t.Errorf("Should have returned error: %v, instead got error: %v", test.wantErr, err)
		}
	}
}
