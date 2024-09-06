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
		log.Fatalf("Error connect database")
	}

	keyRepository := repository.NeyKeyRepository(db)
	keysService := service.NeyKeyService(keyRepository)
	keysHandler := handler.NewHandler(keysService)

	log.Default().Println("Routing...")
	e := echo.New()
	e.POST("/auth", keysHandler.AuthStreamingKey)

	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "working")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
