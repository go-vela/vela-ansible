// SPDX-License-Identifier: Apache-2.0

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
