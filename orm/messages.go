package orm

import "github.com/sosljuk8/analytics/orm/ent"

type MessageStore struct {
	client *ent.Client
}

func NewMessageStore(client *ent.Client) *MessageStore {
	return &MessageStore{
		client: client,
	}
}
