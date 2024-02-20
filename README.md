# 🔨 Htmx

[![Test & Build](https://github.com/zeiss/fiber-htmx/actions/workflows/main.yml/badge.svg)](https://github.com/zeiss/fiber-htmx/actions/workflows/main.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/zeiss/fiber-htmx.svg)](https://pkg.go.dev/github.com/zeiss/fiber-htmx)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeiss/fiber-htmx)](https://goreportcard.com/report/github.com/zeiss/fiber-htmx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)



## Installation

```bash
$ go get github.com/zeiss/fiber-htmx
```

## Components

- [x] [htmx](https://htmx.org/)
- [x] [HTML5](https://www.w3.org/TR/2011/WD-html5-20110405/)
- [ ] [DaisyUI](https://daisyui.com/) (WIP)

The package supports to write HTML5 components and HTMX components in Go.

```go
htmx.Button(
    htmx.Attribute("type", "submit")
    htmx.Text("Button"),
    htmx.HxPost("/api/respond")
)
```

## Examples

See [examples](https://github.com/zeiss/fiber-htmx/tree/master/examples) to understand the provided interfaces.

## License

[MIT](/LICENSE)