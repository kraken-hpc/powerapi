# PowerAPI - A REST-ful API for cluster power control

This repository contains several distinct but related things:

***API Specification***:
- **PowerAPI** -- A proposed [RedFish](https://redfish.dmtf.org)-inspired API for power control of distributed systems.  PowerAPI is [OpenAPI 3.0](https://www.openapis.org/) compliant.  It differs with current RedFish standards in at least two important ways:

   1) It allows for a light-weight API without the overhead of the full [RedFish schema set](https://redfish.dmtf.org/schemas/v1/).
   2) It provides easier ways to query/set the states of many nodes with a single API call.

***Golang packages***
- **pkg/powerman** -- The package `pkg/powerman` provides golang bindings for [libpowerman](https://github.com/chaos/powerman/tree/master/libpowerman), allowing Go programs to directly call the API for the popular [powerman](https://github.com/chaos/powerman) cluster power control software.

- **pkg/powerapi-client** -- Provides a Go client API for using the OpenAPI 3 PowerAPI specification.

***Utilities***
- **cmd/powermanapi** -- Provides gateway between Powerman and PowerAPI, creating a REST-ful interface for Powerman cluster node power control.  This utility is intended to provide an easily reached API for cluster management and automation systems, like [Kraken](https://github.com/hpc/kraken)
- **cmd/pmshell** -- Provides a CLI for interacting with Powerman through the `pkg/powerman` Go bindings.  This exists largely as a way to test the `pkg/powerman` <-> `libpowerman` bindings, but may prove useful as a utility.
- **cmd/powerctl** -- Provides a CLI (very similar  to `cmd/pmshell` in interface) to control power through the PowerAPI.  This can be used,  e.g. in conjunction with `cmd/powermanapi` to use and test REST-ful API control of node power states.


More details on each of these can be found in their respective directories.

The PowerAPI specification can be found at [openapi.yaml](openapi.yaml).

## Authors

This software was written by J. Lowell Wofford <lowell@lanl.gov>