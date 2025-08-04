// SPDX-License-Identifier: Apache-2.0

package playbook

import (
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		failure  bool
		playbook *Playbook
	}{
		{
			failure: true,
			playbook: &Playbook{
				Playbook:   "",
				Options:    &Options{},
				Connection: &Connection{},
				Privilege:  &Privilege{},
			},
		},
		{
			failure: true,
			playbook: &Playbook{
				Playbook:   "testdata/main.yml",
				Options:    &Options{},
				Connection: &Connection{},
				Privilege:  &Privilege{},
			},
		},
	}

	for _, test := range tests {
		err := Exec(t.Context(), test.playbook)

		if test.failure {
			if err == nil {
				t.Errorf("setFlags should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("setFlags returned err: %v", err)
		}
	}
}
