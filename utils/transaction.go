package utils

import (
	"context"
	"database/sql"
)

type txnKey string

const transaction_key txnKey = "transaction_key"

func GetTransactionFromContext(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(transaction_key).(*sql.Tx)
	return tx, ok
}
func SetTransactionInContext(ctx context.Context, tx *sql.Tx) context.Context {
	ctx = context.WithValue(ctx, transaction_key, tx)
	return ctx
}
