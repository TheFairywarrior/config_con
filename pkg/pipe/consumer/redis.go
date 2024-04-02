package consumer

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/thefairywarrior/config_con/pkg/pipe/queue"
)

type RedisMessageData struct {
	ID   string         `json:id`
	Type string         `json:type`
	Data map[string]any `json:data`
}

type RedisMessage struct {
	queue.MessageData
	content RedisMessageData
}

func (msg RedisMessage) GetData() (any, error) {
	return msg.content, nil
}

func NewRedisMessage(id string, content RedisMessageData) RedisMessage {
	return RedisMessage{
		queue.NewMessageWithId(id),
		content,
	}
}

type RedisConsumer struct {
	url      string
	password string
	database int
	channel  string
}

func (con RedisConsumer) Consume(ctx context.Context, q queue.Queue) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     con.url,
		Password: con.password,
		DB:       con.database,
	})

	pubsub := rdb.Subscribe(ctx, con.channel)

	if _, err := pubsub.Receive(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-pubsub.Channel():
			messageData := RedisMessageData{}
			if err := json.Unmarshal([]byte(msg.Payload), &messageData); err != nil {
				q.Add(NewRedisMessage(
					messageData.ID,
					messageData,
				))
			}
		}
	}
}


func NewRedisConsumer(url string, password string, database int, channel string) RedisConsumer {
	return RedisConsumer{
		url:      url,
		password: password,
		database: database,
		channel:  channel,
	}
}
