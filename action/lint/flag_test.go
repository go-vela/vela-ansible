// SPDX-License-Identifier: Apache-2.0

package lint

import (
	"context"
	"os/exec"
	"testing"
)

func TestSetFlags(t *testing.T) {
	tests := []struct {
		lint     *Linter
		expected *exec.Cmd
	}{
		{
			lint: &Linter{
				LintPlaybook:          "testdata/main.yml",
				LintList:              true,
				LintFormat:            "plain",
				LintQuieter:           true,
				LintParseable:         true,
				LintParseableSeverity: true,
				LintProgressive:       true,
				LintProjectDir:        "vela-ansible/",
				LintRules:             []string{"no-tabs", "yaml", "role-name"},
				LintRulesDefault:      true,
				LintShowRelativePath:  true,
				LintTags:              []string{"shell", "yaml"},
				LintTagsList:          true,
				LintForceColor:        true,
			},
			expected: exec.CommandContext(context.Background(), lint, "testdata/main.yml",
				"-L", "-f", "plain", "-q", "-p", "--parseable-severity",
				"--progressive", "--project-dir", "vela-ansible/", "-r",
				"no-tabs,yaml,role-name", "-R", "--show-relpath", "-t",
				"shell,yaml", "-T", "--force-color"),
		},
		{
			lint: &Linter{
				LintPlaybook: "testdata/main.yml",
				LintVerbose:  true,
				LintSkip:     []string{"security"},
				LintWarn:     []string{"experimental", "metadata"},
				LintEnable:   []string{"no-tabs"},
				LintNoColor:  true,
				LintExclude:  []string{"/path1/", "testdata/path2/"},
				LintConfig:   ".ansible-lint",
				LintOffline:  true,
			},
			expected: exec.CommandContext(context.Background(), lint, "testdata/main.yml", "-v", "-x",
				"security", "-w", "experimental,metadata", "--enable-list",
				"no-tabs", "--nocolor", "--exclude", "/path1/,testdata/path2/",
				"-c", ".ansible-lint", "--offline"),
		},
		{
			lint: &Linter{
				LintPlaybook: "testdata/main.yml",
				LintVersion:  true,
				LintList:     true,
				LintQuieter:  true,
			},
			expected: exec.CommandContext(context.Background(), lint, "--version"),
		},
	}

	for _, test := range tests {
		command := setFlags(context.Background(), test.lint)

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
