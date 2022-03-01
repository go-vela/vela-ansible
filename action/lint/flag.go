// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// flag - sets flags for ansible-lint and creates ansible-lint cli

package lint

import (
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// nolint:funlen // false positive
// Command creates the command line with flags to run ansible-lint.
func setFlags(l *Linter) *exec.Cmd {
	logrus.Trace("entered plugin.lint.setFlags")
	defer logrus.Trace("exited plugin.lint.setFlags")

	var flags []string

	if l.LintVersion {
		logrus.Info("ansible-lint: command created")
		return exec.Command(lint, "--version")
	}

	// check if lint.playbook is provided
	if len(l.LintPlaybook) > 0 {
		flags = append(flags, l.LintPlaybook)
	}

	if l.LintList {
		flags = append(flags, "-L")
	}

	if len(l.LintFormat) > 0 {
		flags = append(flags, "-f", l.LintFormat)
	}

	if l.LintQuieter {
		flags = append(flags, "-q")
	}

	if l.LintParseable {
		flags = append(flags, "-p")
	}

	if l.LintParseableSeverity {
		flags = append(flags, "--parseable-severity")
	}

	if l.LintProgressive {
		flags = append(flags, "--progressive")
	}

	if len(l.LintProjectDir) > 0 {
		flags = append(flags, "--project-dir", l.LintProjectDir)
	}

	if len(l.LintRules) > 0 {
		flags = append(flags, "-r", strings.Join(l.LintRules, ","))
	}

	if l.LintRulesDefault {
		flags = append(flags, "-R")
	}

	if l.LintShowRelativePath {
		flags = append(flags, "--show-relpath")
	}

	if len(l.LintTags) > 0 {
		flags = append(flags, "-t", strings.Join(l.LintTags, ","))
	}

	if l.LintTagsList {
		flags = append(flags, "-T")
	}

	if l.LintVerbose {
		flags = append(flags, "-v")
	}

	if len(l.LintSkip) > 0 {
		flags = append(flags, "-x", strings.Join(l.LintSkip, ","))
	}

	if len(l.LintWarn) > 0 {
		flags = append(flags, "-w", strings.Join(l.LintWarn, ","))
	}

	if len(l.LintEnable) > 0 {
		flags = append(flags, "--enable-list", strings.Join(l.LintEnable, ","))
	}

	if l.LintNoColor {
		flags = append(flags, "--nocolor")
	}

	if l.LintForceColor {
		flags = append(flags, "--force-color")
	}

	if len(l.LintExclude) > 0 {
		flags = append(flags, "--exclude", strings.Join(l.LintExclude, ","))
	}

	if len(l.LintConfig) > 0 {
		flags = append(flags, "-c", l.LintConfig)
	}

	if l.LintOffline {
		flags = append(flags, "--offline")
	}

	logrus.Info("ansible-lint: command created")
	
	// ansible-lint cli
	return exec.Command(lint, flags...)
}
