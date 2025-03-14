//go:build tools
// +build tools

package tools

import (
	_ "github.com/air-verse/air"
	_ "github.com/evanw/esbuild/cmd/esbuild"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)
