// Code generated by candi v1.3.1.

package resthandler

import (
	"net/http"

	"github.com/labstack/echo"

	"monorepo/services/notification-service/internal/modules/push-notif/usecase"

	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
	"pkg.agungdp.dev/candi/wrapper"
)

// RestHandler handler
type RestHandler struct {
	mw        interfaces.Middleware
	uc        usecase.PushNotifUsecase
	validator interfaces.Validator
}

// NewRestHandler create new rest handler
func NewRestHandler(mw interfaces.Middleware, uc usecase.PushNotifUsecase, validator interfaces.Validator) *RestHandler {
	return &RestHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root *echo.Group) {
	v1Root := root.Group(candihelper.V1)

	pushnotif := v1Root.Group("/pushnotif")
	pushnotif.GET("", h.hello, echo.WrapMiddleware(h.mw.HTTPBearerAuth))
}

func (h *RestHandler) hello(c echo.Context) error {
	trace := tracer.StartTrace(c.Request().Context(), "DeliveryREST:Hello")
	defer trace.Finish()
	ctx := trace.Context()

	tokenClaim := c.Get(string(candishared.ContextKeyTokenClaim)).(*candishared.TokenClaim) // must using HTTPBearerAuth in middleware for this handler

	return wrapper.NewHTTPResponse(http.StatusOK, h.uc.Hello(ctx)+", with your session ("+tokenClaim.Audience+")").JSON(c.Response())
}
