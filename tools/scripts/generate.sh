#!/usr/bin/env bash
set -o errexit
set -o nounset

######################################################
# Generate client using OpenAPI generator
# Environment variables:
#   OPENAPI_FOLDER - folder containing openapi file
#   OPENAPI_FILE_NAME - openapi file name
#   SDK_FOLDER - folder location for generated client
######################################################

OPENAPI_FOLDER="./openapi"
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-swagger.yaml}
SDK_FOLDER="./api/mmc_atlas_api_client"

transformed_file="atlas-api-transformed.yaml"
client_package="mmc_atlas_api_client"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"
pwd
echo "# Running generation pipeline"
echo "# Running transformation based on $OPENAPI_FILE_NAME to the $transformed_file"
echo "# Generated client to  $OPENAPI_FOLDER and $SDK_FOLDER"
cp "$OPENAPI_FOLDER/$OPENAPI_FILE_NAME" "$openapiFileLocation"

npm install
#npm run sdk:transform -- "$openapiFileLocation"

echo "# Running OpenAPI generator validation"
npm exec openapi-generator-cli -- generate -i "$OPENAPI_FOLDER/$OPENAPI_FILE_NAME"  -g openapi -o ./openapi 
input_file="./openapi/openapi.json"

# Temporary file for storing the modified JSON
temp_file="./openapi/openapi_temp.json"
src_file="./openapi/openapi_transformed.json"
# Perform the replacement using jq

# Replace the original file with the modified file
jq 'walk(if type == "object" and has("atlasresponse.AtlasResponse") then .AtlasResponse = .["atlasresponse.AtlasResponse"] | del(.["atlasresponse.AtlasResponse"]) else . end)' "$input_file" > "$temp_file"
#cp  "$temp_file" "$input_file"
jq "walk(if type == \"string\" and test(\"atlasresponse.AtlasResponse\") then gsub(\"atlasresponse.AtlasResponse\"; \"AtlasResponse\") else . end)" "$temp_file" > "$src_file"
cp  "$temp_file" "$input_file"
echo "# Running Client Generation"

echo "# Running OpenAPI converter validation"
npm exec openapi-generator-cli -- generate \
    -c "./config/config.yaml" -i "$src_file" -o "$SDK_FOLDER" \
    --package-name="$client_package" \
    --type-mappings=integer=int \
    --ignore-file-override=config/.go-ignore
#mkdir "$SDK_FOLDER/digest/"
#cp ./digest/digest.go "$SDK_FOLDER/digest/digest.go"

gofmt -s -w "$SDK_FOLDER/"*.go
