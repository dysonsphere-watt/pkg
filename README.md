# Package Utility for DSAdmin Backend

Currently helps with Consul and Hertz setup.

## Private Repository Setup

If the module is hosted in a private repository, running `go mod tidy` will result in a 403 error. Do bypass this, configure git to use SSH for either this repository or for all projects under the organisation.

```bash
# From anywhere
go env -w GOPRIVATE=github.com/dysonsphere-watt
git config --global url."git@github.com:dysonsphere-watt".insteadOf "https://github.com/dysonsphere-watt"

# From backend modules or anywhere that imports this pkg module
go mod tidy

# Outputs:
# go: finding module for package github.com/dysonsphere-watt/pkg
# go: downloading github.com/dysonsphere-watt/pkg v0.0.0-20240424064517-a8624bdeb150
# go: found github.com/dysonsphere-watt/pkg in github.com/dysonsphere-watt/pkg v0.0.0-20240424064517-a8624bdeb150
```
