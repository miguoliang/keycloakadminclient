#!/bin/bash

docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli generate \
  --skip-validate-spec -i /local/configs/keycloak.yaml -g go -o /local \
  --git-repo-id=keycloakadminclient \
  --git-user-id=miguoliang \
  --additional-properties=structPrefix=true,enumClassPrefix=true,packageName=keycloakadminclient,packageVersion=24.0.2
