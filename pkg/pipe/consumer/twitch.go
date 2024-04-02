package consumer

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/thefairywarrior/config_con/pkg/api"
	"github.com/thefairywarrior/config_con/pkg/pipe/queue"
	"github.com/thefairywarrior/config_con/pkg/utils/override"

	"github.com/gofiber/fiber/v2"
)

type TwitchEventConsumer struct {
	name        string
	eventSecret string
	url         string
}

func (con TwitchEventConsumer) Name() string {
	return con.name
}

func (con TwitchEventConsumer) EventSecret() string {
	return con.eventSecret
}

func (con TwitchEventConsumer) Url() string {
	return con.url
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

type TwitchEventData struct {
	Challenge    string       `json:"challenge"`
	Subscription subscription `json:"subscription"`
	Event        event        `json:"event"`
}

type TwitchEventMessage struct {
	queue.MessageData
	TwitchEventData
}

func (message TwitchEventMessage) GetData() (any, error) {
	jsonData, _ := json.Marshal(message.TwitchEventData)
	var data map[string]any
	err := json.Unmarshal(jsonData, &data)
	return data, err
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

// EventRoute is the actual function that going to be run when the consumer api is hit.
// It connects the headers, verification, and pushing to queue together while also holding the error handling
// for the request.
func (con TwitchEventConsumer) EventRoute(ctx override.FiberContext, q queue.Queue) error {
	signature, timestamp, messageId, messageType, err := getHeaders(ctx)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var payload TwitchEventData
	err = ctx.BodyParser(&payload)
	if err != nil {
		body := fiber.Map{
			"error":   "Invalid body",
			"message": err.Error(),
		}
		log.Println(body)
		ctx.Status(400)
		return ctx.JSON(
			body,
		)
	}
	if !con.verifyEvent(messageId+timestamp+string(ctx.Body()), signature) {
		ctx.Status(402)
		return ctx.JSON(
			fiber.Map{
				"error":   "Invalid signature",
				"message": "Signature does not match",
			},
		)
	}

	if messageType == "webhook_callback_verification" {
		ctx.Status(200)
		return ctx.Send([]byte(payload.Challenge))
	}

	message := TwitchEventMessage{
		MessageData:     queue.NewMessageData(),
		TwitchEventData: payload,
	}
	q.Add(message)

	ctx.Status(200)
	return nil
}

func (con TwitchEventConsumer) Consume(ctx context.Context, q queue.Queue) error {
	api.ApiRoutes <- con.name // Passing the name of the consumer to know when all the api consumers are ready.
	server := api.GetServer()
	return server.AddRoute("POST", con.Url(), func(ctx *fiber.Ctx) error {
		return con.EventRoute(ctx, q)
	})
}

func NewTwitchEventConsumer(name string, eventSecret string, url string) TwitchEventConsumer {
	return TwitchEventConsumer{
		name:        name,
		eventSecret: eventSecret,
		url:         url,
	}
}
