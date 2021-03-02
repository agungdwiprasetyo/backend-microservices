package serviceshared

import (
	"errors"
	"monorepo/sdk"
	masterservice "monorepo/sdk/master-service"
	"net/http"

	"github.com/labstack/echo"
	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/tracer"
	"pkg.agungdp.dev/candi/wrapper"
)

// CheckPermission middleware
func CheckPermission(permissionCode string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := func(permCode string) error {
				trace := tracer.StartTrace(c.Request().Context(), "Middleware:CheckPermission")
				defer trace.Finish()
				ctx := trace.Context()

				tokenClaim := candishared.ParseTokenClaimFromContext(c.Request().Context())
				userID := tokenClaim.Additional.(map[string]interface{})["user_id"].(string)
				isAllowed, err := sdk.GetSDK().MasterService().CheckPermission(ctx, masterservice.PayloadCheckPermission{
					UserID: userID, PermissionCode: permCode,
				})
				if err != nil {
					trace.SetError(err)
				}

				if !isAllowed {
					return errors.New("You dont have permission to access this resource")
				}
				return nil
			}(permissionCode); err != nil {
				return wrapper.NewHTTPResponse(http.StatusForbidden, err.Error()).JSON(c.Response())
			}

			return next(c)
		}
	}
}
