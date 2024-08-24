package handler

import (
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

	if keys.Key != "" {
		log.Default().Println("user authenticated")
		return ctx.String(http.StatusOK, "OK")
	}

	return ctx.String(http.StatusForbidden, "forbidden")
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
