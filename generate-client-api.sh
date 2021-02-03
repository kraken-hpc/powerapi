#!/bin/bash

podman run -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go -o /local/pkg/powerapi-client -c /local/openapi-generator-client.yaml