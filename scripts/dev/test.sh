# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# Note, Bazel consistently throws
# ""There were tests whose specified size is too big.""
# This seems to be persistent issue in OSX with no clear solution.
# Means Bazel test is basically useless.

# Replaced with Go Test b/c it's actually working

#
cd test/client/db-admin
go test ./...

cd ../schema
go test ./...
