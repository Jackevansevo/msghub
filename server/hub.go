package server

import (
	"log"

	"github.com/Jackevansevo/msghub/client"
)

// Publisher represents a single message hub
type Publisher interface {
	Publish(message client.Message)
	Subscribe(topic string, client *client.Client)
	Unsubscribe(topic string, client *client.Client)
}

// Hub does stuff
type Hub struct {
	Publisher
	log         *log.Logger
	subscribers map[string][]*client.Client
}

// NewHub returns a new Hub
func NewHub(log *log.Logger) *Hub {
	return &Hub{
		log:         log,
		subscribers: make(map[string][]*client.Client),
	}
}

// Subscribe subscribes a user to a specific topic
func (hub *Hub) Subscribe(topic string, client *client.Client) {
	subs := hub.subscribers[topic]
	subs = append(subs, client)
	hub.subscribers[topic] = subs
	hub.log.Printf("Subscribers for %v: %v", topic, hub.subscribers[topic])
}

// Publish publishes a message to all subscribed clients
func (hub *Hub) Publish(message client.Message) {
	for _, subscriber := range hub.subscribers[message.Topic] {
		if subscriber != &message.Author {
			// A client shouldn't get it's own message
			subscriber.Messages <- message
		}
	}
}

// Unsubscribe unsubscribes a client from a topic
func (hub *Hub) Unsubscribe(topic string, client *client.Client) {
	// [TODO] Implement
}
