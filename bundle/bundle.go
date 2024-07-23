package bundle

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	Next func(c *fiber.Ctx) bool
	// Bundle is the http.FileSystem to serve.
	Bundle fs.FS `json:"-"`
	// Root is the root directory of the Bundle.
	Root string `json:"root"`
	// PathPrefix is the optional prefix used to serve the filesystem content.
	PathPrefix string `json:"path_prefix"`
	// MaxAge defines the max-age of the Cache-Control header in seconds.
	MaxAge int `json:"max_age"`
	// ContentTypeCharset defines the charset of the Content-Type header.
	ContentTypeCharset string `json:"content_type_charset"`
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:               nil,
	Bundle:             htmx.Bundle,
	Root:               htmx.BundleFoler,
	PathPrefix:         "",
	MaxAge:             0,
	ContentTypeCharset: "",
}

// New ...
// nolint:gocyclo
func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		method := c.Method()

		if method != fiber.MethodGet && method != fiber.MethodHead {
			return c.Next()
		}

		p := strings.TrimPrefix(c.Path(), cfg.PathPrefix)

		f, err := cfg.Bundle.Open(path.Join(htmx.BundleFoler, p))
		if err != nil && errors.Is(err, fs.ErrNotExist) {
			return fiber.ErrNotFound
		}

		if err != nil {
			return fiber.ErrInternalServerError
		}

		stat, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat: %w", err)
		}

		if stat.IsDir() {
			return fiber.ErrForbidden
		}

		c.Type(getFileExtension(stat.Name()))

		c.Status(fiber.StatusOK)

		contentLength := int(stat.Size())

		if method == fiber.MethodGet {
			c.Response().SetBodyStream(f, contentLength)
			return nil
		}

		if method == fiber.MethodHead {
			c.Request().ResetBody()
			c.Response().SkipBody = true
			c.Response().Header.SetContentLength(contentLength)

			if err := f.Close(); err != nil {
				return fmt.Errorf("failed to close: %w", err)
			}

			return nil
		}

		return c.Next()
	}
}

func getFileExtension(p string) string {
	n := strings.LastIndexByte(p, '.')
	if n < 0 {
		return ""
	}
	return p[n:]
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.Bundle == nil {
		cfg.Bundle = ConfigDefault.Bundle
	}

	if cfg.Root == "" {
		cfg.Root = ConfigDefault.Root
	}

	return cfg
}
