package main

import (
	"log"

	"github.com/amirhnajafiz/traffic-shadowing/internal/http"
	"github.com/labstack/echo/v4"
)

func main() {
	h := http.Handler{}
	app := echo.New()

	app.GET("/", h.HandleGetRequests)

	log.Println(app.Start(":8080"))
}
