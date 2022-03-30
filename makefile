# Make will use bash instead of sh
SHELL := /usr/bin/env bash
CC=clang # required by bazel
#ENV=LOCAL # Required by auto-config.

# GNU make man page
# http://www.gnu.org/software/make/manual/make.html

# For some strange reasons, intends & blanks shift in bash when calling 'make' so the formatting below should align intend at least on Bash on OSX.
.PHONY: help
help:
	@echo ' '
	@echo 'Setup: '
	@echo '    make check        	    	  Checks whether all project requirements are present.'
	@echo ' '
	@echo 'Dev: '
	@echo '    make build   		Builds the code base incrementally (fast). Use for coding.'
	@echo '    make rebuild   		Rebuilds dependencies, build files, & code base (slow). Use after go mod changes.  '
	@echo '    make tests        		Runs all tests. See scripts/dev/test.sh for details.'
	@echo '    make stats        		Shows the latest project stats.'

# "---------------------------------------------------------"
# Setup
# "---------------------------------------------------------"
.PHONY: check
check:
	@source scripts/setup/check_requirements.sh

# "---------------------------------------------------------"
# Development
# "---------------------------------------------------------"
.PHONY: build
build:
	@source scripts/dev/build.sh

.PHONY: rebuild
rebuild:
	@source scripts/dev/rebuild.sh


.PHONY: test
test:
	@source scripts/dev/test.sh

.PHONY: stats
stats:
	@source scripts/dev/stats.sh
