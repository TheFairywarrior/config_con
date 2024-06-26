package consumer

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/thefairywarrior/config_con/pkg/pipe/queue"
	"github.com/thefairywarrior/config_con/pkg/utils/test"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestTwitchEventConsumer_EventRoute(t *testing.T) {
	eventSecret := "hello"
	timestamp := "2022-10-10T10:10:10Z"
	messageId := "12345"
	payload := TwitchEventData{
		Challenge: "",
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
	mac := hmac.New(sha256.New, []byte(eventSecret))
	secretBody := messageId + timestamp + string(jsonData)
	mac.Write([]byte(secretBody))
	h := mac.Sum(nil)
	signature := "sha256=" + hex.EncodeToString(h)
	fakeContext := test.FakeFiberContext{
		RequestBody: []byte(jsonData),
		Headers: map[string]string{
			"X-Hub-Signature":                   signature,
			"Twitch-Eventsub-Message-Signature": signature,
			"Twitch-Eventsub-Message-Id":        messageId,
			"Twitch-Eventsub-Message-Timestamp": timestamp,
			"Twitch-Eventsub-Message-Type":      "type",
		},
	}

	consumer := TwitchEventConsumer{
		name:        "test",
		eventSecret: eventSecret,
		url:         "url",
	}

	tQueue := queue.NewLocalQueue(1)

	err := consumer.EventRoute(&fakeContext, tQueue)
	assert.NoError(t, err)
	assert.Equal(t, 200, fakeContext.CurrentStatus)
	<-tQueue.Chan()

	fakeContext = test.FakeFiberContext{
		RequestBody: []byte(jsonData),
		Headers: map[string]string{
			"X-Hub-Signature":                   signature,
			"Twitch-Eventsub-Message-Signature": signature,
			"Twitch-Eventsub-Message-Id":        messageId,
			"Twitch-Eventsub-Message-Timestamp": timestamp,
			"Twitch-Eventsub-Message-Type":      "webhook_callback_verification",
		},
	}
	err = consumer.EventRoute(&fakeContext, tQueue)
	assert.NoError(t, err)
	assert.Equal(t, 200, fakeContext.CurrentStatus)

	brokenData, _ := json.Marshal(fiber.Map{
		"broken": "data",
	})
	fakeContext = test.FakeFiberContext{
		RequestBody: brokenData,
		Headers: map[string]string{
			"X-Hub-Signature":                   signature,
			"Twitch-Eventsub-Message-Signature": "signature",
			"Twitch-Eventsub-Message-Id":        messageId,
			"Twitch-Eventsub-Message-Timestamp": timestamp,
			"Twitch-Eventsub-Message-Type":      "type",
		},
	}
	err = consumer.EventRoute(&fakeContext, tQueue)
	assert.NoError(t, err)
	assert.Equal(t, 402, fakeContext.CurrentStatus)
}

func TestTwitchEventConsumer_Getters(t *testing.T) {
	consumer := TwitchEventConsumer{
		name:        "test",
		eventSecret: "secret",
		url:         "url",
	}
	assert.Equal(t, "test", consumer.Name())
	assert.Equal(t, "secret", consumer.EventSecret())
	assert.Equal(t, "url", consumer.Url())
}

func TestNewTwitchEventConsumer(t *testing.T) {
	consumer := NewTwitchEventConsumer("test", "secret", "url")
	assert.Equal(t, "test", consumer.Name())
	assert.Equal(t, "secret", consumer.EventSecret())
	assert.Equal(t, "url", consumer.Url())
}

func TestTwitchEventConsumer_Consume(t *testing.T) {
	q := queue.NewLocalQueue(1)
	defer q.Close()

	type fields struct {
		name        string
		eventSecret string
		url         string
	}
	type args struct {
		ctx context.Context
		q   queue.Queue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestTwitchEventConsumer_Consume",
			fields: fields{
				name:        "test",
				eventSecret: "secret",
				url:         "url",
			},
			args: args{
				ctx: context.Background(),
				q:   q,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := TwitchEventConsumer{
				name:        tt.fields.name,
				eventSecret: tt.fields.eventSecret,
				url:         tt.fields.url,
			}
			if err := con.Consume(tt.args.ctx, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("TwitchEventConsumer.Consume() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTwitchEventMessage_GetData(t *testing.T) {
	type fields struct {
		MessageData     queue.MessageData
		TwitchEventData TwitchEventData
	}
	tests := []struct {
		name    string
		fields  fields
		want    any
		wantErr bool
	}{
		{
			name: "TestTwitchEventMessage_GetData",
			fields: fields{
				MessageData: queue.NewMessageData(),
				TwitchEventData: TwitchEventData{
					Event: event{
						UserId:               "123",
						UserLogin:            "test",
						UserName:             "test",
						BroadCasterUserId:    "123",
						BroadCasterUserLogin: "test",
						BroadCasterUserName:  "test",
					},
				},
			},
			want: map[string]interface{}{
				"challenge": "",
				"subscription": map[string]any{
					"id":      "",
					"status":  "",
					"type":    "",
					"version": "",
					"cost":    float64(0),
					"condition": map[string]any{
						"broadcaster_user_id": "",
					},
					"transport": map[string]any{
						"method":   "",
						"callback": "",
					},
					"created_at": "",
				},
				"event": map[string]interface{}{
					"broadcaster_user_id":    "123",
					"broadcaster_user_login": "test",
					"broadcaster_user_name":  "test",
					"user_id":                "123",
					"user_login":             "test",
					"user_name":              "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := TwitchEventMessage{
				MessageData:     tt.fields.MessageData,
				TwitchEventData: tt.fields.TwitchEventData,
			}
			got, err := message.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("TwitchEventMessage.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwitchEventMessage.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
