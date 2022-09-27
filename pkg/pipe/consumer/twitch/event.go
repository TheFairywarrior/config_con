package twitch

import (
	"config_con/pkg/api"
	"config_con/pkg/pipe/queue"
	"config_con/pkg/utils/override"
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TwitchEventConsumer struct {
	name        string `yaml:"name"`
	eventSecret string `yaml:"eventSecret"`
	url         string `yaml:"url"`
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

func NewTwitchEventConsumer(name string, eventSecret string, url string) TwitchEventConsumer {
	return TwitchEventConsumer{
		name:        name,
		eventSecret: eventSecret,
		url:         url,
	}
}

func (message TwitchEventMessage) GetData() (any, error) {
	jsonData, _ := json.Marshal(message.TwitchEventData)
	var data map[string]any
	err := json.Unmarshal(jsonData, &data)
	return data, err
}

// EventRoute is the actual function that going to be run when the consumer api is hit.
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

func (con TwitchEventConsumer) Consume(cxt context.Context, q queue.Queue) error {
	server := api.GetServer()
	return server.AddRoute("POST", con.Url(), func(ctx *fiber.Ctx) error {
		return con.EventRoute(ctx, q)
	})
}
