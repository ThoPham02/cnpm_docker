package svc

import (
	"cnpm-be/internal/config"
	"cnpm-be/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserTblModel(sqlx.NewMysql(c.Database.DataSource)),
	}
}
