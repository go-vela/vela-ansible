// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.
package main

import (
	"testing"

	"github.com/go-vela/vela-ansible/action/lint"
	"github.com/go-vela/vela-ansible/action/playbook"
)

func TestPluginValidateSuccess(t *testing.T) {
	tests := []struct {
		plugin *Plugin
	}{
		{
			plugin: &Plugin{
				action: AnsiblePlaybook,
				Playbook: &playbook.Playbook{
					Playbook: "testdata/main.yml",
					Options: &playbook.Options{
						Inventory: []string{"testdata/inventory/hosts.yml"},
					},
				},
			},
		},
		{
			plugin: &Plugin{
				action: AnsibleLint,
				Lint: &lint.Linter{
					LintPlaybook: "testdata/main.yml",
					LintSkip:     []string{"403"},
				},
			},
		},
	}

	for _, test := range tests {
		if err := test.plugin.Validate(); err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}

func TestPluginValidateError(t *testing.T) {
	tests := []struct {
		name    string
		plugin  *Plugin
		wantErr error
	}{
		{
			name: "Invalid action",
			plugin: &Plugin{
				action: "hey",
			},
			wantErr: ErrorInvalidAction,
		},
		{
			name: "Empty playbook",
			plugin: &Plugin{
				action: AnsiblePlaybook,
				Playbook: &playbook.Playbook{
					Playbook: "testdata/empty.yml",
					Options:  &playbook.Options{},
				},
			},
			wantErr: playbook.ErrorEmptyPlaybook,
		},
		{
			name: "Invalid playbook path",
			plugin: &Plugin{
				action: AnsiblePlaybook,
				Playbook: &playbook.Playbook{
					Playbook: "notfound.yml",
					Options:  &playbook.Options{},
				},
			},
			wantErr: playbook.ErrorInvalidPlaybook,
		},
		{
			name: "No playbook provided",
			plugin: &Plugin{
				action: AnsiblePlaybook,
				Playbook: &playbook.Playbook{
					Playbook: "",
					Options:  &playbook.Options{},
				},
			},
			wantErr: playbook.ErrorMissingPlaybook,
		},
		{
			name: "No playbook provided",
			plugin: &Plugin{
				action: AnsiblePlaybook,
				Playbook: &playbook.Playbook{
					Options: &playbook.Options{},
				},
			},
			wantErr: playbook.ErrorMissingPlaybook,
		},
		{
			name: "Empty inventory",
			plugin: &Plugin{
				action: AnsiblePlaybook,
				Playbook: &playbook.Playbook{
					Playbook: "testdata/main.yml",
					Options: &playbook.Options{
						Inventory: []string{"testdata/empty.yml"},
					},
				},
			},
			wantErr: playbook.ErrorEmptyInventory,
		},
		{
			name: "Invalid inventory",
			plugin: &Plugin{
				action: AnsiblePlaybook,
				Playbook: &playbook.Playbook{
					Playbook: "testdata/main.yml",
					Options: &playbook.Options{
						Inventory: []string{" "},
					},
				},
			},
			wantErr: playbook.ErrorInvalidInventory,
		},
		{
			name: "Inventory not provided",
			plugin: &Plugin{
				action: AnsiblePlaybook,
				Playbook: &playbook.Playbook{
					Playbook: "testdata/main.yml",
					Options:  &playbook.Options{},
				},
			},
			wantErr: playbook.ErrorMissingInventory,
		},
		{
			name: "Empty playbook",
			plugin: &Plugin{
				action: AnsibleLint,
				Lint: &lint.Linter{
					LintPlaybook: "testdata/empty.yml",
				},
			},
			wantErr: lint.ErrorEmptyLintPlaybook,
		},
		{
			name: "Playbook path doesn't exist",
			plugin: &Plugin{
				action: AnsibleLint,
				Lint: &lint.Linter{
					LintPlaybook: "notfound.yml",
				},
			},
			wantErr: lint.ErrorInvalidLintPlaybook,
		},
		{
			name: "Playbook not provided",
			plugin: &Plugin{
				action: AnsibleLint,
				Lint:   &lint.Linter{},
			},
			wantErr: lint.ErrorMissingLintPlaybook,
		},
		{
			name: "Playbook empty string",
			plugin: &Plugin{
				action: AnsibleLint,
				Lint: &lint.Linter{
					LintPlaybook: "",
				},
			},
			wantErr: lint.ErrorMissingLintPlaybook,
		},
	}

	for _, test := range tests {
		err := test.plugin.Validate()

		if err == nil {
			t.Errorf("should have returned err")
		}

		if err != test.wantErr {
			t.Errorf("Should have returned error: %v, instead got error: %v", test.wantErr, err)
		}
	}
}

func TestPluginExecError(t *testing.T) {
	tests := []struct {
		name    string
		plugin  *Plugin
		wantErr error
	}{
		{
			name: "Invalid action",
			plugin: &Plugin{
				action: "",
				Playbook: &playbook.Playbook{
					Playbook: "testdata/main.yml",
					Options:  &playbook.Options{},
				},
			},
			wantErr: ErrorInvalidAction,
		},
	}

	for _, test := range tests {
		err := test.plugin.Exec()

		if err == nil {
			t.Errorf("should have returned err")
		}

		if err != test.wantErr {
			t.Errorf("Should have returned error: %v, instead got error: %v", test.wantErr, err)
		}
	}
}
