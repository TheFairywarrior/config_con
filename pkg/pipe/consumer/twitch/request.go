package twitch

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/utils/override"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type condition struct {
	BroadCasterUserId string `json:"broadcaster_user_id"`
}

type transport struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
}

type subscription struct {
	Id        string    `json:"id"`
	Status    string    `json:"status"`
	EventType string    `json:"type"`
	Version   string    `json:"version"`
	Cost      int       `json:"cost"`
	Condition condition `json:"condition"`
	Transport transport `json:"transport"`
	CreatedAt string    `json:"created_at"`
}

type event struct {
	UserId               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadCasterUserId    string `json:"broadcaster_user_id"`
	BroadCasterUserLogin string `json:"broadcaster_user_login"`
	BroadCasterUserName  string `json:"broadcaster_user_name"`
}

type TwitchEventData struct {
	Challenge    string       `json:"challenge"`
	Subscription subscription `json:"subscription"`
	Event        event        `json:"event"`
}

type TwitchEventMessage struct {
	queue.MessageData
	TwitchEventData
}

// getHeaders takes the fiber context in and checks if all of the correct headers are present.
// If headers are missing an error is returned.
// If all of the headers are present, it seperates and returns them.
func getHeaders(ctx override.FiberContext) (string, string, string, string, error) {
	signature, sOk := ctx.GetReqHeaders()["Twitch-Eventsub-Message-Signature"]
	timestamp, tOk := ctx.GetReqHeaders()["Twitch-Eventsub-Message-Timestamp"]
	messageId, mOk := ctx.GetReqHeaders()["Twitch-Eventsub-Message-Id"]
	messageType, mTOk := ctx.GetReqHeaders()["Twitch-Eventsub-Message-Type"]

	if !sOk || !tOk || !mOk || !mTOk {
		return "", "", "", "", fmt.Errorf("missing headers, required headers are twitch-eventsub-message-signature, twitch-eventsub-message-timestamp, twitch-eventsub-message-id, twitch-eventsub-message-type")
	}

	return signature, timestamp, messageId, messageType, nil
}

// verifyEvent takes the event signature and body and verifies it against the secret.
func (con TwitchEventConsumer) verifyEvent(message, messageSignature string) bool {
	prefix := "sha256="
	mac := hmac.New(sha256.New, []byte(con.eventSecret))
	mac.Write([]byte(message))
	sigCheck := prefix + hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(sigCheck), []byte(messageSignature))
}
