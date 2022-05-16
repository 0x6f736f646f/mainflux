// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package rabbitmq_test

import (
	"fmt"
	"testing"

	"github.com/mainflux/mainflux/pkg/messaging"
	"github.com/mainflux/mainflux/pkg/messaging/rabbitmq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	topic       = "topic"
	chansPrefix = "channels"
	channel     = "9b7b1b3f-b1b0-46a8-a717-b8213f9eda3b"
	subtopic    = "engine"
)

var (
	msgChan = make(chan messaging.Message)
	data    = []byte("payload")
)

func TestPubsub(t *testing.T) {
	// Test Subscribe and Unsubscribe
	subcases := []struct {
		desc         string
		topic        string
		topicID      string
		errorMessage error
		pubsub       bool //true for subscribe and false for unsubscribe
	}{
		{
			desc:         "Susbcribe to a topic with an ID",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "topicid1",
			errorMessage: nil,
			pubsub:       true,
		},
		{
			desc:         "Susbcribe to the same topic with a different ID",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "topicid2",
			errorMessage: nil,
			pubsub:       true,
		},
		{
			desc:         "Susbcribe to an already subscribed topic with an ID",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "topicid1",
			errorMessage: rabbitmq.ErrAlreadySubscribed,
			pubsub:       true,
		},
		{
			desc:         "Unsubscribe to a topic with an ID",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "topicid1",
			errorMessage: nil,
			pubsub:       false,
		},
		{
			desc:         "Unsubscribe to a non-existent topic with an ID",
			topic:        "h",
			topicID:      "topicid1",
			errorMessage: rabbitmq.ErrNotSubscribed,
			pubsub:       false,
		},
		{
			desc:         "Unsubscribe to the same topic with a different ID",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "topicid2",
			errorMessage: nil,
			pubsub:       false,
		},
		{
			desc:         "Unsubscribe to the same topic with a different ID not subscribed",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "topicid3",
			errorMessage: rabbitmq.ErrNotSubscribed,
			pubsub:       false,
		},
		{
			desc:         "Unsubscribe to an already unsubscribed topic with an ID",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "topicid1",
			errorMessage: rabbitmq.ErrNotSubscribed,
			pubsub:       false,
		},
		{
			desc:         "Susbcribe to a topic with a subtopic with an ID",
			topic:        fmt.Sprintf("%s.%s.%s", chansPrefix, topic, subtopic),
			topicID:      "topicid1",
			errorMessage: nil,
			pubsub:       true,
		},
		{
			desc:         "Susbcribe to an already subscribed topic with a subtopic with an ID",
			topic:        fmt.Sprintf("%s.%s.%s", chansPrefix, topic, subtopic),
			topicID:      "topicid1",
			errorMessage: rabbitmq.ErrAlreadySubscribed,
			pubsub:       true,
		},
		{
			desc:         "Unsubscribe to a topic with a subtopic with an ID",
			topic:        fmt.Sprintf("%s.%s.%s", chansPrefix, topic, subtopic),
			topicID:      "topicid1",
			errorMessage: nil,
			pubsub:       false,
		},
		{
			desc:         "Unsubscribe to an already unsubscribed topic with a subtopic with an ID",
			topic:        fmt.Sprintf("%s.%s.%s", chansPrefix, topic, subtopic),
			topicID:      "topicid1",
			errorMessage: rabbitmq.ErrNotSubscribed,
			pubsub:       false,
		},
		{
			desc:         "Susbcribe to an empty topic with an ID",
			topic:        "",
			topicID:      "topicid1",
			errorMessage: rabbitmq.ErrEmptyTopic,
			pubsub:       true,
		},
		{
			desc:         "Unsubscribe to an empty topic with an ID",
			topic:        "",
			topicID:      "topicid1",
			errorMessage: rabbitmq.ErrEmptyTopic,
			pubsub:       false,
		},
		{
			desc:         "Susbcribe to a topic with empty id",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "",
			errorMessage: rabbitmq.ErrEmptyID,
			pubsub:       true,
		},
		{
			desc:         "Unsubscribe to a topic with empty id",
			topic:        fmt.Sprintf("%s.%s", chansPrefix, topic),
			topicID:      "",
			errorMessage: rabbitmq.ErrEmptyID,
			pubsub:       false,
		},
	}

	for _, pc := range subcases {
		if pc.pubsub == true {
			err := pubsub.Subscribe(pc.topicID, pc.topic, handler{})
			if pc.errorMessage == nil {
				require.Nil(t, err, fmt.Sprintf("%s got unexpected error: %s", pc.desc, err))
			} else {
				assert.Equal(t, err, pc.errorMessage)
			}
		} else {
			err := pubsub.Unsubscribe(pc.topicID, pc.topic)
			if pc.errorMessage == nil {
				require.Nil(t, err, fmt.Sprintf("%s got unexpected error: %s", pc.desc, err))
			} else {
				assert.Equal(t, err, pc.errorMessage)
			}
		}
	}
}

type handler struct{}

func (h handler) Handle(msg messaging.Message) error {
	msgChan <- msg
	return nil
}

func (h handler) Cancel() error {
	return nil
}