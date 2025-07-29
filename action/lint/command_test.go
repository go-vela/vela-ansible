// SPDX-License-Identifier: Apache-2.0

package lint

import (
	"os/exec"
	"testing"
)

func TestCommand_AnsibleLint_Valid(t *testing.T) {
	e := exec.CommandContext(t.Context(), "echo", "hello")

	err := execCmd(e)
	if err != nil {
		t.Errorf("execCmd returned err: %v", err)
	}
}
