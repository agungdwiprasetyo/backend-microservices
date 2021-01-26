// Code generated by candi v1.3.1.

package resthandler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"

	"monorepo/services/line-chatbot/internal/modules/chatbot/domain"
	"monorepo/services/line-chatbot/internal/modules/chatbot/usecase"
	"monorepo/services/line-chatbot/pkg/shared"

	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
	"pkg.agungdwiprasetyo.com/candi/wrapper"
)

// RestHandler handler
type RestHandler struct {
	mw        interfaces.Middleware
	uc        usecase.ChatbotUsecase
	validator interfaces.Validator
}

// NewRestHandler create new rest handler
func NewRestHandler(mw interfaces.Middleware, uc usecase.ChatbotUsecase, validator interfaces.Validator) *RestHandler {
	return &RestHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root *echo.Group) {
	bot := root.Group("/v1/bot")

	bot.POST("/callback", h.callback)
	bot.POST("/pushmessage", h.pushMessage, h.mw.HTTPBearerAuth())
}

func (h *RestHandler) callback(c echo.Context) error {
	trace := tracer.StartTrace(c.Request().Context(), "ChatbotDeliveryREST:Callback")
	defer trace.Finish()
	ctx := trace.Context()

	req := c.Request()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	signature := req.Header.Get("X-Line-Signature")
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusUnauthorized, err.Error()).JSON(c.Response())
	}

	hash := hmac.New(sha256.New, []byte(shared.GetEnv().LineClientSecret))
	_, err = hash.Write(body)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	if !hmac.Equal(decoded, hash.Sum(nil)) {
		return wrapper.NewHTTPResponse(http.StatusUnauthorized, err.Error()).JSON(c.Response())
	}

	request := struct {
		Events []*linebot.Event `json:"events"`
	}{}
	if err = json.Unmarshal(body, &request); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	if err := h.uc.ProcessCallback(ctx, request.Events); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "ok").JSON(c.Response())
}

func (h *RestHandler) pushMessage(c echo.Context) error {
	trace := tracer.StartTrace(c.Request().Context(), "ChatbotDeliveryREST:PushMessage")
	defer trace.Finish()
	ctx := trace.Context()

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	if err := h.validator.ValidateDocument("push-message", body); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(c.Response())
	}

	var payload domain.PushMessagePayload
	if err = json.Unmarshal(body, &payload); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	if err := h.uc.PushMessageToChannel(ctx, payload); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "ok").JSON(c.Response())
}
