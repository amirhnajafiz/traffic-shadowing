package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
// manages to handle the requests.
type Handler struct{}

// HandleGetRequests
// handling user get requests.
func (h *Handler) HandleGetRequests(ctx echo.Context) error {
	log.Printf("user %s", ctx.RealIP())

	return ctx.NoContent(http.StatusNoContent)
}
