package client

import (
	"github.com/Jackevansevo/msghub/uuid"
)

// Client represents a base client
type Client struct {
	UUID     uuid.UUID
	Messages chan Message
}

// Message defines a message
type Message struct {
	Author Client
	Topic  string
	Body   string
}

// NewClient returns a new BaseClient
func NewClient() Client {
	uuid, _ := uuid.New()
	return Client{uuid, make(chan Message)}
}
