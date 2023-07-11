// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package lint

import (
	"os/exec"
	"reflect"
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
			expected: exec.Command(lint, "testdata/main.yml",
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
			expected: exec.Command(lint, "testdata/main.yml", "-v", "-x",
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
			expected: exec.Command(lint, "--version"),
		},
	}

	for _, test := range tests {
		command := setFlags(test.lint)

		if !reflect.DeepEqual(command, test.expected) {
			t.Errorf("Command is %v, want %v", command, test.expected)
		}
	}
}
