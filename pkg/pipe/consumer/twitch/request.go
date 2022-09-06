package twitch

import (
	"config_con/pkg/utils/override"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// getHeaders takes the fiber context in and checks if all of the correct headers are present.
// If headers are missing an error is returned.
// If all of the headers are present, it seperates and returns them.
func getHeaders(ctx override.FiberContext) (string, string, string, string, error) {
	signature, sOk := ctx.GetReqHeaders()["twitch-eventsub-message-signature"]
	timestamp, tOk := ctx.GetReqHeaders()["twitch-eventsub-message-timestamp"]
	messageId, mOk := ctx.GetReqHeaders()["twitch-eventsub-message-id"]
	messageType, mTOk := ctx.GetReqHeaders()["twitch-eventsub-message-type"]

	if !sOk || !tOk || !mOk || !mTOk {
		return "", "", "", "", fmt.Errorf("missing headers, required headers are twitch-eventsub-message-signature, twitch-eventsub-message-timestamp, twitch-eventsub-message-id, twitch-eventsub-message-type")
	}

	return signature, timestamp, messageId, messageType, nil
}

// verifyEvent takes the event signature and body and verifies it against the secret.
func (con TwitchEventConsumer) verifyEvent(message, messageSignature string) bool {
	prefix := "sha256="
	mac := hmac.New(sha256.New, []byte(con.EventSecret))
	mac.Write([]byte(prefix + message))
	sigCheck := prefix + hex.EncodeToString(mac.Sum(nil))
	return messageSignature == sigCheck
}
