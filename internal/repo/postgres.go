package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const schema = `
 	CREATE TABLE IF NOT EXISTS strategies (
 	    id serial primary key,
 		strategy_addr text,
 		vault_addr text,
 		network text,
 		estimated_total_assets real,
		apy real,
 	    created_at timestamp
 	);
	CREATE TABLE IF NOT EXISTS vaults (
	    id serial primary key,
 		vault_addr text,
 		network text,
 		total_supply real,
 		price_per_share real,
 		created_at timestamp
 	);
	CREATE TABLE IF NOT EXISTS prices (
	    id serial primary key,
 		token text,
 		price real,
 		created_at timestamp
 	);
	CREATE TABLE IF NOT EXISTS balances (
	    id serial primary key,
 		vault_addr text,
 		network text,
 		user_addr text,
 		created_at timestamp
 	);
 	`

type DB struct {
	db  *sql.DB
	ctx context.Context
}

func (d *DB) Close() {
	d.db.Close()
}

func (d *DB) Ping() error {
	ctx, cancel := context.WithTimeout(d.ctx, 1*time.Second)
	defer cancel()
	err := d.db.PingContext(ctx)
	if err != nil {
		log.Print("ERROR: ", err, "\n")
		return fmt.Errorf("PingContext failed: %w", err)
	}
	return nil
}

func New(ctx context.Context, psqlConn string) *DB {
	db, err := sql.Open("pgx", psqlConn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(schema); err != nil {
		log.Fatal(err)
	}
	return &DB{
		db:  db,
		ctx: ctx,
	}
}
