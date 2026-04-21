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

	// if err := simple_sql.InsertRow(
	// 	ctx,
	// 	conn,
	// 	"Обед3",
	// 	"Покушать борщика",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	if err := simple_sql.UpdateRaw(ctx, conn); err != nil {
		panic(err)
	}

	if err := simple_sql.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("succeed!")
}
