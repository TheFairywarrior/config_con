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

type TwitchEventData struct {
	Subscription subscription `json:"subscription"`
	Event        event        `json:"event"`
}

type TwitchEventMessage struct {
	queue.MessageData
	TwitchEventData
}

func (message TwitchEventMessage) GetData() any {
	return message.TwitchEventData
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

	if messageType == "webhook_callback_verification" {
		ctx.Status(200)
		return ctx.JSON(
			fiber.Map{
				"message": "Verified",
			},
		)
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
	payloadJson, _ := json.Marshal(payload)
	if !con.verifyEvent(messageId+timestamp+string(payloadJson), signature) {
		ctx.Status(402)
		return ctx.JSON(
			fiber.Map{
				"error":   "Invalid signature",
				"message": "Signature does not match",
			},
		)
	}

	message := TwitchEventMessage{
		MessageData:     queue.NewMessageData(),
		TwitchEventData: payload,
	}
	q.Add(message)

	ctx.Status(200)
	return ctx.JSON(fiber.Map{
		"message": "Success",
	})
}

func (con TwitchEventConsumer) Consume(cxt context.Context, q queue.Queue) error {
	server := api.GetServer()
	return server.AddRoute("GET", con.Url, func(ctx *fiber.Ctx) error {
		return con.EventRoute(ctx, q)
	})
}
