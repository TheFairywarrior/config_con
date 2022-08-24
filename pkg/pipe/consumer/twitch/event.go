package twitch

import (
	"config_con/pkg/pipe/queue"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TwitchEventConsumer struct {
	Name        string `yaml:"name"`
	EventSecret string `yaml:"eventSecret"`
	Url         string `yaml:"url"`
}

type twitchEventPayload struct {
	subscription struct {
		id        string
		status    string
		eventType string `json:"type"`
		version   string
		cost      int
		condition struct {
			broadCasterUserId string `json: "broadcaster_user_id"`
		}
		transport struct {
			method   string
			callback string
		}
		createdAt string `json:"created_at"`
	}
	event struct {
		userId               string `json:"user_id"`
		userLogin            string `json:"user_login"`
		userName             string `json:"user_name"`
		broadCasterUserId    string `json:"broadcaster_user_id"`
		broadCasterUserLogin string `json:"broadcaster_user_login"`
		broadCasterUserName  string `json:"broadcaster_user_name"`
	}
}

// verifyEvent takes the event signature and body and verifies it against the secret.
func (con TwitchEventConsumer) verifyEvent(message, messageSignature string) bool {
	prefix := "sha256="
	mac := hmac.New(sha256.New, []byte(con.EventSecret))
	mac.Write([]byte(prefix + message))
	sigCheck := mac.Sum(nil)
	return hmac.Equal([]byte(messageSignature), sigCheck)
}

func (con TwitchEventConsumer) EventRoute(ctx *fiber.Ctx) error {
	signature := ctx.GetReqHeaders()["twitch-eventsub-message-signature"]
	timestamp := ctx.GetReqHeaders()["twitch-eventsub-message-timestamp"]
	messageId := ctx.GetReqHeaders()["twitch-eventsub-message-id"]
	var payload twitchEventPayload
	err := ctx.BodyParser(&payload)
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
	payloadJson, _ := json.Marshal(&payload)
	if !con.verifyEvent(messageId + timestamp + string(payloadJson), signature) {
		return ctx.Status(400).JSON(
			fiber.Map{
				"error":   "Invalid signature",
				"message": "Signature does not match",
			},
		)
	}

	return nil
}

func (con TwitchEventConsumer) Consume(cxt context.Context, q queue.TransformerQueue) error {
	return nil
}
