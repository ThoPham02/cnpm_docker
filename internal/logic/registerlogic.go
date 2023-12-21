package logic

import (
	"context"
	"database/sql"
	"time"

	"cnpm-be/internal/svc"
	"cnpm-be/internal/types"
	"cnpm-be/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	l.Logger.Info("Register ", req)

	var user *model.UserTbl

	var userInsert = &model.UserTbl{
		Id:       time.Now().UnixMilli(),
		Fullname: sql.NullString{Valid: true, String: req.Fullname},
		Email:    req.Email,
		HashPw:   sql.NullString{Valid: true, String: req.Password},
	}

	user, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    1,
				Message: "DB Error",
			},
		}, nil
	}
	if user != nil {
		l.Logger.Info("Email was existed")
		return &types.RegisterRes{
			Result: types.Result{
				Code:    4,
				Message: "Email was existed",
			},
		}, nil
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, userInsert)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    1,
				Message: "DB Error",
			},
		}, nil
	}

	resp = &types.RegisterRes{
		Result: types.Result{
			Code:    0,
			Message: "Register Success",
		},
		User: types.User{
			ID:       userInsert.Id,
			FullName: userInsert.Fullname.String,
			Email:    userInsert.Email,
		},
	}

	l.Logger.Info("Register Success", resp)
	return
}
