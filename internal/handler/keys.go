package handler

import (
	"fmt"
	"io"
	"lf/goLiveStreaming/internal/model"
	"lf/goLiveStreaming/internal/service"
	"log"
	"net/http"
	"strings"

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
	passedKeyValue := getStreamKey(fields)

	keys, err := kh.keysService.AuthStreamingKey(passedKeyValue.Name, passedKeyValue.Key)

	if err != nil {
		return ctx.String(http.StatusBadRequest, "problem with stream key")
	}

	if keys.Key == "" {
		return ctx.String(http.StatusForbidden, "forbidden user")
	}

	log.Default().Println("user authenticated")

	newStreamURL := fmt.Sprintf("rmtp://127.0.0.1:1935/hls-live/%s", keys.Name)
	log.Default().Println("Redirecting to:", newStreamURL)
	return ctx.Redirect(http.StatusFound, newStreamURL)
}

func getStreamKey(s []byte) model.Keys {
	var authValues model.Keys

	pairs := strings.Split(string(s), "&")

	for _, pair := range pairs {
		splitPair := strings.Split(pair, "=")
		key := splitPair[0]
		value := splitPair[1]

		if key == "name" {
			allPassedValues := strings.Split(value, "_")
			authValues.Name = allPassedValues[0]
			authValues.Key = allPassedValues[1]
		}
	}

	return authValues
}
