package sse

import (
	"bufio"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// Event defines the interface for content that can be broadcast to clients.
type Event interface {
	String() string // Represent the envelope contents as a string for transmission.
}

// Sender is the interface for sending server-sent events.
type Sender interface {
	ID() string
	Events() chan Event
}

// Config is the configuration for the server-sent events server.
type Config struct{}

// NewSSEHandler creates a new server-sent events handler.
func NewSSEHandler(sender Sender, config ...Config) fiber.Handler {
	_ = configDefault(config...)

	return func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "text/event-stream")
		c.Set(fiber.HeaderCacheControl, "no-cache")
		c.Set(fiber.HeaderConnection, "keep-alive")
		c.Set(fiber.HeaderTransferEncoding, "chunked")

		c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			for {
				select {
				case msg, ok := <-sender.Events():
					if !ok {
						return
					}

					_, err := fmt.Fprint(w, msg.String())
					if err != nil {
						return
					}

					err = w.Flush()
					if err != nil {
						// Refreshing page in web browser will establish a new
						// SSE connection, but only (the last) one is alive, so
						// dead connections must be closed here.
						return
					}
				case <-c.Context().Done():
					return
				}
			}
		}))

		return nil
	}
}

// ConfigDefault is the default configuration for the server-sent events server.
var ConfigDefault = Config{}

// Helper function to set default values
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	return cfg
}
