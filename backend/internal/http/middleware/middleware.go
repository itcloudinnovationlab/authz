package middleware

import (
	"github.com/eko/authz/backend/internal/manager"
	"github.com/eko/authz/backend/internal/security/paseto"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

const (
	AuthenticationKey = "authentication"
	AuthorizationKey  = "authorization"
)

type Middlewares map[string]fiber.Handler

func (m Middlewares) Get(name string) fiber.Handler {
	return m[name]
}

func NewMiddlewares(
	logger *slog.Logger,
	manager manager.Manager,
	tokenManager paseto.Manager,
) Middlewares {
	return Middlewares{
		AuthenticationKey: Authentication(logger, tokenManager),
		AuthorizationKey:  Authorization(logger, manager),
	}
}