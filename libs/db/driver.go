package db

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

type Drivers struct {
	Master, Replica dialect.Driver
}

var _ dialect.Driver = (*Drivers)(nil)

func (d *Drivers) Query(ctx context.Context, query string, args, v interface{}) error {
	return d.Replica.Query(ctx, query, args, v)
}

func (d *Drivers) Exec(ctx context.Context, query string, args, v interface{}) error {
	return d.Master.Exec(ctx, query, args, v)
}

func (d *Drivers) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.Master.Tx(ctx)
}

func (d *Drivers) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	return d.Master.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
}

func (d *Drivers) Dialect() string {
	return dialect.MySQL
}

func (d *Drivers) Close() error {
	werr := d.Master.Close()
	if werr != nil {
		return werr
	}
	rerr := d.Replica.Close()
	if rerr != nil {
		return rerr
	}
	return nil
}
