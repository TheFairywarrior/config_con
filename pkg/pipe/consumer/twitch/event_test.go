package twitch

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/utils/test"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitchEventConsumer_EventRoute(t *testing.T) {
	eventSecret := "hello"
	timestamp := "2022-10-10T10:10:10Z"
	messageId := "12345"
	payload := twitchEventPayload{
		Subscription: subscription{
			Id:        "id",
			Status:    "status",
			EventType: "type",
			Version:   "version",
			Cost:      1,
			Condition: condition{
				BroadCasterUserId: "broadcaster_user_id",
			},
			Transport: transport{
				Method:   "method",
				Callback: "callback",
			},
			CreatedAt: "created_at",
		},
		Event: event{
			UserId:               "user_id",
			UserLogin:            "user_login",
			UserName:             "user_name",
			BroadCasterUserId:    "broadcaster_user_id",
			BroadCasterUserLogin: "broadcaster_user_login",
			BroadCasterUserName:  "broadcaster_user_name",
		},
	}
	jsonData, _ := json.Marshal(payload)
	mac := hmac.New(sha256.New, []byte(eventSecret))
	secretBody := messageId + timestamp + string(jsonData)
	mac.Write([]byte("sha256=" + secretBody))
	h := mac.Sum(nil)
	signature := "sha256=" + hex.EncodeToString(h)
	fakeContext := test.FakeFiberContext{
		Body: []byte(jsonData),
		Headers: map[string]string{
			"X-Hub-Signature":                   signature,
			"twitch-eventsub-message-signature": signature,
			"twitch-eventsub-message-id":        messageId,
			"twitch-eventsub-message-timestamp": timestamp,
			"twitch-eventsub-message-type":      "type",
		},
	}

	consumer := TwitchEventConsumer{
		Name:        "test",
		EventSecret: eventSecret,
		Url:         "url",
	}

	tQueue := queue.NewTransformerQueue(1)

	err := consumer.EventRoute(&fakeContext, tQueue)
	assert.NoError(t, err)
	assert.Equal(t, 200, fakeContext.CurrentStatus)
	<-tQueue.Chan()
}
