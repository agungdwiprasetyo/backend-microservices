package serviceshared

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/tracer"
	"pkg.agungdp.dev/candi/wrapper"
)

var errForbidden = errors.New("You dont have permission to access this resource")

// ACLChecker abstraction
type ACLChecker interface {
	CheckPermission(ctx context.Context, userID string, permissionCode string) (err error)
}

// CheckPermission middleware
func CheckPermission(aclChecker ACLChecker, permissionCode string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := func(permCode string) (err error) {
				trace := tracer.StartTrace(c.Request().Context(), "Middleware:CheckPermission")
				defer func() {
					if r := recover(); r != nil {
						err = errForbidden
					}
					trace.Finish()
				}()
				ctx := trace.Context()

				tokenClaim := candishared.ParseTokenClaimFromContext(c.Request().Context())
				userID := tokenClaim.Additional.(map[string]interface{})["user_id"].(string)
				err = aclChecker.CheckPermission(ctx, userID, permCode)
				if err != nil {
					trace.SetError(err)
					err = errForbidden
				}
				return
			}(permissionCode); err != nil {
				return wrapper.NewHTTPResponse(http.StatusForbidden, err.Error()).JSON(c.Response())
			}

			return next(c)
		}
	}
}
