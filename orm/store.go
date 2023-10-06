package orm

import (
	"context"
	"github.com/estatie/pet/pkg/ant"
	"github.com/sosljuk8/analytics/internal/svc"
	"github.com/sosljuk8/analytics/orm/ent"
	"github.com/sosljuk8/analytics/orm/ent/message"
)

type Stores struct {
	Messages *MessageStore
	ent      *ent.Client
}

func NewStore(config ant.Config) *Stores {
	client := ant.MustClient(config, ent.Open)
	return &Stores{
		Messages: NewMessageStore(client),
		ent:      client,
	}
}

func (s *Stores) Load(req svc.RangeQuery) ([]svc.PrimeTime, error) {
	// s.ent.Message.Query().All(context.Background())
	times := make([]svc.PrimeTime, 0)

	messages, err := s.ent.Message.Query().
		Select("sent").
		Where(message.SentGTE(req.GetStart())).
		Where(message.SentLTE(req.GetFinish())).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	for _, message := range messages {
		//times = append(times, svc.PrimeTime{Time: message.Sent})
	}

	return times, nil
}

func (s *Stores) Close() error {
	return s.ent.Close()
}
