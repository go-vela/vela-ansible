// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"

	"github.com/go-vela/vela-ansible/action/lint"
	"github.com/go-vela/vela-ansible/action/playbook"
	"github.com/go-vela/vela-ansible/version"
)

func main() {
	// capture application version information.
	pluginVersion := version.New()

	// serialize the version information as pretty JSON
	bytes, err := json.MarshalIndent(pluginVersion, "", "  ")
	if err != nil {
		logrus.Fatal(err)
	}

	// output the version information to stdout
	fmt.Fprintf(os.Stdout, "%s\n", string(bytes))

	app := &cli.Command{
		Name:      "vela-ansible",
		Usage:     "Vela Ansible Plugin for running ansible-playbook and ansible-lint.",
		Copyright: "Copyright 2022 Target Brands, Inc. All rights reserved.",
		Authors: []any{
			"Vela Admins <vela@target.com>",
		},
		Action:  run,
		Version: pluginVersion.Semantic(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "log.level",
				Usage:   "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
				Value:   "info",
				Sources: cli.EnvVars("PARAMETER_LOG_LEVEL", "VELA_LOG_LEVEL"),
			},

			&cli.StringFlag{
				Name:    "action",
				Usage:   "set plugin action - options: (lint|playbook)",
				Value:   "lint",
				Sources: cli.EnvVars("PARAMETER_ACTION", "ANSIBLE_ACTION"),
			},
		},
	}

	// ansible-lint flags
	app.Flags = append(app.Flags, lint.Flags...)

	// ansible-playbook flags
	app.Flags = append(app.Flags, playbook.Flags...)

	if err := app.Run(context.Background(), os.Args); err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(_ context.Context, cmd *cli.Command) error {
	// setting the log level
	switch cmd.String("log.level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "i", "info", "Info", "INFO":
		fallthrough
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.WithFields(logrus.Fields{
		"code":     "https://github.com/go-vela/vela-ansible",
		"docs":     "https://go-vela.github.io/docs/plugins/registry/pipeline/ansible/",
		"registry": "https://hub.docker.com/r/target/vela-ansible",
	}).Info("Vela Ansible Plugin")

	// create plugin
	p := &Plugin{
		action: cmd.String("action"),
	}

	// configure plugin depending on action
	switch strings.ToLower(p.action) {
	case AnsibleLint:
		p.Lint = lint.Config(cmd)
	case AnsiblePlaybook:
		p.Playbook = playbook.Config(cmd)
	}

	//validate configuration
	if err := p.Validate(); err != nil {
		return err
	}

	// execute the plugin
	return p.Exec()
}
