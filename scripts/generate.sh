#!/bin/bash

curl -o configs/keycloak.yaml https://www.keycloak.org/docs-api/latest/rest-api/openapi.yaml

docker run --rm --user $(id -u):$(id -g) -v $(pwd):/local mikefarah/yq e 'del(.. | select(.deprecated? == true))' /local/configs/keycloak.yaml > configs/keycloak_clean.yaml

docker run --rm --user $(id -u):$(id -g) -v $(pwd):/local openapitools/openapi-generator-cli generate \
  --skip-validate-spec -i /local/configs/keycloak_clean.yaml -g go -o /local \
  --git-repo-id=keycloakadminclient \
  --git-user-id=miguoliang \
  --additional-properties=structPrefix=true,enumClassPrefix=true,packageName=keycloakadminclient,packageVersion=24.0.2
