#!/bin/sh

# Check if go.mod exists, if not, initialize the module and tidy up dependencies
if [ ! -f go.mod ]; then
    go mod init load-tester
    go mod tidy
fi
