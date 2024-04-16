#!/bin/bash

./node_modules/.bin/openapi-generator-cli generate --skip-validate-spec -i keycloak.yaml -g go -o ../ \
  --git-repo-id=keycloakadminclient \
  --git-user-id=miguoliang \
  --additional-properties=structPrefix=true,enumClassPrefix=true,packageName=keycloakadminclient,packageVersion=24.0.2