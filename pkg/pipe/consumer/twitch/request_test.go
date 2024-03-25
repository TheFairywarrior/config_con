package twitch

import (
	"github.com/thefairywarrior/config_con/pkg/pipe/queue"
	"reflect"
	"testing"
)

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
