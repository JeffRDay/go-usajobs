#!/bin/bash

read -p "input name: " NAME

cp cli/cmd/academichonors.go cli/cmd/${NAME}.go
cp cli/cmd/academichonors_test.go cli/cmd/${NAME}_test.go
cp examples/cli/academichonors.sh examples/cli/${NAME}.sh
chmod +x examples/cli/${NAME}.sh
