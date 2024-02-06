package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/katallaxie/pkg/logger"
	"github.com/spf13/cobra"
	htmx "github.com/zeiss/fiber-htmx"
)

// Config ...
type Config struct {
	Flags *Flags
}

// Flags ...
type Flags struct {
	Addr string
}

var cfg = &Config{
	Flags: &Flags{},
}

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":8080", "addr")

	rootCmd.SilenceUsage = true
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	logger.RedirectStdLog(logger.LogSink)

	app := fiber.New()
	app.Get("/", htmx.NewCompHandler(indexPage))

	app.Post("/api/respond", htmx.NewHtmxHandler(func(hx *htmx.Htmx) error {
		if !hx.IsHxRequest() {
			return nil
		}

		_, err := hx.WriteHTML("<div>New Content</div>")
		return err
	}))

	err := app.Listen(cfg.Flags.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

var indexPage = htmx.HTML5(htmx.HTML5Props{
	Title:    "index",
	Language: "en",
	Head: []htmx.Node{
		htmx.Script(htmx.Attribute("src", "https://unpkg.com/htmx.org@1.9.10"), htmx.Attribute("type", "application/javascript")),
	},
	Body: []htmx.Node{
		htmx.Button(htmx.Text("Button"), htmx.HxPost("/api/respond"), htmx.HxSwap("outerHTML"), htmx.ClassNames{"inline-block cursor-pointer rounded-md bg-gray-800 px-4 py-3 text-center text-sm font-semibold uppercase text-white transition duration-200 ease-in-out hover:bg-gray-900": true}),
	},
})
