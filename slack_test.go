package slack

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PostWebhookMessage(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		message        *WebhookMessage
		expectedErr    error
		expectedStatus int
	}{
		{
			name: "success",
			url:  os.Getenv("SLACK_WEBHOOK_URL"),
			message: &WebhookMessage{
				Channel:  os.Getenv("SLACK_CHANNEL_ID"),
				UserName: "test",
				Text:     "hello!",
			},
			expectedStatus: 200,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := PostWebhookMessage(test.url, test.message)
			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			if test.expectedStatus != 0 {
				assert.Equal(t, res.StatusCode, test.expectedStatus)
			} else {
				assert.Nil(t, res)
			}
		})
	}
}
