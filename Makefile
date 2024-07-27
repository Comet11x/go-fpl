SHELL              = bash
TARGET             = ./target
TARGET_RELEASE     = $(TARGET)/release

GO_FPL_NAME        = fpl
GO_FPL_SRC         = ./cmd/$(GO_FPL_NAME)/main.go
GO_FPL_BIN         = $(TARGET_RELEASE)/$(GO_FPL_NAME)

GO_FPL_CORE_NAME   = fpl_core 
GO_FPL_CORE_SRC    = ./examples/core/main.go
GO_FPL_CORE_BIN    = $(TARGET_RELEASE)/$(GO_FPL_CORE_NAME)

TEST_COVERAGE_DIR  = test_coverage
TEST_COVERAGE_FILE = test_coverage
CORE_TEST_COVERAGE = $(TEST_COVERAGE_DIR)/core/$(TEST_COVERAGE_FILE)


add_dependencies:
	@echo "This project does not have any dependencies"

make_target:
	@mkdir -p $(TARGET_RELEASE)

make_test_coverage_dirs:
	@mkdir -p $(TEST_COVERAGE_DIR)/core

init_githooks_hooks:
	@cp ./githooks/pre-commit.sh .git/hooks/pre-commit

run_tests:
	@go test ./test/... -v

measure_test_coverage: make_test_coverage_dirs
	go test -v ./pkg/core/... -coverprofile=$(CORE_TEST_COVERAGE)
	go tool cover -func=$(CORE_TEST_COVERAGE)

show_core_test_coverage: measure_test_coverage
	go tool cover -html=$(CORE_TEST_COVERAGE)

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
