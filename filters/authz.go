package filters

import (
	"github.com/gofiber/fiber/v2"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
)

// NewAuthzParamFilter creates a new filter to check the access of a principal to a specific object.
func NewAuthzParamFilter(action authz.AuthzAction, param string, checker authz.AuthzChecker) htmx.FilterFunc {
	return func(c *fiber.Ctx) error {
		p := authz.AuthzObject(c.Params(param, ""))

		principal, _, _, err := authz.AuthzFromContext(c)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		allowed, err := checker.Allowed(c.Context(), principal, p, action)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		if !allowed {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		return nil
	}
}
