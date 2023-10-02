package primetime

import (
	"context"
	"github.com/sosljuk8/analytics/internal/svc"
	"github.com/sosljuk8/analytics/internal/types"
	"math/rand"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.RangeQuery) (resp []types.PrimeTime, err error) {
	times := make([]types.PrimeTime, 10)
	for i := 0; i < 10; i++ {
		times[i] = types.PrimeTime{
			Hour:  uint8(rand.Intn(64)),
			Count: rand.Intn(1000),
		}
	}

	return times, nil
}
