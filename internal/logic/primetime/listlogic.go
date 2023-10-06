package primetime

import (
	"context"
	"github.com/sosljuk8/analytics/internal/svc"
	"github.com/sosljuk8/analytics/internal/types"

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

func (l *ListLogic) List(req *types.RangeQuery) (resp []svc.PrimeTime, err error) {

	times, err := l.svcCtx.Store.Load(req)
	if err != nil {
		return nil, err
	}
	return times, nil
}
