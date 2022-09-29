package queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMessageData_Getters(t *testing.T) {
	testTime := time.Date(2022, 9, 29, 0, 0, 0, 0, time.Now().Location())
	data := MessageData{
		id:        "test",
		timestamp: testTime,
	}
	assert.Equal(t, "test", data.ID())
	assert.Equal(t, testTime, data.Timestamp())
	assert.Equal(t, data, data.Meta())
}
