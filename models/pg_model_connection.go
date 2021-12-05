package models

import (
	//pgx "github.com/jackc/pgx/v4"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

/*
const (
	host      = "161.35.226.104"
	port      = "7000"
	user      = "postgresxd"
	password  = "postgresxd"
	dbname_pg = "postgresxd"
)*/

var PostgresCN = Conectar_Pg_DB()

/*
func Conectar_Pg_DB() *sql.DB {

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname_pg, password, port)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Error en el servidor interno en el driver de PostgreSQL, mayor detalle: " + err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error en el servidor interno al intentar conectarse con la base de datos, mayor detalle: " + err.Error())
	}

	return db
}*/

func Conectar_Pg_DB() *pgxpool.Pool {

	urlString := "postgres://postgresxd:postgresxd@161.35.226.104:7000/postgresxd?pool_max_conns=20"

	config, error_connec_pg := pgxpool.ParseConfig(urlString)

	if error_connec_pg != nil {
		log.Fatal("Error en el servidor interno en el driver de PostgreSQL, mayor detalle: " + error_connec_pg.Error())
		return nil
	}

	conn, _ := pgxpool.ConnectConfig(context.Background(), config)

	return conn
}

//ChequeoConnection es el Ping a la BD
/*func ChequeoConnection_Pg() int {

	err := PostgresCN.Ping()
	if err != nil {
		return 0
	}
	return 1

}*/
