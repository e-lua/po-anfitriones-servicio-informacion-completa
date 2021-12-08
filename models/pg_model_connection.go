package models

import (
	//pgx "github.com/jackc/pgx/v4"
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCN = Conectar_Pg_DB()

var (
	once_pg sync.Once
	p_pg    *pgxpool.Pool
)

func Conectar_Pg_DB() *pgxpool.Pool {

	once_pg.Do(func() {
		urlString := "postgres://postgresxd:postgresxd@161.35.226.104:7000/postgresxd?pool_max_conns=15"
		config, _ := pgxpool.ParseConfig(urlString)
		p_pg, _ = pgxpool.ConnectConfig(context.Background(), config)
	})

	return p_pg
}
