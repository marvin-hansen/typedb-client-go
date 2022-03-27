# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# "---------------------------------------------------------"
# "-                                                       -"
# "-  Tests for all dependencies required by make          -"
# "-                                                       -"
# "---------------------------------------------------------"

command bash --version >/dev/null 2>&1 || {
  # command sudo apt-get -qqq -y install curl
   echo "Please install bash"
   exit
}
echo "* Bash installed"

command clang --version >/dev/null 2>&1 || {
  # command sudo apt-get -qqq -y install curl
   echo "Please install clang"
   exit
}
echo "* Clang installed"

command bazel version >/dev/null 2>&1 || {
      echo "Please install Bazel"
      echo "https://docs.bazel.build/versions/main/install.html"
      exit
}
echo "* Bazel installed"

command go version >/dev/null 2>&1 || {
   echo "Please install Golang"
       echo "https://go.dev/doc/install"
       exit
}
echo "* Golang installed"

echo ""
echo "==============================="
echo "All dependencies installed."
echo "==============================="
echo ""