// SPDX-License-Identifier: Apache-2.0

// config - ansible-lint configuration

package lint

import (
	"github.com/urfave/cli/v2"
)

// apply Linter configuration.
func Config(c *cli.Context) *Linter {
	// lint configuration
	return &Linter{
		LintPlaybook:          c.String("playbook"),
		LintVersion:           c.Bool("lint.version"),
		LintList:              c.Bool("lint.list"),
		LintQuieter:           c.Bool("lint.quieter"),
		LintParseable:         c.Bool("lint.parseable"),
		LintParseableSeverity: c.Bool("lint.parseable.severity"),
		LintRules:             c.StringSlice("lint.rules"),
		LintRulesDefault:      c.Bool("lint.rules.default"),
		LintShowRelativePath:  c.Bool("lint.show.relative.path"),
		LintTags:              c.StringSlice("lint.tags"),
		LintTagsList:          c.Bool("lint.tags.list"),
		LintVerbose:           c.Bool("lint.verbose"),
		LintSkip:              c.StringSlice("lint.skip"),
		LintNoColor:           c.Bool("lint.no.color"),
		LintForceColor:        c.Bool("lint.force.color"),
		LintExclude:           c.StringSlice("lint.exclude"),
		LintConfig:            c.String("lint.config"),
		LintFormat:            c.String("lint.format"),
		LintProgressive:       c.Bool("lint.progressive"),
		LintProjectDir:        c.String("lint.project.dir"),
		LintWarn:              c.StringSlice("lint.warn"),
		LintEnable:            c.StringSlice("lint.enable"),
		LintOffline:           c.Bool("lint.offline"),
	}
}
