package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (h *Handler) HandleGetRequests(ctx echo.Context) error {
	log.Printf("user %s", ctx.RealIP())

	return ctx.NoContent(http.StatusNoContent)
}
