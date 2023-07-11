// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package playbook

import (
	"os/exec"
	"testing"
)

func TestCommand_AnsiblePlaybook_Valid(t *testing.T) {
	e := exec.Command("echo", "hello")

	err := execCmd(e)
	if err != nil {
		t.Errorf("execCmd returned err: %v", err)
	}
}
