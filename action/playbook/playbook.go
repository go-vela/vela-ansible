// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// playbook - ansible-playbook flags, struct, and binary

package playbook

import (
	"github.com/urfave/cli/v2"
)

const (
	playbook = "/usr/bin/ansible-playbook"
)

type (
	// ansible-playbook cli based on:
	// https://docs.ansible.com/ansible/latest/cli/ansible-playbook.html
	Playbook struct {
		Playbook   string
		Options    *Options
		Connection *Connection
		Privilege  *Privilege
	}
	Options struct {
		SkipTags          []string
		ExtraVars         []string
		Inventory         []string
		Tags              []string
		StartAtTask       string
		VaultID           string
		VaultPasswordFile string
		ModulePath        string
		Forks             string
		Limit             string
		AskVaultPass      bool
		FlushCache        bool
		ForceHandlers     bool
		ListHosts         bool
		ListTags          bool
		ListTasks         bool
		Step              bool
		SyntaxCheck       bool
		Version           bool
		Check             bool
		Difference        bool
		Verbose           bool
		VerboseMore       bool
		VerboseDebug      bool
	}
	Connection struct {
		PrivateKey    string
		SCPExtraArgs  []string
		SFTPExtraArgs []string
		SSHCommonArgs []string
		SSHExtraArgs  []string
		Timeout       string
		Connection    string
		PasswordFile  string
		AskPass       bool
		User          string
	}
	Privilege struct {
		BecomeMethod       string
		BecomeUser         string
		BecomePasswordFile string
		AskBecomePass      bool
		Become             bool
	}
)

var Flags = []cli.Flag{
	// ansible-playbook flags
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_PLAYBOOK"},
		Name:    "playbook",
		Usage:   "playbook",
	},
	// options flags
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_ASKVAULTPASS", "PLAYBOOK_OPTIONS_ASKVAULTPASS"},
		Name:    "options.ask.vault.pass",
		Usage:   "ask for vault password",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_FLUSHCACHE", "PLAYBOOK_OPTIONS_FLUSHCACHE"},
		Name:    "options.flush.cache",
		Usage:   "clear the fact cache for every host in inventory",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_FORCEHANDLERS", "PLAYBOOK_OPTIONS_FORCEHANDLERS"},
		Name:    "options.force.handlers",
		Usage:   "run handlers even if a task fails",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_LISTHOSTS", "PLAYBOOK_OPTIONS_LISTHOSTS"},
		Name:    "options.list.hosts",
		Usage:   "outputs a list of matching hosts; does not execute anything else",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_LISTTAGS", "PLAYBOOK_OPTIONS_LISTTAGS"},
		Name:    "options.list.tags",
		Usage:   "list all available tags",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_LISTTASKS", "PLAYBOOK_OPTIONS_LISTTASKS"},
		Name:    "options.list.tasks",
		Usage:   "list all tasks that would be executed",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_SKIPTAGS", "PLAYBOOK_OPTIONS_SKIPTAGS"},
		Name:    "options.skip.tags",
		Usage:   "only run plays and tasks whose tag do not match these value",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_STARTATTASK", "PLAYBOOK_OPTIONS_STARTATTASK"},
		Name:    "options.start.at.task",
		Usage:   "start the playbook at the task matching this name",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_STEP", "PLAYBOOK_OPTIONS_STEP"},
		Name:    "options.step",
		Usage:   "one-step-at-a-time: confirm each task before running",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_SYNTAXCHECK", "PLAYBOOK_OPTIONS_SYNTAXCHECK"},
		Name:    "options.syntax.check",
		Usage:   "perform a syntax check on the playbook, but do not execute it",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_VAULTID", "PLAYBOOK_OPTIONS_VAULTID"},
		Name:    "options.valut.id",
		Usage:   "the vault identity to use",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_VAULTPASSWORDFILE", "PLAYBOOK_OPTIONS_VAULTPASSWORDFILE"},
		Name:    "options.vault.password.file",
		Usage:   "vault password file",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_VERSION", "PLAYBOOK_OPTIONS_VERSION"},
		Name:    "options.version",
		Usage: "show program's version number, config file location, configured module search path, " +
			"module location executable location and exit",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_CHECK", "PLAYBOOK_OPTIONS_CHECK"},
		Name:    "options.check",
		Usage:   "don't make any changes; instead, try to predict some of the changes that may occur",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_DIFFERENCE", "PLAYBOOK_OPTIONS_DIFFERENCE"},
		Name:    "options.difference",
		Usage: "when changing (small) files and template, show the difference in those files; " +
			"works great with --check",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_MODULEPATH", "PLAYBOOK_OPTIONS_MODULEPATH"},
		Name:    "options.module.path",
		Usage: "prepend colon-separated path(s) to module library " +
			"(default=~/.ansible/plugins/modules:/usr/share/ansible/plugins/modules)",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_EXTRAVARS", "PLAYBOOK_OPTIONS_EXTRAVARS"},
		Name:    "options.extra.vars",
		Usage:   "set additional variables as key=value or YAML/JSON, if filename prepend with @",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_FORKS", "PLAYBOOK_OPTIONS_FORKS"},
		Name:    "options.forks",
		Usage:   "specify number of parallel proccesses to use (default=5)",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_INVENTORY", "PLAYBOOK_OPTIONS_INVENTORY"},
		Name:    "options.inventory",
		Usage:   "specify inventory host path or comma separated host list",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_LIMIT", "PLAYBOOK_OPTIONS_LIMIT"},
		Name:    "options.limit",
		Usage:   "further limit selected hosts to additional pattern",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_TAGS", "PLAYBOOK_OPTIONS_TAGS"},
		Name:    "options.tags",
		Usage:   "only run plays and tasks tagged with these values",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_VERBOSE", "PLAYBOOK_OPTIONS_VERBOSE"},
		Name:    "options.verbose",
		Usage:   "verbose mode",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_VERBOSEMORE", "PLAYBOOK_OPTIONS_VERBOSEMORE"},
		Name:    "options.verbose.more",
		Usage:   "verbose mode: more verbose",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_OPTIONS_VERBOSEDEBUG", "PLAYBOOK_OPTIONS_VERBOSEDEBUG"},
		Name:    "options.verbose.debug",
		Usage:   "verbose mode: connection debugging",
	},
	// connection flags
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_PRIVATEKEY", "PLAYBOOK_CONNECTION_PRIVATEKEY"},
		Name:    "connection.private.key",
		Usage:   "use this file to authenticate the connection",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_SCPEXTRAARGS", "PLAYBOOK_CONNECTION_SCPEXTRAARGS"},
		Name:    "connection.scp.extra.args",
		Usage:   "specify extra arguments to pass to scp only",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_SFTPEXTRAARGS", "PLAYBOOK_CONNECTION_SFTPEXTRAARGS"},
		Name:    "connection.sftp.extra.args",
		Usage:   "specify extra arguments to pass to sfpt only",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_SSHCOMMONARGS", "PLAYBOOK_CONNECTION_SSHCOMMONARG"},
		Name:    "connection.ssh.common.args",
		Usage:   "specify common arguments to pass to sftp/scp/ssh",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_SSHEXTRAARGS", "PLAYBOOK_CONNECTION_SSHEXTRAARGS"},
		Name:    "connection.ssh.extra.args",
		Usage:   "specify extra arguments to pass to ssh only",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_TIMEOUT", "PLAYBOOK_CONNECTION_TIMEOUT"},
		Name:    "connection.timeout",
		Usage:   "override the connection timeout in seconds (default=10)",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_CONNECTION", "PLAYBOOK_CONNECTION_CONNECTION"},
		Name:    "connection.connection",
		Usage:   "connection type to use (default=smart)",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_USER", "PLAYBOOK_CONNECTION_USER"},
		Name:    "connection.user",
		Usage:   "connect as this user (default=none)",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_CONNECTION_PASSWORDFILE", "PLAYBOOK_CONNECTION_PASSWORDFILE"},
		Name:    "connection.pass.file",
		Usage:   "connection password file",
	},
	// privilege flags
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_PRIVILEGE_BECOMEMETHOD", "PLAYBOOK_PRIVILEGE_BECOMEMETHOD"},
		Name:    "privilege.become.method",
		Usage:   "privilege escalation method to use (default=sudo)",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_PRIVILEGE_BECOMEUSER", "PLAYBOOK_PRIVILEGE_BECOMEUSER"},
		Name:    "privilege.become.user",
		Usage:   "run operation as this user (default=root)",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_PRIVILEGE_ASKBECOMEPASS", "PLAYBOOK_PRIVILEGE_ASKBECOMEPASS"},
		Name:    "privilege.ask.become.pass",
		Usage:   "ask for privilege escalation password",
	},
	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_PRIVILEGE_BECOMEPASSFILE", "PLAYBOOK_PRIVILEGE_BECOMEPASSFILE"},
		Name:    "privilege.become.pass.file",
		Usage:   "become password file",
	},
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_PRIVILEGE_BECOME", "PLAYBOOK_PRIVILEGE_BECOME"},
		Name:    "privilege.become",
		Usage:   "run operations with become (does not imply password prompting)",
	},
}
