# 🔨 Htmx

[![Test & Build](https://github.com/zeiss/fiber-htmx/actions/workflows/main.yml/badge.svg)](https://github.com/zeiss/fiber-htmx/actions/workflows/main.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/zeiss/fiber-htmx.svg)](https://pkg.go.dev/github.com/zeiss/fiber-htmx)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeiss/fiber-htmx)](https://goreportcard.com/report/github.com/zeiss/fiber-htmx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)

A Go package to write HTML5 and HTMX components in Go. The package is designed to work with [fiber](http://gofiber.io) and [htmx](https://htmx.org/).

## Features

- Write declartive HTML5 components in Go without using templates and with the full-power of a type-safe language, auto-completion, and refactoring.
- Full support for HTMX components.
- No dependencies on JavaScript frameworks.
- Fast rendering of HTML5 and HTMX components.
- Easy to use and learn.
- Easy to extend and customize.

### Example

Creating a button leveraging htmx is as easy as this.

```go
htmx.Button(
    htmx.Attribute("type", "submit")
    htmx.Text("Button"),
    htmx.HxPost("/api/respond")
)
```

## Installation

```bash
$ go get github.com/zeiss/fiber-htmx
```

## Components

Write HTML5 and HTMX components in Go.

```go
func HelloWorld() htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            "font-semibold",
        },
        htmx.Text("Hello World"),
    )
}
```

There are different types of composition. For example, passing children to a component.

```go
func HelloWorld(children ...htmx.Node) htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            "font-semibold",
        },
        htmx.Text("Hello World"),
        htmx.Div(
            htmx.ClassNames{
                "text-red-500",
            },
            htmx.Group(children...),
        ),
    )
}
```
Styling of components is done with the `htmx.ClassNames` type.

```go
func HelloWorld() htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            "font-semibold": true,
            "text-red-500": true,
        },
        htmx.Text("Hello World"),
    )
}
```

There are also helpers to make the life with styling easier by merging classes.

```go
func HelloWorld(classes htmx.ClassNames) htmx.Node {
    return htmx.Div(
        htmx.Merge(
            htmx.ClassNames{
                "font-semibold",
                "text-red-500",
            },
            classes,
        )
        htmx.Text("Hello World"),
    )
}
```

There are additional complex components that help to write HTML5 and HTMX components in Go.

- [x] [htmx](https://htmx.org/)
- [x] [HTML5](https://www.w3.org/TR/2011/WD-html5-20110405/)
- [ ] [DaisyUI](https://daisyui.com/) (WIP)
- [ ] [Heroicons](https://heroicons.com/) (WIP)

There is also the option to use `htmx.Controller` to encapsulate the logic of the components.

```go
type HelloWorldController struct {
    htmx.DefaultController
}

func (c *HelloWorldController) Get() error {
    return htmx.RenderComp(
		c.Ctx(),
		htmx.HTML5(
            htmx.HTML5Props{
                Title:    "index",
                Language: "en",
                Head: []htmx.Node{},
            },
            htmx.Div(
                htmx.ClassNames{},
                htmx.Text("Hello World"),
            ),
        ),
    )    
}

app := fiber.New()
app.Get("/", htmx.NewHxControllerHandler(&HelloWorldController{}))

app.Listen(":3000")
```

## Examples

See [examples](https://github.com/zeiss/fiber-htmx/tree/master/examples) to understand the provided interfaces.

## Benchmarks

```bash
BenchmarkElement-10                     12964440                77.40 ns/op
Benchmark_AttrRender-10                 16038232                74.15 ns/op
Benchmark_HTML5_Render-10                   1392            847193 ns/op
Benchmark_ClassNames_Render-10           3166761               378.2 ns/op
```

Rendering `10.000` nodes took `>0.8ms`. The package is fast enough to render HTML5 and HTMX components. 

## License

[MIT](/LICENSE)
