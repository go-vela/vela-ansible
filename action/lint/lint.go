// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// lint - ansible-lint flags, struct, and binary

package lint

import (
	"github.com/urfave/cli/v2"
)

const (
	lint = "/usr/bin/ansible-lint"
)

type (
	// ansible-lint cli based on:
	// https://github.com/ansible/ansible-lint
	Linter struct {
		LintRules             []string
		LintTags              []string
		LintSkip              []string
		LintWarn              []string
		LintEnable            []string
		LintExclude           []string
		LintPlaybook          string
		LintConfig            string
		LintFormat            string
		LintProjectDir        string
		LintVersion           bool
		LintList              bool
		LintQuieter           bool
		LintParseable         bool
		LintParseableSeverity bool
		LintRulesDefault      bool
		LintShowRelativePath  bool
		LintTagsList          bool
		LintVerbose           bool
		LintNoColor           bool
		LintForceColor        bool
		LintProgressive       bool
		LintOffline           bool
	}
)

var Flags = []cli.Flag{
	// ansible-lint flags
	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_VERSION", "LINT_VERSION"},
		Name:    "lint.version",
		Usage:   "ansible-lint version",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_LIST", "LINT_LIST"},
		Name:    "lint.list",
		Usage:   "list all the rules",
	},

	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_LINT_FORMAT", "LINT_FORMAT"},
		Name:    "lint.format",
		Usage: "format used rules output, (default: rich) options: " +
			"{rich,plain,rst,codeclimate,quiet,pep8}",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_QUIETER", "LINT_QUIETER"},
		Name:    "lint.quieter",
		Usage:   "quieter, although not silent output",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_PARSEABLE", "LINT_PARSEABLE"},
		Name:    "lint.parseable",
		Usage:   "parseable output in the format of pep8",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_PARSEABLESEVERITY", "LINT_PARSEABLESEVERITY"},
		Name:    "lint.parseable.severity",
		Usage:   "parseable output including severity of rule",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_PROGRESSIVE", "LINT_PROGRESSIVE"},
		Name:    "lint.progressive",
		Usage: " Return success if it detects a reduction in number of " +
			"violations compared with previous git commit. This " +
			"feature works only in git repositories.",
	},

	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_LINT_PROJECTDIR", "LINT_PROJECTDIR"},
		Name:    "lint.project.dir",
		Usage: " Location of project/repository, autodetected based on " +
			"location of configuration file.",
	},

	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_LINT_RULE", "LINT_RULE"},
		Name:    "lint.rules",
		Usage: "specify one or more rules directories using one or more -r arguments. " +
			"Any -r flags overridethe default rules in /path/to/ansible-lint/lib/ansiblelint/rules," +
			" unless -R is also used.",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_RULESDEFAULT", "LINT_RULESDEFAULT"},
		Name:    "lint.rules.default",
		Usage: "use default rules in /path/to/ansible-lint/lib/ansiblelint/rules in addition" +
			" to any extra rules directories specified with -r. There is no need to specify this if no" +
			" -r flags are used",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_SHOWRELATIVEPATH", "LINT_SHOWRELATIVEPATH"},
		Name:    "lint.show.relative.path",
		Usage:   "display path relative to CWD",
	},

	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_LINT_TAGS", "LINT_TAGS"},
		Name:    "lint.tags",
		Usage:   "only check rules whose id/tags match these values",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_TAGSLIST", "LINT_TAGSLIST"},
		Name:    "lint.tags.list",
		Usage:   "list all the tags",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_VERBOSE", "LINT_VERBOSE"},
		Name:    "lint.verbose",
		Usage:   "increase verbosity level",
	},

	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_LINT_SKIP", "LINT_SKIP"},
		Name:    "lint.skip",
		Usage:   "only check rules whose id/tags do not match these values",
	},

	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_LINT_WARN", "LINT_WARN"},
		Name:    "lint.warn",
		Usage: "only warn about these rules, unless overridden in " +
			"config file defaults to 'experimental'",
	},

	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_LINT_ENABLE", "LINT_ENABLE"},
		Name:    "lint.enable",
		Usage:   "activate optional rules by their tag name",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_NOCOLOR", "LINT_NOCOLOR"},
		Name:    "lint.no.color",
		Usage:   "disable colored output",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_FORCECOLOR", "LINT_FORCECOLOR"},
		Name:    "lint.force.color",
		Usage:   "try force colored output (relying on ansible's code)",
	},

	&cli.StringSliceFlag{
		EnvVars: []string{"PARAMETER_LINT_EXCLUDE", "LINT_EXCLUDE"},
		Name:    "lint.exclude",
		Usage:   "path to directories or files to skip. this option is repeatable",
	},

	&cli.StringFlag{
		EnvVars: []string{"PARAMETER_LINT_CONFIG", "LINT_CONFIG"},
		Name:    "lint.config",
		Usage:   "specify configuration file to use. defaults to \".ansible-lint\"",
	},

	&cli.BoolFlag{
		EnvVars: []string{"PARAMETER_LINT_OFFLINE", "LINT_OFFLINE"},
		Name:    "lint.offline",
		Usage:   "Disable installation of requirements.yml",
	},
}
