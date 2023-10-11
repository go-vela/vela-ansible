// SPDX-License-Identifier: Apache-2.0

package playbook

import (
	"errors"
	"testing"
)

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		playbook *Playbook
	}{
		{
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{"testdata/inventory/hosts.yml"},
				},
			},
		},
		{
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{"10.10.10.10"},
				},
			},
		},
		{
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{"10.10.10.10"},
					Version:   true,
				},
			},
		},
	}

	for _, test := range tests {
		if err := Validate(test.playbook); err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}

func TestValidateError(t *testing.T) {
	tests := []struct {
		name     string
		playbook *Playbook
		wantErr  error
	}{
		{
			name: "Empty playbook",
			playbook: &Playbook{
				Playbook: "testdata/empty.yml",
				Options:  &Options{},
			},
			wantErr: ErrorEmptyPlaybook,
		},
		{
			name: "Invalid playbook path",
			playbook: &Playbook{
				Playbook: "notfound.yml",
				Options:  &Options{},
			},
			wantErr: ErrorInvalidPlaybook,
		},
		{
			name: "No playbook provided",
			playbook: &Playbook{
				Playbook: "",
				Options:  &Options{},
			},
			wantErr: ErrorMissingPlaybook,
		},
		{
			name: "No playbook provided",
			playbook: &Playbook{
				Options: &Options{},
			},
			wantErr: ErrorMissingPlaybook,
		},
		{
			name: "Empty inventory",
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{"testdata/inventory/empty.yml"},
				},
			},
			wantErr: ErrorEmptyInventory,
		},
		{
			name: "Invalid inventory",
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{"notfound.yml"},
				},
			},
			wantErr: ErrorInvalidInventory,
		},
		{
			name: "Inventory empty string",
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{""},
				},
			},
			wantErr: ErrorInvalidInventory,
		},
		{
			name: "Inventory not provided",
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options:  &Options{},
			},
			wantErr: ErrorMissingInventory,
		},
		{
			name: "Invalid ip address",
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{"256.256.256.256"},
				},
			},
			wantErr: ErrorInvalidInventory,
		},
	}

	for _, test := range tests {
		err := Validate(test.playbook)

		if err == nil {
			t.Errorf("should have returned err")
		}

		if !errors.Is(err, test.wantErr) {
			t.Errorf("Should have returned error: %v, instead got error: %v", test.wantErr, err)
		}
	}
}
