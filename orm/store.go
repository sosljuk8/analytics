package orm

import (
	"context"
	"github.com/estatie/pet/pkg/ant"
	"github.com/sosljuk8/analytics/internal/svc"
	"github.com/sosljuk8/analytics/internal/types"
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

// Migrate runs the migration tool to create all schema resources.
func (s *Stores) Migrate() error {
	return s.ent.Schema.Create(context.Background())
}

func (s *Stores) Load(req svc.RangeQuery) ([]svc.PrimeTime, error) {
	// s.ent.Message.Query().All(context.Background())
	times := []svc.PrimeTime{
		&types.PrimeTime{
			Hour:  10,
			Count: 3,
		},
		&types.PrimeTime{
			Hour:  11,
			Count: 4,
		},
	}

	_, err := s.ent.Message.Create().
		SetSent(req.GetAfter()).
		SetDirection("in").
		SetCrmRid(123).
		SetMailboxId(5).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	messages, err := s.ent.Message.Query().
		Select("sent").
		Where(message.SentLTE(req.GetBefore())).
		Where(message.SentGTE(req.GetAfter())).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	for _, _ = range messages {
		times = append(times, &types.PrimeTime{
			Hour:  0,
			Count: 0,
		})
	}

	return times, nil
}

func (s *Stores) Close() error {
	return s.ent.Close()
}
