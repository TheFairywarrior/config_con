package twitch

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/utils/override"
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
	Subscription subscription `json:"Subscription"`
	Event        event        `json:"event"`
}

// verifyEvent takes the event signature and body and verifies it against the secret.
func (con TwitchEventConsumer) verifyEvent(message, messageSignature string) bool {
	prefix := "sha256="
	mac := hmac.New(sha256.New, []byte(con.EventSecret))
	mac.Write([]byte(prefix + message))
	sigCheck := mac.Sum(nil)
	return hmac.Equal([]byte(messageSignature), sigCheck)
}

func (con TwitchEventConsumer) EventRoute(ctx override.FiberContext, q queue.TransformerQueue) error {
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
	if !con.verifyEvent(messageId+timestamp+string(payloadJson), signature) {
		return ctx.Status(400).JSON(
			fiber.Map{
				"error":   "Invalid signature",
				"message": "Signature does not match",
			},
		)
	}
	q.Add(payload)

	return nil
}

func (con TwitchEventConsumer) Consume(cxt context.Context, q queue.TransformerQueue) error {
	return nil
}
