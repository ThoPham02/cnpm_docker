package logic

import (
	"context"

	"cnpm-be/internal/svc"
	"cnpm-be/internal/types"
	"cnpm-be/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserRes, err error) {
	l.Logger.Info("Get User ", req)

	var user *model.UserTbl

	user, err = l.svcCtx.UserModel.FindOne(l.ctx, req.UserID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetUserRes{
			Result: types.Result{
				Code:    1,
				Message: "DB Error",
			},
		}, nil
	}
	if user == nil {
		l.Logger.Info("User not found")
		return &types.GetUserRes{
			Result: types.Result{
				Code:    2,
				Message: "User not found",
			},
		}, nil
	}

	resp = &types.GetUserRes{
		Result: types.Result{
			Code:    0,
			Message: "Get User Success",
		},
		User: types.User{
			ID:       user.Id,
			FullName: user.Fullname.String,
			Email:    user.Email,
		},
	}

	l.Logger.Info("Get User Success", resp)
	return
}
