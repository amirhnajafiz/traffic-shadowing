package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amirhnajafiz/traffic-shadowing/internal/http"
	"github.com/labstack/echo/v4"
)

// main function for starting
// a http service with echo framework.
func main() {
	// getting service port from command line arguments
	var port string

	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	} else {
		panic(fmt.Errorf("must give the port argument"))
	}

	// exporting service logs into a logs[port].txt
	f, err := os.OpenFile("logs"+os.Args[1]+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("creating log file failed: %w", err))
	}

	log.SetOutput(f)

	// creating a new handler
	h := http.Handler{}
	// creating a new echo app
	app := echo.New()
	// set the default handler
	app.GET("/", h.HandleGetRequests)
	// starting the application
	log.Println(app.Start(port))
}
