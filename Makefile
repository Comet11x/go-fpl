TARGET           = ./target
TARGET_RELEASE   = $(TARGET)/release

GO_FPL_NAME      = fpl
GO_FPL_SRC       = ./cmd/$(GO_FPL_NAME)/main.go
GO_FPL_BIN       = $(TARGET_RELEASE)/$(GO_FPL_NAME)

GO_FPL_CORE_NAME = fpl_core 
GO_FPL_CORE_SRC  = ./examples/core/main.go
GO_FPL_CORE_BIN  = $(TARGET_RELEASE)/$(GO_FPL_CORE_NAME)

add_dependencies:
	@echo "This project does not have any dependencies"

make_target:
	@mkdir -p $(TARGET_RELEASE)

init_githooks_hooks:
	@cp ./githooks/pre-commit.sh .git/hooks/pre-commit

run_tests:
	@go test ./test/... -v

build: make_target init_githooks_hooks run_tests
	go build $(GO_FPL_SRC)
	mv main $(GO_FPL_BIN)

build_example_core: make_target init_githooks_hooks run_tests
	go build $(GO_FPL_CORE_SRC)
	mv main $(GO_FPL_CORE_BIN)

run: build
	$(GO_FPL_BIN)

run_examples: build_example_core
	$(GO_FPL_CORE_BIN)
