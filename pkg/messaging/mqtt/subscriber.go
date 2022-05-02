// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogo/protobuf/proto"

	log "github.com/mainflux/mainflux/logger"
	"github.com/mainflux/mainflux/pkg/errors"
	"github.com/mainflux/mainflux/pkg/messaging"
)

var (
	errSubscribeTimeout       = errors.New("failed to subscribe due to timeout reached")
	errUnsubscribeTimeout     = errors.New("failed to unsubscribe due to timeout reached")
	errUnsubscribeDeleteTopic = errors.New("failed to unsubscribe due to deletion of topic")
	errAlreadySubscribed      = errors.New("already subscribed to topic")
	errNotSubscribed          = errors.New("not subscribed")
	errEmptyTopic             = errors.New("empty topic")
	errEmptyID                = errors.New("empty ID")
)

var _ messaging.Subscriber = (*subscriber)(nil)

type subscription struct {
	client mqtt.Client
	topics []string
}

// contains checks if a topic is present
func (sub subscription) contains(topic string) bool {
	for _, v := range sub.topics {
		if v == topic {
			return true
		}
	}
	return false
}

// Finds the index of an item in the topics
func (sub subscription) indexOf(element string) int {
	for k, v := range sub.topics {
		if element == v {
			return k
		}
	}
	return -1
}

// Deletes a topic from the slice
func (sub subscription) delete(topic string) bool {
	index := sub.indexOf(topic)
	if index == -1 {
		return false
	}
	copy(sub.topics[index:], sub.topics[index+1:])
	sub.topics[len(sub.topics)-1] = ""
	return true
}

type subscriber struct {
	address       string
	timeout       time.Duration
	logger        log.Logger
	subscriptions map[string]subscription
}

// NewSubscriber returns a new MQTT message subscriber.
func NewSubscriber(address string, timeout time.Duration, logger log.Logger) (messaging.Subscriber, error) {
	ret := subscriber{
		address:       address,
		timeout:       timeout,
		logger:        logger,
		subscriptions: make(map[string]subscription),
	}
	return ret, nil
}

func (sub subscriber) Subscribe(id, topic string, handler messaging.MessageHandler) error {
	if id == "" {
		return errEmptyID
	}
	if topic == "" {
		return errEmptyTopic
	}
	// Check client ID
	s, ok := sub.subscriptions[id]
	if !ok {
		opts := mqtt.NewClientOptions().SetUsername(username).AddBroker(sub.address)
		client := mqtt.NewClient(opts)
		token := client.Connect()
		if token.Error() != nil {
			return token.Error()
		}
		s = subscription{
			client: client,
			topics: []string{topic},
		}
	}
	// Check topic
	ok = s.contains(topic)
	if ok {
		return errAlreadySubscribed
	}
	token := s.client.Subscribe(topic, qos, sub.mqttHandler(handler))
	if token.Error() != nil {
		return token.Error()
	}
	if ok := token.WaitTimeout(sub.timeout); !ok {
		return errSubscribeTimeout
	}
	return token.Error()
}

func (sub subscriber) Unsubscribe(id, topic string) error {
	if id == "" {
		return errEmptyID
	}
	if topic == "" {
		return errEmptyTopic
	}
	// Check client ID
	s, ok := sub.subscriptions[id]
	switch ok {
	case true:
		// Check topic
		if ok := s.contains(topic); !ok {
			return errNotSubscribed
		}
	default:
		return errNotSubscribed
	}
	token := s.client.Unsubscribe(topic)
	if token.Error() != nil {
		return token.Error()
	}

	ok = token.WaitTimeout(sub.timeout)
	if !ok {
		return errUnsubscribeTimeout
	}
	if ok := s.delete(topic); !ok {
		return errUnsubscribeDeleteTopic
	}
	if len(s.topics) == 0 {
		delete(sub.subscriptions, id)
	}
	return token.Error()
}

func (sub subscriber) mqttHandler(h messaging.MessageHandler) mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		var msg messaging.Message
		if err := proto.Unmarshal(m.Payload(), &msg); err != nil {
			sub.logger.Warn(fmt.Sprintf("Failed to unmarshal received message: %s", err))
			return
		}
		if err := h.Handle(msg); err != nil {
			sub.logger.Warn(fmt.Sprintf("Failed to handle Mainflux message: %s", err))
		}
	}
}
