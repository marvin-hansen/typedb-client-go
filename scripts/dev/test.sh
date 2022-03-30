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
# Caching causing problems when the DeleteDB test has been cached,
# reported as passed when in fact no delete operation was executed.

# Testing create & delete DB
cd test/client/db_admin/
go test -count=1 ./...

# Setup required for all subsequent tests.
 cd ../a_pre_test/
go test -count=1 ./...

# Test create & get schema
cd ../schema/
go test -count=1 ./...

# Test queries
cd ../query/
go test -count=1 ./...

# Teardown of everything that was created during pre-test setup
cd ../z_post_test/
go test -count=1 ./...
