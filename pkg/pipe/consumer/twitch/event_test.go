package twitch

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
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
	mac.Write([]byte(secretBody))
	h := mac.Sum(nil)
	signature := "sha256=" + hex.EncodeToString(h)
	fmt.Println(signature)

}
