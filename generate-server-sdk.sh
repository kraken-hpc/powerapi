#!/bin/bash

podman run -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go-server -o /local/cmd/powermanapi -c /local/openapi-generator-server.yaml
