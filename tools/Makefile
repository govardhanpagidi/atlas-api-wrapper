export OPENAPI_FOLDER=../openapi
export SDK_FOLDER=../mmc
export DOC_FOLDER=../docs

.DEFAULT_GOAL := clean_and_generate

.PHONY: clean_client
clean_client:
	rm -Rf ../mmc_atlas_api_client

.PHONY: generate_client
generate_client:
	./scripts/generate.sh

.PHONY: clean_and_generate
clean_and_generate: clean_client generate_client

.PHONY: generate_tests
generate_tests:	
	rm -Rf ../test/generated/test*
	./scripts/generate_tests.sh
