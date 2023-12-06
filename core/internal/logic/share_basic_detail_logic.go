package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetialReply, err error) {
	// 对分享记录的点击次数加1
	_, err = l.svcCtx.Engine.Exec("update share_basic set click_num = click_num + 1 where identity = ?", req.Identity)
	if err != nil {
		return nil, err
	}
	// 获取资源的详细信息
	resp = &types.ShareBasicDetialReply{}
	l.svcCtx.Engine.Table("share_basic").
		Select("share_basic.repository_identity, repository_pool.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Where("share_basic.identity = ?", req.Identity).Get(resp)

	return
}
