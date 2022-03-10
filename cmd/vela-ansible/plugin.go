// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// plugin - action controller of the plugin.

package main

import (
	"errors"
	"strings"

	"github.com/go-vela/vela-ansible/action/lint"
	"github.com/go-vela/vela-ansible/action/playbook"
	"github.com/sirupsen/logrus"
)

// action constants.
const (
	AnsibleLint     = "lint"
	AnsiblePlaybook = "playbook"
)

var (
	// ErrorInvalidAction is returned when the user provided an invalid action.
	ErrorInvalidAction = errors.New("invalid action, please choose lint or playbook")
)

type (
	Plugin struct {
		// action argument loaded for the plugin
		action string
		// Playbook arguments loaded for the plugin
		Playbook *playbook.Playbook
		// Lint arguments loaded for the plugin
		Lint *lint.Linter
	}
)

// Validate plugin based on action.
func (p *Plugin) Validate() error {
	logrus.Trace("entered plugin.Validate")
	defer logrus.Trace("exited plugin.Validate")

	switch strings.ToLower(p.action) {
	case AnsibleLint:
		if err := lint.Validate(p.Lint); err != nil {
			return err
		}
	case AnsiblePlaybook:
		if err := playbook.Validate(p.Playbook); err != nil {
			return err
		}
	default:
		return ErrorInvalidAction
	}

	return nil
}

// Exec configs and runs ansible plugin.
func (p *Plugin) Exec() error {
	logrus.Trace("entered plugin.Exec")
	defer logrus.Trace("exited plugin.Exec")

	logrus.Info("start: ansible plugin")

	switch strings.ToLower(p.action) {
	case AnsibleLint:
		logrus.Info("ansible-lint")

		if err := lint.Exec(p.Lint); err != nil {
			return err
		}
	case AnsiblePlaybook:
		logrus.Info("ansible-playbook")

		if err := playbook.Exec(p.Playbook); err != nil {
			return err
		}
	default:
		return ErrorInvalidAction
	}

	return nil
}
