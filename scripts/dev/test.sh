# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# note, Bazel consistently throws
# ""There were tests whose specified size is too big.""
# So the reported test results are basically useless.

# Needs to be replaced with Go Test

# Build all test sources
bazel build //test:build

# run DB admin tests first i.e. create & delete DB
bazel test //test/client/db_admin/... --test_output=streamed --nocache_test_results

# Run schema tests next
bazel test //test/client/schema/... --test_output=streamed --nocache_test_results

# Run all query tests
bazel test //test/client/query/... --test_output=streamed --nocache_test_results