package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	*sqlx.DB
}

func NewSqlClient() *PostgresClient {
	connStr := "postgres://vctciqprbgdjtl:679973ef9de0ff3e5fa3c790c4d93ad56a1b8e6b872317dd42a712cf7f857976@ec2-54-173-138-144.compute-1.amazonaws.com/d4suet4hvbb9vh"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln("[DB Connection Error]", err)
	}

	return &PostgresClient{db}
}
