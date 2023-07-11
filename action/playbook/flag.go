// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// flag - sets flags for ansible-playbook and creates ansible-playbook cli

package playbook

import (
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// nolint: funlen, gocyclo // ignore statements for flags
// setFlags creates the command line with flags to run ansible-playbook.
func setFlags(p *Playbook) *exec.Cmd {
	logrus.Trace("entered plugin.playbook.setFlags")
	defer logrus.Trace("exited plugin.playbook.setFlags")

	var flags []string

	if p.Options.Version {
		logrus.Info("ansible-playbook: command created")
		return exec.Command(playbook, "--version")
	}

	// check if playbook is provided
	if len(p.Playbook) > 0 {
		flags = append(flags, p.Playbook)
	}

	if p.Options.AskVaultPass {
		flags = append(flags, "--ask-vault-pass")
	}

	if p.Options.FlushCache {
		flags = append(flags, "--flush-cache")
	}

	if p.Options.ForceHandlers {
		flags = append(flags, "--force-handlers")
	}

	if p.Options.ListHosts {
		flags = append(flags, "--list-hosts")
	}

	if p.Options.ListTasks {
		flags = append(flags, "--list-tasks")
	}

	if len(p.Options.SkipTags) > 0 {
		flags = append(flags, "--skip-tags", strings.Join(p.Options.SkipTags, ","))
	}

	if len(p.Options.StartAtTask) > 0 {
		flags = append(flags, "--start-at-task", p.Options.StartAtTask)
	}

	if p.Options.Step {
		flags = append(flags, "--step")
	}

	if p.Options.SyntaxCheck {
		flags = append(flags, "--syntax-check")
	}

	if len(p.Options.VaultID) > 0 {
		flags = append(flags, "--vault-id", p.Options.VaultID)
	}

	if len(p.Options.VaultPasswordFile) > 0 {
		flags = append(flags, "--vault-password-file", p.Options.VaultPasswordFile)
	}

	if p.Options.Check {
		flags = append(flags, "--check")
	}

	if p.Options.Difference {
		flags = append(flags, "--diff")
	}

	if len(p.Options.ModulePath) > 0 {
		flags = append(flags, "--module-path", p.Options.ModulePath)
	}

	if len(p.Options.ExtraVars) > 0 {
		flags = append(flags, "--extra-vars", strings.Join(p.Options.ExtraVars, ", "))
	}

	if len(p.Options.Forks) > 0 {
		flags = append(flags, "--forks", p.Options.Forks)
	}

	if len(p.Options.Inventory) > 0 {
		flags = append(flags, "--inventory", strings.Join(p.Options.Inventory, ", "))
	}

	if len(p.Options.Limit) > 0 {
		flags = append(flags, "--limit", p.Options.Limit)
	}

	if len(p.Options.Tags) > 0 {
		flags = append(flags, "--tags", strings.Join(p.Options.Tags, ","))
	}

	if p.Options.Verbose {
		flags = append(flags, "--verbose")
	}

	if p.Options.VerboseMore {
		flags = append(flags, "-vvv")
	}

	if p.Options.VerboseDebug {
		flags = append(flags, "-vvvv")
	}

	if len(p.Connection.PrivateKey) > 0 {
		flags = append(flags, "--private-key", p.Connection.PrivateKey)
	}

	if len(p.Connection.SCPExtraArgs) > 0 {
		flags = append(flags, "--scp-extra-args", strings.Join(p.Connection.SCPExtraArgs, ","))
	}

	if len(p.Connection.SFTPExtraArgs) > 0 {
		flags = append(flags, "--sftp-extra-args", strings.Join(p.Connection.SFTPExtraArgs, ","))
	}

	if len(p.Connection.SSHCommonArgs) > 0 {
		flags = append(flags, "--ssh-common-args", strings.Join(p.Connection.SSHCommonArgs, ","))
	}

	if len(p.Connection.SSHExtraArgs) > 0 {
		flags = append(flags, "--ssh-extra-args", strings.Join(p.Connection.SSHExtraArgs, ","))
	}

	if len(p.Connection.Timeout) > 0 {
		flags = append(flags, "--timeout", p.Connection.Timeout)
	}

	if len(p.Connection.Connection) > 0 {
		flags = append(flags, "--connection", p.Connection.Connection)
	}

	if p.Connection.AskPass {
		flags = append(flags, "--ask-pass")
	}

	if len(p.Connection.PasswordFile) > 0 {
		flags = append(flags, "--connection-password-file", p.Connection.PasswordFile)
	}

	if len(p.Connection.User) > 0 {
		flags = append(flags, "--user", p.Connection.User)
	}

	if len(p.Privilege.BecomeMethod) > 0 {
		flags = append(flags, "--become-method", p.Privilege.BecomeMethod)
	}

	if len(p.Privilege.BecomeUser) > 0 {
		flags = append(flags, "--become-user", p.Privilege.BecomeUser)
	}

	if p.Privilege.AskBecomePass {
		flags = append(flags, "--ask-become-pass")
	}

	if len(p.Privilege.BecomePasswordFile) > 0 {
		flags = append(flags, "--become-password-file", p.Privilege.BecomePasswordFile)
	}

	if p.Privilege.Become {
		flags = append(flags, "--become")
	}

	logrus.Info("ansible-playbook: command created")

	// ansible-playbook cli
	return exec.Command(playbook, flags...)
}
