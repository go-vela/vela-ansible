// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// executor - executes ansible-lint

package lint

import (
	"github.com/sirupsen/logrus"
)

// Exec runs ansible-lint with given flags.
func Exec(l *Linter) error {
	logrus.Trace("entered plugin.lint.Exec")
	defer logrus.Trace("exited plugin.lint.Exec")

	logrus.Debug("running ansible-lint with provided configuration")

	// sets ansible-lint flags
	cmd := setFlags(l)

	logrus.Info("ansible-lint: running")
	// execute ansible-lint cli
	if err := execCmd(cmd); err != nil {
		return err
	}

	logrus.Info("ansible-lint: complete")

	return nil
}
