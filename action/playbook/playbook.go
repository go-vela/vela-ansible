// SPDX-License-Identifier: Apache-2.0

// playbook - ansible-playbook flags, struct, and binary

package playbook

import (
	"github.com/urfave/cli/v3"
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
		Name:    "playbook",
		Usage:   "playbook",
		Sources: cli.EnvVars("PARAMETER_PLAYBOOK"),
	},
	// options flags
	&cli.BoolFlag{
		Name:    "options.ask.vault.pass",
		Usage:   "ask for vault password",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_ASKVAULTPASS", "PLAYBOOK_OPTIONS_ASKVAULTPASS"),
	},
	&cli.BoolFlag{
		Name:    "options.flush.cache",
		Usage:   "clear the fact cache for every host in inventory",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_FLUSHCACHE", "PLAYBOOK_OPTIONS_FLUSHCACHE"),
	},
	&cli.BoolFlag{
		Name:    "options.force.handlers",
		Usage:   "run handlers even if a task fails",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_FORCEHANDLERS", "PLAYBOOK_OPTIONS_FORCEHANDLERS"),
	},
	&cli.BoolFlag{
		Name:    "options.list.hosts",
		Usage:   "outputs a list of matching hosts; does not execute anything else",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_LISTHOSTS", "PLAYBOOK_OPTIONS_LISTHOSTS"),
	},
	&cli.BoolFlag{
		Name:    "options.list.tags",
		Usage:   "list all available tags",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_LISTTAGS", "PLAYBOOK_OPTIONS_LISTTAGS"),
	},
	&cli.BoolFlag{
		Name:    "options.list.tasks",
		Usage:   "list all tasks that would be executed",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_LISTTASKS", "PLAYBOOK_OPTIONS_LISTTASKS"),
	},
	&cli.StringSliceFlag{
		Name:    "options.skip.tags",
		Usage:   "only run plays and tasks whose tag do not match these value",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_SKIPTAGS", "PLAYBOOK_OPTIONS_SKIPTAGS"),
	},
	&cli.StringFlag{
		Name:    "options.start.at.task",
		Usage:   "start the playbook at the task matching this name",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_STARTATTASK", "PLAYBOOK_OPTIONS_STARTATTASK"),
	},
	&cli.BoolFlag{
		Name:    "options.step",
		Usage:   "one-step-at-a-time: confirm each task before running",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_STEP", "PLAYBOOK_OPTIONS_STEP"),
	},
	&cli.BoolFlag{
		Name:    "options.syntax.check",
		Usage:   "perform a syntax check on the playbook, but do not execute it",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_SYNTAXCHECK", "PLAYBOOK_OPTIONS_SYNTAXCHECK"),
	},
	&cli.StringFlag{
		Name:    "options.vault.id",
		Usage:   "the vault identity to use",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_VAULTID", "PLAYBOOK_OPTIONS_VAULTID"),
	},
	&cli.StringFlag{
		Name:    "options.vault.password.file",
		Usage:   "vault password file",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_VAULTPASSWORDFILE", "PLAYBOOK_OPTIONS_VAULTPASSWORDFILE"),
	},
	&cli.BoolFlag{
		Name: "options.version",
		Usage: "show program's version number, config file location, configured module search path, " +
			"module location executable location and exit",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_VERSION", "PLAYBOOK_OPTIONS_VERSION"),
	},
	&cli.BoolFlag{
		Name:    "options.check",
		Usage:   "don't make any changes; instead, try to predict some of the changes that may occur",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_CHECK", "PLAYBOOK_OPTIONS_CHECK"),
	},
	&cli.BoolFlag{
		Name: "options.difference",
		Usage: "when changing (small) files and template, show the difference in those files; " +
			"works great with --check",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_DIFFERENCE", "PLAYBOOK_OPTIONS_DIFFERENCE"),
	},
	&cli.StringFlag{
		Name: "options.module.path",
		Usage: "prepend colon-separated path(s) to module library " +
			"(default=~/.ansible/plugins/modules:/usr/share/ansible/plugins/modules)",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_MODULEPATH", "PLAYBOOK_OPTIONS_MODULEPATH"),
	},
	&cli.StringSliceFlag{
		Name:    "options.extra.vars",
		Usage:   "set additional variables as key=value or YAML/JSON, if filename prepend with @",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_EXTRAVARS", "PLAYBOOK_OPTIONS_EXTRAVARS"),
	},
	&cli.StringFlag{
		Name:    "options.forks",
		Usage:   "specify number of parallel proccesses to use (default=5)",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_FORKS", "PLAYBOOK_OPTIONS_FORKS"),
	},
	&cli.StringSliceFlag{
		Name:    "options.inventory",
		Usage:   "specify inventory host path or comma separated host list",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_INVENTORY", "PLAYBOOK_OPTIONS_INVENTORY"),
	},
	&cli.StringFlag{
		Name:    "options.limit",
		Usage:   "further limit selected hosts to additional pattern",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_LIMIT", "PLAYBOOK_OPTIONS_LIMIT"),
	},
	&cli.StringSliceFlag{
		Name:    "options.tags",
		Usage:   "only run plays and tasks tagged with these values",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_TAGS", "PLAYBOOK_OPTIONS_TAGS"),
	},
	&cli.BoolFlag{
		Name:    "options.verbose",
		Usage:   "verbose mode",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_VERBOSE", "PLAYBOOK_OPTIONS_VERBOSE"),
	},
	&cli.BoolFlag{
		Name:    "options.verbose.more",
		Usage:   "verbose mode: more verbose",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_VERBOSEMORE", "PLAYBOOK_OPTIONS_VERBOSEMORE"),
	},
	&cli.BoolFlag{
		Name:    "options.verbose.debug",
		Usage:   "verbose mode: connection debugging",
		Sources: cli.EnvVars("PARAMETER_OPTIONS_VERBOSEDEBUG", "PLAYBOOK_OPTIONS_VERBOSEDEBUG"),
	},
	// connection flags
	&cli.StringFlag{
		Name:    "connection.private.key",
		Usage:   "use this file to authenticate the connection",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_PRIVATEKEY", "PLAYBOOK_CONNECTION_PRIVATEKEY"),
	},
	&cli.StringSliceFlag{
		Name:    "connection.scp.extra.args",
		Usage:   "specify extra arguments to pass to scp only",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_SCPEXTRAARGS", "PLAYBOOK_CONNECTION_SCPEXTRAARGS"),
	},
	&cli.StringSliceFlag{
		Name:    "connection.sftp.extra.args",
		Usage:   "specify extra arguments to pass to sfpt only",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_SFTPEXTRAARGS", "PLAYBOOK_CONNECTION_SFTPEXTRAARGS"),
	},
	&cli.StringSliceFlag{
		Name:    "connection.ssh.common.args",
		Usage:   "specify common arguments to pass to sftp/scp/ssh",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_SSHCOMMONARGS", "PLAYBOOK_CONNECTION_SSHCOMMONARG"),
	},
	&cli.StringSliceFlag{
		Name:    "connection.ssh.extra.args",
		Usage:   "specify extra arguments to pass to ssh only",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_SSHEXTRAARGS", "PLAYBOOK_CONNECTION_SSHEXTRAARGS"),
	},
	&cli.StringFlag{
		Name:    "connection.timeout",
		Usage:   "override the connection timeout in seconds (default=10)",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_TIMEOUT", "PLAYBOOK_CONNECTION_TIMEOUT"),
	},
	&cli.StringFlag{
		Name:    "connection.connection",
		Usage:   "connection type to use (default=smart)",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_CONNECTION", "PLAYBOOK_CONNECTION_CONNECTION"),
	},
	&cli.StringFlag{
		Name:    "connection.user",
		Usage:   "connect as this user (default=none)",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_USER", "PLAYBOOK_CONNECTION_USER"),
	},
	&cli.StringFlag{
		Name:    "connection.password.file",
		Usage:   "connection password file",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_PASSWORDFILE", "PLAYBOOK_CONNECTION_PASSWORDFILE"),
	},
	&cli.BoolFlag{
		Name:    "connection.ask.pass",
		Usage:   "ask for connection password",
		Sources: cli.EnvVars("PARAMETER_CONNECTION_ASKPASS", "PLAYBOOK_CONNECTION_ASKPASS"),
	},
	// privilege flags
	&cli.StringFlag{
		Name:    "privilege.become.method",
		Usage:   "privilege escalation method to use (default=sudo)",
		Sources: cli.EnvVars("PARAMETER_PRIVILEGE_BECOMEMETHOD", "PLAYBOOK_PRIVILEGE_BECOMEMETHOD"),
	},
	&cli.StringFlag{
		Name:    "privilege.become.user",
		Usage:   "run operation as this user (default=root)",
		Sources: cli.EnvVars("PARAMETER_PRIVILEGE_BECOMEUSER", "PLAYBOOK_PRIVILEGE_BECOMEUSER"),
	},
	&cli.BoolFlag{
		Name:    "privilege.ask.become.pass",
		Usage:   "ask for privilege escalation password",
		Sources: cli.EnvVars("PARAMETER_PRIVILEGE_ASKBECOMEPASS", "PLAYBOOK_PRIVILEGE_ASKBECOMEPASS"),
	},
	&cli.StringFlag{
		Name:    "privilege.become.password.file",
		Usage:   "become password file",
		Sources: cli.EnvVars("PARAMETER_PRIVILEGE_BECOMEPASSFILE", "PLAYBOOK_PRIVILEGE_BECOMEPASSFILE"),
	},
	&cli.BoolFlag{
		Name:    "privilege.become",
		Usage:   "run operations with become (does not imply password prompting)",
		Sources: cli.EnvVars("PARAMETER_PRIVILEGE_BECOME", "PLAYBOOK_PRIVILEGE_BECOME"),
	},
}
