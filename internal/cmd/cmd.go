package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amirhnajafiz/traffic-shadowing/internal/http"
	"github.com/labstack/echo/v4"
)

func main() {
	var port string

	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	} else {
		panic(fmt.Errorf("must give the port argument"))
	}

	f, err := os.OpenFile("logs"+os.Args[1]+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("creating log file failed: %w", err))
	}

	log.SetOutput(f)

	h := http.Handler{}
	app := echo.New()

	app.GET("/", h.HandleGetRequests)

	log.Println(app.Start(port))
}
