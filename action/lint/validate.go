// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// validate - validates ansible-lint to make sure that a playbook is defined.

package lint

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// ErrorEmptyLintPlaybook is returned when the playbook provided is empty.
	ErrorEmptyLintPlaybook = errors.New("ansible-lint: playbook is empty")

	// ErrorInvalidLintPlaybook is returned when the playbook path provided is invalid.
	ErrorInvalidLintPlaybook = errors.New("ansible-lint: playbook invalid path")

	// ErrorMissingLintPlaybook is returned when there is no playbook provided.
	ErrorMissingLintPlaybook = errors.New("ansible-lint: playbook not specified")
)

// Validate ansible-lint.
func Validate(l *Linter) error {
	logrus.Trace("entered plugin.lint.Validate")
	defer logrus.Trace("exited plugin.lint.Validate")

	logrus.Debug("validating ansible-lint configuration")

	if l.LintVersion {
		return nil
	}

	// check for valid playbook
	logrus.Info("validating ansible-lint")

	if len(l.LintPlaybook) != 0 {
		file, err := os.Stat(l.LintPlaybook)
		if err != nil {
			return ErrorInvalidLintPlaybook
		}

		if file.Size() == 0 {
			return ErrorEmptyLintPlaybook
		}
	} else {
		return ErrorMissingLintPlaybook
	}

	return nil
}
