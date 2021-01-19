package linebotapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"monorepo/services/line-chatbot/pkg/shared"
	"monorepo/services/line-chatbot/pkg/shared/domain"
	"net/http"
	"net/url"
	"strings"
	"time"

	"pkg.agungdwiprasetyo.com/candi/candiutils"
)

// botHTTP implementation
type botHTTP struct {
	httpReq candiutils.HTTPRequest
}

// NewLineBotHTTP constructor
func NewLineBotHTTP() Linebot {
	return &botHTTP{
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
		),
	}
}

// ProcessText method
func (b *botHTTP) ProcessText(ctx context.Context, text string) string {
	url := fmt.Sprintf("%s?input=%s", shared.GetEnv().ChatbotHost, url.QueryEscape(text))
	body, _, err := b.httpReq.Do(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return ""
	}

	var output struct {
		Output string `json:"output"`
	}
	json.Unmarshal(body, &output)

	return strings.TrimLeftFunc(output.Output, func(r rune) bool {
		return r == '-' || r == ' '
	})
}

// PushMessage push message to line channel
func (b *botHTTP) PushMessage(ctx context.Context, message *domain.LineMessage) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(true)
	encoder.Encode(message)

	url := "https://api.line.me/v2/bot/message/push"
	body, _, err := b.httpReq.Do(ctx, http.MethodPost, url, buffer.Bytes(), map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", shared.GetEnv().LineClientToken),
	})
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
