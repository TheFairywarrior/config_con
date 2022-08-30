package twitch

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/utils/override"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TwitchEventConsumer struct {
	Name        string `yaml:"name"`
	EventSecret string `yaml:"eventSecret"`
	Url         string `yaml:"url"`
}

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

type twitchEventPayload struct {
	Subscription subscription `json:"subscription"`
	Event        event        `json:"event"`
}

// verifyEvent takes the event signature and body and verifies it against the secret.
func (con TwitchEventConsumer) verifyEvent(message, messageSignature string) bool {
	prefix := "sha256="
	mac := hmac.New(sha256.New, []byte(con.EventSecret))
	mac.Write([]byte(prefix + message))
	sigCheck := prefix + hex.EncodeToString(mac.Sum(nil))
	return messageSignature == sigCheck
}

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


// EventRoute is the actual function that going to be run when the consumer api is hit.
func (con TwitchEventConsumer) EventRoute(ctx override.FiberContext, q queue.TransformerQueue) error {
	signature, timestamp, messageId, messageType, err := getHeaders(ctx)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var payload twitchEventPayload
	err = ctx.BodyParser(&payload)
	if err != nil {
		body := fiber.Map{
			"error":   "Invalid body",
			"message": err.Error(),
		}
		log.Println(body)
		return ctx.Status(400).JSON(
			body,
		)
	}
	payloadJson, _ := json.Marshal(payload)
	if !con.verifyEvent(messageId+timestamp+string(payloadJson), signature) {
		return ctx.Status(400).JSON(
			fiber.Map{
				"error":   "Invalid signature",
				"message": "Signature does not match",
			},
		)
	}

	if messageType == "webhook_callback_verification" {
		return ctx.Status(200).JSON(
			fiber.Map{
				"message": "Verified",
			},
		)
	}
	q.Add(payload)

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func (con TwitchEventConsumer) Consume(cxt context.Context, q queue.TransformerQueue) error {
	return nil
}
