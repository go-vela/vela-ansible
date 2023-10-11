// SPDX-License-Identifier: Apache-2.0

// command - runs ansible-playbook cli

package playbook

import (
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// execCmd is a helper function to run the provided command.
func execCmd(e *exec.Cmd) error {
	logrus.Tracef("executing cmd %s", strings.Join(e.Args, " "))

	// set command stdout to OS stdout
	e.Stdout = os.Stdout
	// set command stderr to OS stderr
	e.Stderr = os.Stderr

	// runs ansible-lint cli
	return e.Run()
}
