// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"

	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	//Example...
	// In production: query from DB using l.svcCtx.DB
	return &types.UserInfoResp{
		Id:       req.Id,
		Username: gofakeit.Username(),
		Email:    gofakeit.Email(),
	}, nil
}
