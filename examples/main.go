package main

import (
	"log"
	"log/slog"
	"os"

	fiber_htmx "github.com/zeiss/fiber-htmx/v3"

	"github.com/gofiber/fiber/v3"
	safe "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/spf13/pflag"
	reload "github.com/zeiss/fiber-reload/v3"
	slogx "github.com/zeiss/fiber-slogx"
	htmx "github.com/zeiss/htmx"
	"github.com/zeiss/htmx/buttons"
)

var addr = ":3000"

type GetHelloWorldController struct {
	fiber_htmx.UnimplementedController
}

func (ctrl *GetHelloWorldController) Get(c fiber.Ctx) error {
	return Demo().Render(c)
}

func Demo() htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    "Demo",
			Language: "en",
			Head: []htmx.Node{
				htmx.Script(
					htmx.Src("https://unpkg.com/fiber-reload@0.9.0/reload.js"),
					htmx.Type("text/javascript"),
				),
			},
		},
		htmx.Body(
			buttons.Button(buttons.ButtonProps{}, htmx.Text("Demo")),
		),
	)
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger = logger.With("env", "development")

	pflag.StringVar(&addr, "addr", addr, "address to listen on")

	cfg := slogx.Config{}

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(slogx.New(logger.WithGroup("http"), cfg))
	app.Use(safe.New())
	app.Use(reload.Environment(reload.Development))
	reload.WithHotReload(app)

	ctrl := &GetHelloWorldController{}

	app.Get("/", fiber_htmx.Handle(ctrl))

	err := app.Listen(addr)
	if err != nil {
		log.Fatal(err)
	}
}
