// SPDX-License-Identifier: Apache-2.0

// lint - ansible-lint flags, struct, and binary

package lint

import (
	"github.com/urfave/cli/v3"
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
		Name:    "lint.version",
		Usage:   "ansible-lint version",
		Sources: cli.EnvVars("PARAMETER_LINT_VERSION", "LINT_VERSION"),
	},

	&cli.BoolFlag{
		Name:    "lint.list",
		Usage:   "list all the rules",
		Sources: cli.EnvVars("PARAMETER_LINT_LIST", "LINT_LIST"),
	},

	&cli.StringFlag{
		Name: "lint.format",
		Usage: "format used rules output, (default: rich) options: " +
			"{rich,plain,rst,codeclimate,quiet,pep8}",
		Sources: cli.EnvVars("PARAMETER_LINT_FORMAT", "LINT_FORMAT"),
	},

	&cli.BoolFlag{
		Name:    "lint.quieter",
		Usage:   "quieter, although not silent output",
		Sources: cli.EnvVars("PARAMETER_LINT_QUIETER", "LINT_QUIETER"),
	},

	&cli.BoolFlag{
		Name:    "lint.parseable",
		Usage:   "parseable output in the format of pep8",
		Sources: cli.EnvVars("PARAMETER_LINT_PARSEABLE", "LINT_PARSEABLE"),
	},

	&cli.BoolFlag{
		Name:    "lint.parseable.severity",
		Usage:   "parseable output including severity of rule",
		Sources: cli.EnvVars("PARAMETER_LINT_PARSEABLESEVERITY", "LINT_PARSEABLESEVERITY"),
	},

	&cli.BoolFlag{
		Name: "lint.progressive",
		Usage: " Return success if it detects a reduction in number of " +
			"violations compared with previous git commit. This " +
			"feature works only in git repositories.",
		Sources: cli.EnvVars("PARAMETER_LINT_PROGRESSIVE", "LINT_PROGRESSIVE"),
	},

	&cli.StringFlag{
		Name: "lint.project.dir",
		Usage: " Location of project/repository, autodetected based on " +
			"location of configuration file.",
		Sources: cli.EnvVars("PARAMETER_LINT_PROJECTDIR", "LINT_PROJECTDIR"),
	},

	&cli.StringFlag{
		Name: "lint.rules",
		Usage: "specify one or more rules directories using one or more -r arguments. " +
			"Any -r flags overridethe default rules in /path/to/ansible-lint/lib/ansiblelint/rules," +
			" unless -R is also used.",
		Sources: cli.EnvVars("PARAMETER_LINT_RULE", "LINT_RULE"),
	},

	&cli.BoolFlag{
		Name: "lint.rules.default",
		Usage: "use default rules in /path/to/ansible-lint/lib/ansiblelint/rules in addition" +
			" to any extra rules directories specified with -r. There is no need to specify this if no" +
			" -r flags are used",
		Sources: cli.EnvVars("PARAMETER_LINT_RULESDEFAULT", "LINT_RULESDEFAULT"),
	},

	&cli.BoolFlag{
		Name:    "lint.show.relative.path",
		Usage:   "display path relative to CWD",
		Sources: cli.EnvVars("PARAMETER_LINT_SHOWRELATIVEPATH", "LINT_SHOWRELATIVEPATH"),
	},

	&cli.StringSliceFlag{
		Name:    "lint.tags",
		Usage:   "only check rules whose id/tags match these values",
		Sources: cli.EnvVars("PARAMETER_LINT_TAGS", "LINT_TAGS"),
	},

	&cli.BoolFlag{
		Name:    "lint.tags.list",
		Usage:   "list all the tags",
		Sources: cli.EnvVars("PARAMETER_LINT_TAGSLIST", "LINT_TAGSLIST"),
	},

	&cli.BoolFlag{
		Name:    "lint.verbose",
		Usage:   "increase verbosity level",
		Sources: cli.EnvVars("PARAMETER_LINT_VERBOSE", "LINT_VERBOSE"),
	},

	&cli.StringSliceFlag{
		Name:    "lint.skip",
		Usage:   "only check rules whose id/tags do not match these values",
		Sources: cli.EnvVars("PARAMETER_LINT_SKIP", "LINT_SKIP"),
	},

	&cli.StringSliceFlag{
		Name: "lint.warn",
		Usage: "only warn about these rules, unless overridden in " +
			"config file defaults to 'experimental'",
		Sources: cli.EnvVars("PARAMETER_LINT_WARN", "LINT_WARN"),
	},

	&cli.StringSliceFlag{
		Name:    "lint.enable",
		Usage:   "activate optional rules by their tag name",
		Sources: cli.EnvVars("PARAMETER_LINT_ENABLE", "LINT_ENABLE"),
	},

	&cli.BoolFlag{
		Name:    "lint.no.color",
		Usage:   "disable colored output",
		Sources: cli.EnvVars("PARAMETER_LINT_NOCOLOR", "LINT_NOCOLOR"),
	},

	&cli.BoolFlag{
		Name:    "lint.force.color",
		Usage:   "try force colored output (relying on ansible's code)",
		Sources: cli.EnvVars("PARAMETER_LINT_FORCECOLOR", "LINT_FORCECOLOR"),
	},

	&cli.StringSliceFlag{
		Name:    "lint.exclude",
		Usage:   "path to directories or files to skip. this option is repeatable",
		Sources: cli.EnvVars("PARAMETER_LINT_EXCLUDE", "LINT_EXCLUDE"),
	},

	&cli.StringFlag{
		Name:    "lint.config",
		Usage:   "specify configuration file to use. defaults to \".ansible-lint\"",
		Sources: cli.EnvVars("PARAMETER_LINT_CONFIG", "LINT_CONFIG"),
	},

	&cli.BoolFlag{
		Name:    "lint.offline",
		Usage:   "Disable installation of requirements.yml",
		Sources: cli.EnvVars("PARAMETER_LINT_OFFLINE", "LINT_OFFLINE"),
	},
}
