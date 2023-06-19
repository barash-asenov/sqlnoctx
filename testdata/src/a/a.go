package a

import (
	"context"
	"database/sql"
)

func main() {
	db, _ := sql.Open("test", "test")

	db.Exec("select * from users")                              // want `should use ExecContext instead`
	db.ExecContext(context.Background(), "select * from users") // OK

	db.Ping()                            // want `should use PingContext instead`
	db.PingContext(context.Background()) // OK

	db.Prepare("select * from users")                              // want `should use PrepareContext instead`
	db.PrepareContext(context.Background(), "select * from users") // OK

	db.Query("select * from users")                              // want `should use QueryContext instead`
	db.QueryContext(context.Background(), "select * from users") // OK

	db.QueryRow("select * from users")                              // want `should use QueryRowContext instead`
	db.QueryRowContext(context.Background(), "select * from users") // OK
}
