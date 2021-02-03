package translator

import (
	"context"
	"encoding/json"
	"monorepo/services/line-chatbot/pkg/shared"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo"
	"pkg.agungdp.dev/candi/candiutils"
)

// translatorHTTPImpl implementation
type translatorHTTPImpl struct {
	httpReq candiutils.HTTPRequest
}

// NewTranslatorHTTP constructor
func NewTranslatorHTTP() Translator {
	return &translatorHTTPImpl{
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
		),
	}
}

// Translate method
func (t *translatorHTTPImpl) Translate(ctx context.Context, from, to, text string) (result string) {
	value := url.Values{}
	value.Set("key", shared.GetEnv().TranslatorKey)
	value.Set("lang", from+"-"+to)
	value.Add("text", text)

	resp, respCode, err := t.httpReq.Do(ctx, http.MethodPost, shared.GetEnv().TranslatorHost, []byte(value.Encode()), map[string]string{
		echo.HeaderContentType: echo.MIMEApplicationForm,
	})
	if err != nil {
		return err.Error()
	}

	if respCode >= http.StatusBadRequest {
		return "<Cannot process request>"
	}

	var response struct {
		Code int      `json:"code"`
		Lang string   `json:"lang"`
		Text []string `json:"text"`
	}

	json.Unmarshal(resp, &response)
	return strings.Join(response.Text, " ")
}
