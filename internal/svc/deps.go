// DO NOT EDIT! This file is generated automatically by pet deps.
package svc

import (
	"context"
)

type (
	// Deps provides dependencies interface for groups of handlers.
	Deps interface {
		Primetime() PrimetimeGroup
	}

	// PrimetimeGroup provides interface for handlers of Primetime group.
	PrimetimeGroup interface {
		List(context.Context, RangeQuery) ([]PrimeTime, error)
	}
	// PrimeTime provides DTO interface for PrimeTime type.
	PrimeTime interface {
		GetHour() uint8
		GetCount() int
	}
	// Message provides DTO interface for Message type.
	Message interface {
		GetId() int
		GetDirection() string
		GetSent() string
		GetMailboxId() int
		GetCrmRid() int
	}
	// PrimeCache provides DTO interface for PrimeCache type.
	PrimeCache interface {
		GetHour() uint8
		GetCount() int
	}
	// RangeQuery provides DTO interface for RangeQuery type.
	RangeQuery interface {
		GetBefore() string
		GetAfter() string
	}
)
