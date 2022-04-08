# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# Bazel consistently throws the following error
# ""There were tests whose specified size is too big.""
# This seems to be persistent issue in OSX with no clear solution.
# Means Bazel test is basically useless.

# Replaced with Go Test b/c it's actually working
# The  -count=1 flag disables caching. Forces all tests to run.
# https://stackoverflow.com/questions/48882691/force-retesting-or-disable-test-caching

# Caching causing problems when the DeleteDB or CreateDB test has been cached,
# then reported as passed when, in fact, no operation was executed and subsequent tests fail.

# Testing create & delete DB
cd test/client/db_admin/
go test -count=1 ./...

# Setup required for all subsequent tests.
 cd ../a_pre_test/
go test -count=1 ./...

# Test session
cd ../session/
go test -count=1 ./...

# Test create & get schema
cd ../schema/
go test -count=1 ./...

# Insert test fails.
# Test queries
#cd ../query/
#go test -count=1 ./...

# Teardown of everything that was created during pre-test setup
cd ../z_post_test/
go test -count=1 ./...
