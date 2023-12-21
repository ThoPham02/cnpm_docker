package logic

import (
	"context"

	"cnpm-be/internal/svc"
	"cnpm-be/internal/types"
	"cnpm-be/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	l.Logger.Info("Login ", req)

	var user *model.UserTbl

	user, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    1,
				Message: "DB Error",
			},
		}, nil
	}
	if user == nil {
		l.Logger.Info("User not found")
		return &types.LoginRes{
			Result: types.Result{
				Code:    2,
				Message: "User not found",
			},
		}, nil
	}

	if user.HashPw.String != req.Password {
		l.Logger.Info("PassWord is wrong")
		return &types.LoginRes{
			Result: types.Result{
				Code:    3,
				Message: "Password is wrong",
			},
		}, nil
	}

	resp = &types.LoginRes{
		Result: types.Result{
			Code:    0,
			Message: "Login Success",
		},
		User: types.User{
			ID:       user.Id,
			FullName: user.Fullname.String,
			Email:    user.Email,
		},
	}

	l.Logger.Info("Login Success", resp)
	return
}
