package twitch

import (
	"encoding/json"
	"testing"
)

func TestTwitchEventConsumer_EventRoute(t *testing.T) {
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
}
