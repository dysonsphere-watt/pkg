#!/bin/bash
# Runs go linter on all files. Run it in a UNIX environment

# [golangci-lint]
# Folow instructions to install golangci-lint here: https://golangci-lint.run/welcome/install/#local-installation
# If on windows, install golangci-lint in Git Bash and run it from there.
# If on UNIX, install locally then run.

if ! golangci-lint run -c .golangci.yml; 
then
    printf "[golangci-lint]\tLint errors found\n"
else
    printf "[golangci-lint]\tNo Lint errors found\n"
fi