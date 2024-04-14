APP_NAME=booking
BIN=./booking

# We always use go in module mode
export GO111MODULE=on
GO_EXEC=go
GOFMT = gofmt

# Detecting OS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    OS = LINUX
    REPLACE = sed -i"" -e
endif
ifeq ($(UNAME_S),Darwin)
    OS = OSX
    REPLACE = sed -i "" -e
endif

.PHONY: start
start:
	echo "Starting the app..."
	@bash run.sh


.PHONY: deps
deps: ; $(info $(M) install depends…) @ ## Install depends
	$(info #Install dependencies...)
	$(GO_EXEC) mod tidy

.PHONY: build
build: deps ; $(info $(M) run build…) @ ## Run build for deploy
	$(info #Building app...)
	$(BUILD_ENVPARMS) $(GO_EXEC) build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/$(APP_NAME)

.PHONY: run
run: build ; $(info $(M) running the app…) @ ## Run the app
	$(info #Running app...)
	$(BIN)

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ": .*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
