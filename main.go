package main

import (
	"GolangPostgres/feature_postgre/simple_connection"
	"GolangPostgres/feature_postgre/simple_sql"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreatConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreatTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	fmt.Println(tasks)

	fmt.Println("succeed!")
}
