package handler

import (
	"fmt"
	"io"
	"lf/goLiveStreaming/internal/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IKeysHandler interface {
	AuthStreamingKey(ctx echo.Context) error
}

type keysHandler struct {
	keysService service.IKeyService
}

func NewHandler(serv service.IKeyService) IKeysHandler {
	return &keysHandler{
		keysService: serv,
	}
}

func (kh *keysHandler) AuthStreamingKey(ctx echo.Context) error {
	log.Default().Println("Running auth...")
	body := ctx.Request().Body
	defer body.Close()

	fields, _ := io.ReadAll(body)
	fmt.Println(string(fields))

	return ctx.String(http.StatusOK, "WORKING")
}
