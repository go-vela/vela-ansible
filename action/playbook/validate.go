// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// validate - validates ansible-playbook to make sure that a playbook and inventory is defined.

package playbook

import (
	"errors"
	"net"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// ErrorEmptyPlaybook is returned when the playbook provided is empty.
	ErrorEmptyPlaybook = errors.New("ansible-playbook: playbook is empty")

	// ErrorInvalidPlaybook is returned when the playbook path provided is invalid.
	ErrorInvalidPlaybook = errors.New("ansible-playbook: playbook invalid path")

	// ErrorMissingPlaybook is returned when the there is no playbook provided.
	ErrorMissingPlaybook = errors.New("ansible-playbook: playbook not specified")

	// ErrorEmptyInventory is returned when the inventory provided is empty.
	ErrorEmptyInventory = errors.New("ansible-playbook: inventory is empty")

	// ErrorInvalidInventory is returned when the inventory path provided is invalid.
	ErrorInvalidInventory = errors.New("ansible-playbook: inventory invalid path or ip address")

	// ErrorMissingInventory is returned when there is no inventory provided.
	ErrorMissingInventory = errors.New("ansible-playbook: inventory not specified")
)

// Validate ansible-playbook.
func Validate(p *Playbook) error {
	logrus.Trace("entered plugin.playbook.Validate")
	defer logrus.Trace("exited plugin.playbook.Validate")

	logrus.Debug("validating ansible-playbook configuration")

	if p.Options.Version {
		return nil
	}

	// validate playbook
	logrus.Info("validating ansible-playbook")

	if len(p.Playbook) != 0 {
		file, err := os.Stat(p.Playbook)
		if err != nil {
			return ErrorInvalidPlaybook
		}

		if file.Size() == 0 {
			return ErrorEmptyPlaybook
		}
	} else {
		return ErrorMissingPlaybook
	}

	// validate inventory
	if len(p.Options.Inventory) != 0 {
		for _, host := range p.Options.Inventory {
			if net.ParseIP(host) == nil {
				file, err := os.Stat(host)
				if err != nil {
					return ErrorInvalidInventory
				}

				if file.Size() == 0 {
					return ErrorEmptyInventory
				}
			} else {
				continue
			}
		}
	} else {
		return ErrorMissingInventory
	}

	return nil
}
