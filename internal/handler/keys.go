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

type KeysHandler interface {
	AuthStreamingKey(ctx echo.Context) error
}

type keysHandler struct {
	keysService service.KeyService
}

func NewHandler(s service.KeyService) *keysHandler {
	return &keysHandler{
		keysService: s,
	}
}

func (h *keysHandler) AuthStreamingKey(ctx echo.Context) error {
	log.Default().Println("Running auth...")
	body := ctx.Request().Body
	defer body.Close()
	fields, _ := io.ReadAll(body)
	authValues := getKeyValues(fields)

	keys, err := h.keysService.AuthStreamingKey(authValues.Name, authValues.Key)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Error finding key")
	}

	if keys.Key == "" {
		log.Default().Println("Forbidden User")
		return ctx.String(http.StatusForbidden, "")
	}

	log.Default().Println("User authenticated")

	newStreamURL := fmt.Sprintf("rtmp://127.0.0.1:1935/hls-live/%s", keys.Name)
	log.Default().Println("Redirecting to:", newStreamURL)
	return ctx.Redirect(http.StatusFound, newStreamURL)
}

func getKeyValues(s []byte) model.Keys {
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
