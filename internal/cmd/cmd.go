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

	h := http.Handler{}
	app := echo.New()

	app.GET("/", h.HandleGetRequests)

	log.Println(app.Start(port))
}
