package simple_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreatConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("conn_string")

	return pgx.Connect(ctx, connString)
}
