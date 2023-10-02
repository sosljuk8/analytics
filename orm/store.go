package orm

import (
	"github.com/estatie/pet/pkg/ant"
	"github.com/sosljuk8/analytics/orm/ent"
)

type Store struct {
	Messages *MessageStore
	ent      *ent.Client
}

func NewStore(config ant.Config) *Store {
	client := ant.MustClient(config, ent.Open)
	return &Store{
		Messages: NewMessageStore(client),
		ent:      client,
	}
}

func (s *Store) Close() error {
	return s.ent.Close()
}
