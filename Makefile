# SPDX-License-Identifier: Apache-2.0

# capture the current date we build the application from
BUILD_DATE = $(shell date +%Y-%m-%dT%H:%M:%SZ)

# check if a git commit sha is already set
ifndef GITHUB_SHA
	# capture the current git commit sha we build the application from
	GITHUB_SHA = $(shell git rev-parse HEAD)
endif

# check if a git tag is already set
ifndef GITHUB_TAG
	# capture the current git tag we build the application from
	GITHUB_TAG = $(shell git describe --tag --abbrev=0)
endif

# check if a go version is already set
ifndef GOLANG_VERSION
	# capture the current go version we build the application from
	GOLANG_VERSION = $(shell go version | awk '{ print $$3 }')
endif

# create a list of linker flags for building the golang application
LD_FLAGS = -X github.com/go-vela/vela-ansible/version.Commit=${GITHUB_SHA} -X github.com/go-vela/vela-ansible/version.Date=${BUILD_DATE} -X github.com/go-vela/vela-ansible/version.Go=${GOLANG_VERSION} -X github.com/go-vela/vela-ansible/version.Tag=${GITHUB_TAG}

# The `clean` target is intended to clean the workspace
# and prepare the local changes for submission.
#
# Usage: `make clean`
clean: tidy vet fmt fix

# The `run` target is intended to build and
# execute the Docker image for the plugin.
#
# Usage: `make run`
run: build docker-build docker-run

# The `tidy` target is intended to clean up
# the Go module files (go.mod & go.sum).
#
# Usage: `make tidy`
tidy:
	@echo
	@echo "### Tidying Go module"
	@go mod tidy

# The `vet` target is intended to inspect the
# Go source code for potential issues.
#
# Usage: `make vet`
vet:
	@echo
	@echo "### Vetting Go code"
	@go vet ./...

# The `fmt` target is intended to format the
# Go source code to meet the language standards.
#
# Usage: `make fmt`
fmt:
	@echo
	@echo "### Formatting Go Code"
	@go fmt ./...

# The `fix` target is intended to rewrite the
# Go source code using old APIs.
#
# Usage: `make fix`
fix:
	@echo
	@echo "### Fixing Go Code"
	@go fix ./...

# The `test` target is intended to run
# the tests for the Go source code.
#
# Usage: `make test`
test:
	@echo
	@echo "### Testing Go Code"
	@go test -race ./...

# The `test-cover` target is intended to run
# the tests for the Go source code and then
# open the test coverage report.
#
# Usage: `make test-cover`
test-cover:
	@echo
	@echo "### Creating test coverage report"
	@go test -race -covermode=atomic -coverprofile=coverage.out ./...
	@echo
	@echo "### Opening test coverage report"
	@go tool cover -html=coverage.out

# The `build` target is intended to compile
# the Go source code into a binary.
#
# Usage: `make build`
build:
	@echo
	@echo "### Building release/vela-ansible binary"
	GOOS=linux CGO_ENABLED=0 \
		go build -a \
		-ldflags '${LD_FLAGS}' \
		-o release/vela-ansible \
		github.com/go-vela/vela-ansible/cmd/vela-ansible

# The `build-static` target is intended to compile
# the Go source code into a statically linked binary.
#
# Usage: `make build-static`
build-static:
	@echo
	@echo "### Building static release/vela-ansible binary"
	GOOS=linux CGO_ENABLED=0 \
		go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/vela-ansible \
		github.com/go-vela/vela-ansible/cmd/vela-ansible

# The `build-static-ci` target is intended to compile
# the Go source code into a statically linked binary
# when used within a CI environment.
#
# Usage: `make build-static-ci`
build-static-ci:
	@echo
	@echo "### Building CI static release/vela-ansible binary"
	@go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/vela-ansible \
		github.com/go-vela/vela-ansible/cmd/vela-ansible

# The `check` target is intended to output all
# dependencies from the Go module that need updates.
#
# Usage: `make check`
check: check-install
	@echo
	@echo "### Checking dependencies for updates"
	@go list -u -m -json all | go-mod-outdated -update

# The `check-direct` target is intended to output direct
# dependencies from the Go module that need updates.
#
# Usage: `make check-direct`
check-direct: check-install
	@echo
	@echo "### Checking direct dependencies for updates"
	@go list -u -m -json all | go-mod-outdated -direct

# The `check-full` target is intended to output
# all dependencies from the Go module.
#
# Usage: `make check-full`
check-full: check-install
	@echo
	@echo "### Checking all dependencies for updates"
	@go list -u -m -json all | go-mod-outdated

# The `check-install` target is intended to download
# the tool used to check dependencies from the Go module.
#
# Usage: `make check-install`
check-install:
	@echo
	@echo "### Installing psampaz/go-mod-outdated"
	@go get -u github.com/psampaz/go-mod-outdated

# The `bump-deps` target is intended to upgrade
# non-test dependencies for the Go module.
#
# Usage: `make bump-deps`
bump-deps: check
	@echo
	@echo "### Upgrading dependencies"
	@go get -u ./...

# The `bump-deps-full` target is intended to upgrade
# all dependencies for the Go module.
#
# Usage: `make bump-deps-full`
bump-deps-full: check
	@echo
	@echo "### Upgrading all dependencies"
	@go get -t -u ./...

# The `docker-build` target is intended to build
# the Docker image for the plugin.
#
# Usage: `make docker-build`
docker-build:
	@echo
	@echo "### Building vela-ansible:local image"
	@docker build --no-cache -t vela-ansible:local .

# The `docker-test` target is intended to execute
# the Docker image for the plugin with test variables.
#
# Usage: `make docker-test`
docker-test:
	@echo
	@echo "### Testing vela-ansible:local image"
	@docker run --rm \
	  -e PARAMETER_LINT_VERSION=true \
	  -e PARAMETER_PLAYBOOK="playbook.yml" \
      -e PARAMETER_OPTIONS_INVENTORY="hosts.yml" \
      -e PARAMETER_CONNECTION_USER="root" \
      -e PARAMETER_CONNECTION_PRIVATEKEY="id_rsa" \
      vela-ansible:local

# The `docker-run` target is intended to execute
# the Docker image for the plugin.
#
# Usage: `make docker-run`
docker-run:
	@echo
	@echo "### Executing vela-ansible:local image"
	@docker run --rm \
	  -e PARAMETER_PLAYBOOK="" \
      -e PARAMETER_OPTIONS_INVENTORY="" \
      -e PARAMETER_CONNECTION_USER=""\
      ansible:local
