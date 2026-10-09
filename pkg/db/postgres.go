package db

import (
	"database/sql"
	"edu-crm/config"
	"fmt"
	_"github.com/lib/pq"

	
)

func ConnectDb(cfg config.Config) (*sql.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	    cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresDb,
    )
	conn, err := sql.Open("postgres", psqlString)
	if err != nil {
		fmt.Printf("could not open db connection: %v\n", err)
		return nil, err
	}
	return conn, nil
}
