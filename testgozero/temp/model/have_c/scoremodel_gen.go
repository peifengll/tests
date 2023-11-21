// Code generated by goctl. DO NOT EDIT.

package have_c

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	scoreFieldNames          = builder.RawFieldNames(&Score{})
	scoreRows                = strings.Join(scoreFieldNames, ",")
	scoreRowsExpectAutoSet   = strings.Join(stringx.Remove(scoreFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	scoreRowsWithPlaceHolder = strings.Join(stringx.Remove(scoreFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheScoreIdPrefix = "cache:score:id:"
)

type (
	scoreModel interface {
		Insert(ctx context.Context, data *Score) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Score, error)
		Update(ctx context.Context, data *Score) error
		Delete(ctx context.Context, id int64) error
	}

	defaultScoreModel struct {
		sqlc.CachedConn
		table string
	}

	Score struct {
		Id    int64          `db:"id"`
		StuId int64          `db:"stu_id"`
		CName sql.NullString `db:"c_name"`
		Grade sql.NullInt64  `db:"grade"`
	}
)

func newScoreModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultScoreModel {
	return &defaultScoreModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`score`",
	}
}

func (m *defaultScoreModel) withSession(session sqlx.Session) *defaultScoreModel {
	return &defaultScoreModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`score`",
	}
}

func (m *defaultScoreModel) Delete(ctx context.Context, id int64) error {
	scoreIdKey := fmt.Sprintf("%s%v", cacheScoreIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, scoreIdKey)
	return err
}

func (m *defaultScoreModel) FindOne(ctx context.Context, id int64) (*Score, error) {
	scoreIdKey := fmt.Sprintf("%s%v", cacheScoreIdPrefix, id)
	var resp Score
	err := m.QueryRowCtx(ctx, &resp, scoreIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", scoreRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultScoreModel) Insert(ctx context.Context, data *Score) (sql.Result, error) {
	scoreIdKey := fmt.Sprintf("%s%v", cacheScoreIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, scoreRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.StuId, data.CName, data.Grade)
	}, scoreIdKey)
	return ret, err
}

func (m *defaultScoreModel) Update(ctx context.Context, data *Score) error {
	scoreIdKey := fmt.Sprintf("%s%v", cacheScoreIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf( "update %s set %s where `id` = ?", m.table, scoreRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.StuId, data.CName, data.Grade, data.Id)
	}, scoreIdKey)
	return err
}

func (m *defaultScoreModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheScoreIdPrefix, primary)
}

func (m *defaultScoreModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", scoreRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultScoreModel) tableName() string {
	return m.table
}