// SPDX-License-Identifier: Apache-2.0

// executor - executes ansible-playbook

package playbook

import (
	"github.com/sirupsen/logrus"
)

// Exec runs ansible-playbook with given flags.
func Exec(p *Playbook) error {
	logrus.Trace("entered plugin.playbook.Exec")
	defer logrus.Trace("exited plugin.playbook.Exec")

	logrus.Debug("running ansible-playbook with provided configuration")

	// sets ansible-playbook flags
	cmd := setFlags(p)

	logrus.Info("ansible-playbook: running")
	// execute ansible-playbook cli
	if err := execCmd(cmd); err != nil {
		return err
	}

	logrus.Info("ansible-playbook: complete")

	return nil
}
