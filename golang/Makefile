BUILD_DIR := ./build
DOCKER    := $(shell which docker)

# Enable latest docker features
export DOCKER_CLI_EXPERIMENTAL  := enabled
export DOCKER_BUILDKIT          := 1
export COMPOSE_DOCKER_CLI_BUILD := 1
export BUILDKIT_PROGRESS        := plain

.PHONY: build
build:
	$(DOCKER) build -t golang-example:latest .

.PHONY: clean
clean:
	@echo "Removing all go download and build caches"
	$(DOCKER) builder prune --filter type=exec.cachemount

	@echo "Removing all artifacts"
	rm -rf $(BUILD_DIR)