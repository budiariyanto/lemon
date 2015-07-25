package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Postgresql struct {
	ConnectionString string
	pool             *sql.DB
}

func NewPostgresql(connectionString string) (*Postgresql, error) {
	postgresql := new(Postgresql)
	postgresql.ConnectionString = connectionString

	return postgresql, nil
}

func (psql *Postgresql) Connect() (*sql.DB, error) {
	if psql.pool == nil {
		pool, err := sql.Open("postgres", psql.ConnectionString)

		if err != nil {
			return nil, err
		}

		psql.pool = pool
		pool.SetMaxIdleConns(10)
	}

	return psql.pool, nil
}
