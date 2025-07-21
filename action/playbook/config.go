// SPDX-License-Identifier: Apache-2.0

// config - ansible-playbook configuration

package playbook

import (
	"github.com/urfave/cli/v3"
)

// apply Playbook configuration.
func Config(c *cli.Command) *Playbook {
	return &Playbook{
		// playbook
		Playbook: c.String("playbook"),
		// playbook Options configurations
		Options: &Options{
			AskVaultPass:      c.Bool("options.ask.vault.pass"),
			FlushCache:        c.Bool("options.flush.cache"),
			ForceHandlers:     c.Bool("options.force.handlers"),
			ListHosts:         c.Bool("options.list.hosts"),
			ListTags:          c.Bool("options.list.tags"),
			ListTasks:         c.Bool("options.list.tasks"),
			SkipTags:          c.StringSlice("options.skip.tags"),
			StartAtTask:       c.String("options.start.at.task"),
			Step:              c.Bool("options.step"),
			SyntaxCheck:       c.Bool("options.syntax.check"),
			VaultID:           c.String("options.vault.id"),
			VaultPasswordFile: c.String("options.vault.password.file"),
			Version:           c.Bool("options.version"),
			Check:             c.Bool("options.check"),
			Difference:        c.Bool("options.difference"),
			ModulePath:        c.String("options.module.path"),
			ExtraVars:         c.StringSlice("options.extra.vars"),
			Forks:             c.String("options.forks"),
			Inventory:         c.StringSlice("options.inventory"),
			Limit:             c.String("options.limit"),
			Tags:              c.StringSlice("options.tags"),
			Verbose:           c.Bool("options.verbose"),
			VerboseMore:       c.Bool("options.verbose.more"),
			VerboseDebug:      c.Bool("options.verbose.debug"),
		},
		// playbook Connection configurations
		Connection: &Connection{
			PrivateKey:    c.String("connection.private.key"),
			SCPExtraArgs:  c.StringSlice("connection.scp.extra.args"),
			SFTPExtraArgs: c.StringSlice("connection.sftp.extra.args"),
			SSHCommonArgs: c.StringSlice("connection.ssh.common.args"),
			SSHExtraArgs:  c.StringSlice("connection.ssh.extra.args"),
			Timeout:       c.String("connection.timeout"),
			Connection:    c.String("connection.connection"),
			AskPass:       c.Bool("connection.ask.pass"),
			PasswordFile:  c.String("connection.pass.file"),
			User:          c.String("connection.user"),
		},
		// playbook Privilege configurations
		Privilege: &Privilege{
			BecomeMethod:       c.String("privilege.become.method"),
			BecomeUser:         c.String("privilege.become.user"),
			AskBecomePass:      c.Bool("privilege.ask.become.pass"),
			BecomePasswordFile: c.String("privilege.become.pass.file"),
			Become:             c.Bool("privilege.become"),
		},
	}
}
