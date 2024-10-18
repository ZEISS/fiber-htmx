.DEFAULT_GOAL := release

GO 					?= go
GO_RUN_TOOLS 		?= $(GO) run -modfile ./tools/go.mod
GO_TEST 			?= $(GO_RUN_TOOLS) gotest.tools/gotestsum --format pkgname
GO_RELEASER 		?= $(GO_RUN_TOOLS) github.com/goreleaser/goreleaser
GO_BENCHSTAT 		?= $(GO_RUN_TOOLS) golang.org/x/perf/cmd/benchstat
GO_MOD 				?= $(shell ${GO} list -m)

.PHONY: release
release: ## Release the project.
	$(GO_RELEASER) release --clean

.PHONY: generate
generate: ## Generate code.
	$(GO) generate ./...

.PHONY: start
start: ## Run air live reload.
	$(GO_RUN_TOOLS) github.com/air-verse/air

.PHONY: bundle
bundle: ## Bundle the project.
	$(GO_RUN_TOOLS) github.com/evanw/esbuild/cmd/esbuild --format=esm --packages=external --outdir=dist src/*.ts
	# $(GO_RUN_TOOLS) github.com/evanw/esbuild/cmd/esbuild --bundle --sourcemap --platform=neutral --packages=external --outfile=dist/fiber-htmx.esm.js src/main.ts
	# $(GO_RUN_TOOLS) github.com/evanw/esbuild/cmd/esbuild --bundle --minify --sourcemap --outfile=dist/fiber-htmx.min.js src/main.ts

.PHONY: fmt
fmt: ## Run go fmt against code.
	$(GO_RUN_TOOLS) mvdan.cc/gofumpt -w .

.PHONY: bench
bench: ## Run benchmarks.
	$(GO) test -bench=. ./...

.PHONY: vet
vet: ## Run go vet against code.
	$(GO) vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	mkdir -p .test/reports
	$(GO_TEST) --junitfile .test/reports/unit-test.xml -- -race ./... -count=1 -short -cover -coverprofile .test/reports/unit-test-coverage.out

.PHONY: lint
lint: ## Run lint.
	$(GO_RUN_TOOLS) github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m -c .golangci.yml

.PHONY: clean
clean: ## Remove previous build.
	rm -rf .test .dist
	find . -type f -name '*.gen.go' -exec rm {} +
	git checkout go.mod

.PHONY: help
help: ## Display this help screen.
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
