#!/bin/bash

set -e

type go-licenses &> /dev/null || go get github.com/google/go-licenses

MODULES+=("github.com/jraman567/corim/corim")
MODULES+=("github.com/jraman567/corim/comid")
MODULES+=("github.com/jraman567/corim/cocli")

for module in ${MODULES[@]}
do
  echo ">> retrieving licenses [ ${module} ]"
  go-licenses csv ${module}
done
