// SPDX-License-Identifier: Apache-2.0

package playbook

import (
	"context"
	"os/exec"
	"testing"
)

func TestSetFlags(t *testing.T) {
	tests := []struct {
		playbook *Playbook
		expected *exec.Cmd
	}{
		{
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory:         []string{"testdata/inventory/hosts.yml"},
					AskVaultPass:      true,
					FlushCache:        true,
					ForceHandlers:     true,
					ListHosts:         true,
					ListTasks:         true,
					SkipTags:          []string{"security", "metadata"},
					StartAtTask:       "yaml",
					Step:              true,
					SyntaxCheck:       true,
					VaultID:           "hey",
					VaultPasswordFile: "test/pass",
					Check:             true,
					Difference:        true,
					ModulePath:        ".ansible/plugins/",
				},
				Connection: &Connection{},
				Privilege:  &Privilege{},
			},
			expected: exec.CommandContext(context.Background(), playbook, "testdata/main.yml", "--ask-vault-pass",
				"--flush-cache", "--force-handlers", "--list-hosts", "--list-tasks",
				"--skip-tags", "security,metadata", "--start-at-task", "yaml", "--step",
				"--syntax-check", "--vault-id", "hey", "--vault-password-file", "test/pass",
				"--check", "--diff", "--module-path", ".ansible/plugins/", "--inventory",
				"testdata/inventory/hosts.yml"),
		},
		{
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory:    []string{"10.10.10.10", "10.10.10.11"},
					ExtraVars:    []string{"ansible_config=config"},
					Forks:        "4",
					Limit:        "pattern",
					Tags:         []string{"tag1", "tag2"},
					Verbose:      true,
					VerboseMore:  true,
					VerboseDebug: true,
				},
				Connection: &Connection{},
				Privilege:  &Privilege{},
			},
			expected: exec.CommandContext(context.Background(), playbook, "testdata/main.yml", "--extra-vars", "ansible_config=config",
				"--forks", "4", "--inventory", "10.10.10.10, 10.10.10.11", "--limit", "pattern",
				"--tags", "tag1,tag2", "--verbose", "-vvv", "-vvvv"),
		},
		{
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{"10.10.10.10", "10.10.10.11"},
				},
				Connection: &Connection{
					PrivateKey:    "privkey",
					SCPExtraArgs:  []string{"-l"},
					SFTPExtraArgs: []string{"-f", "-l"},
					SSHCommonArgs: []string{"ProxyCommand"},
					SSHExtraArgs:  []string{"-R"},
					Timeout:       "9",
					Connection:    "smart",
					AskPass:       true,
					PasswordFile:  "passfile",
					User:          "None",
				},
				Privilege: &Privilege{},
			},
			expected: exec.CommandContext(context.Background(), playbook, "testdata/main.yml", "--inventory",
				"10.10.10.10, 10.10.10.11", "--private-key", "privkey", "--scp-extra-args",
				"-l", "--sftp-extra-args", "-f,-l", "--ssh-common-args", "ProxyCommand",
				"--ssh-extra-args", "-R", "--timeout", "9", "--connection", "smart",
				"--ask-pass", "--connection-password-file", "passfile", "--user", "None"),
		},
		{
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Inventory: []string{"10.10.10.10", "10.10.10.11"},
				},
				Connection: &Connection{},
				Privilege: &Privilege{
					BecomeMethod:       "sudo",
					BecomeUser:         "root",
					AskBecomePass:      true,
					BecomePasswordFile: "passfile",
					Become:             true,
				},
			},
			expected: exec.CommandContext(context.Background(), playbook, "testdata/main.yml", "--inventory",
				"10.10.10.10, 10.10.10.11", "--become-method", "sudo", "--become-user",
				"root", "--ask-become-pass", "--become-password-file", "passfile",
				"--become"),
		},
		{
			playbook: &Playbook{
				Playbook: "testdata/main.yml",
				Options: &Options{
					Version: true,
				},
				Connection: &Connection{},
				Privilege:  &Privilege{},
			},
			expected: exec.CommandContext(context.Background(), playbook, "--version"),
		},
	}

	for _, test := range tests {
		command := setFlags(context.Background(), test.playbook)

		if command.Path != test.expected.Path {
			t.Errorf("Command path is %v, want %v", command.Path, test.expected.Path)
		}

		if len(command.Args) != len(test.expected.Args) {
			t.Errorf("Command args length is %d, want %d", len(command.Args), len(test.expected.Args))
		}

		for i, arg := range command.Args {
			if i < len(test.expected.Args) && arg != test.expected.Args[i] {
				t.Errorf("Command arg[%d] is %v, want %v", i, arg, test.expected.Args[i])
			}
		}
	}
}
