package orm

import (
	"context"
	"github.com/sosljuk8/analytics/orm/ent"
)

type MessageStore struct {
	ent *ent.Client
}

func NewMessageStore(client *ent.Client) *MessageStore {
	return &MessageStore{
		ent: client,
	}
}

// CreateInitialMessage creates a new message in the database
func (r *MessageStore) CreateInitialMessage() (*ent.Message, error) {
	return r.ent.Message.Create().
		SetDirection("in").
		SetCrmRid(0).
		SetMailboxId(0).
		SetSent("2020-01-20 16:22:29").
		Save(context.Background())
}

// LoadAllMessages loads all messages from the database
func (r *MessageStore) LoadAllMessages() ([]*ent.Message, error) {
	return r.ent.Message.Query().All(context.Background())
}
