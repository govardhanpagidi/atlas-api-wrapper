#!/usr/bin/env bash
set -o errexit
set -o nounset


FOLDER="./api/mmc_atlas_api_client"
OPENAPI_FOLDER="./openapi"
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-swagger.yaml}
transformed_file="atlas-api-transformed.yaml"
client_package="mmc_atlas_api_client"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"

echo "# Running Test Generation"

npm exec openapi-generator-cli -- generate \
    -c "./config/config.yaml" -i "$openapiFileLocation" -o "$FOLDER" \
    --package-name="$client_package" \
    --type-mappings=integer=int \
    --ignore-file-override=config/.go-ignore-tests

