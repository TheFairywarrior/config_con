package twitch

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/utils/test"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestTwitchEventConsumer_EventRoute(t *testing.T) {
	eventSecret := "hello"
	timestamp := "2022-10-10T10:10:10Z"
	messageId := "12345"
	payload := TwitchEventData{
		Challenge: "",
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
	mac.Write([]byte(secretBody))
	h := mac.Sum(nil)
	signature := "sha256=" + hex.EncodeToString(h)
	fakeContext := test.FakeFiberContext{
		RequestBody: []byte(jsonData),
		Headers: map[string]string{
			"X-Hub-Signature":                   signature,
			"Twitch-Eventsub-Message-Signature": signature,
			"Twitch-Eventsub-Message-Id":        messageId,
			"Twitch-Eventsub-Message-Timestamp": timestamp,
			"Twitch-Eventsub-Message-Type":      "type",
		},
	}

	consumer := TwitchEventConsumer{
		Name:        "test",
		EventSecret: eventSecret,
		Url:         "url",
	}

	tQueue := queue.NewQueue(1)

	err := consumer.EventRoute(&fakeContext, tQueue)
	assert.NoError(t, err)
	assert.Equal(t, 200, fakeContext.CurrentStatus)
	<-tQueue.Chan()

	fakeContext = test.FakeFiberContext{
		RequestBody: []byte(jsonData),
		Headers: map[string]string{
			"X-Hub-Signature":                   signature,
			"Twitch-Eventsub-Message-Signature": signature,
			"Twitch-Eventsub-Message-Id":        messageId,
			"Twitch-Eventsub-Message-Timestamp": timestamp,
			"Twitch-Eventsub-Message-Type":      "webhook_callback_verification",
		},
	}
	err = consumer.EventRoute(&fakeContext, tQueue)
	assert.NoError(t, err)
	assert.Equal(t, 200, fakeContext.CurrentStatus)

	brokenData, _ := json.Marshal(fiber.Map{
		"broken": "data",
	})
	fakeContext = test.FakeFiberContext{
		RequestBody: brokenData,
		Headers: map[string]string{
			"X-Hub-Signature":                   signature,
			"Twitch-Eventsub-Message-Signature": "signature",
			"Twitch-Eventsub-Message-Id":        messageId,
			"Twitch-Eventsub-Message-Timestamp": timestamp,
			"Twitch-Eventsub-Message-Type":      "type",
		},
	}
	err = consumer.EventRoute(&fakeContext, tQueue)
	assert.NoError(t, err)
	assert.Equal(t, 402, fakeContext.CurrentStatus)
}
