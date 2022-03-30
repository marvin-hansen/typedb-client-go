# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# Format bazel BUILD and .bzl files with a standard convention.
bazel run //:buildifier

# Build all sources
bazel build //:build