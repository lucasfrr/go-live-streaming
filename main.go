package main

import (
	"lf/goLiveStreaming/config/db"
	"lf/goLiveStreaming/internal/handler"
	"lf/goLiveStreaming/internal/repository"
	"lf/goLiveStreaming/internal/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := db.OpenConn()
	if err != nil {
		log.Fatalf("error connect database")
	}

	keysRepo := repository.NeyKeyRepository(db)
	keysServ := service.NeyKeyService(keysRepo)
	keysHandl := handler.NewHandler(keysServ)

	log.Default().Println("Routing...")

	e := echo.New()

	e.POST("/auth", keysHandl.AuthStreamingKey)

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "WORKING")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
