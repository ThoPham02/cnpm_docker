package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserTblModel = (*customUserTblModel)(nil)

type (
	// UserTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTblModel.
	UserTblModel interface {
		userTblModel
		withSession(session sqlx.Session) UserTblModel
		FindOneByEmail(ctx context.Context, email string) (*UserTbl, error)
	}

	customUserTblModel struct {
		*defaultUserTblModel
	}
)

// NewUserTblModel returns a model for the database table.
func NewUserTblModel(conn sqlx.SqlConn) UserTblModel {
	return &customUserTblModel{
		defaultUserTblModel: newUserTblModel(conn),
	}
}

func (m *customUserTblModel) withSession(session sqlx.Session) UserTblModel {
	return NewUserTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultUserTblModel) FindOneByEmail(ctx context.Context, email string) (*UserTbl, error) {
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userTblRows, m.table)
	var resp UserTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
