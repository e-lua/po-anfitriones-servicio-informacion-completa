package models

import (
	"context"

	pgx "github.com/jackc/pgx/v4"
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

func Conectar_Pg_DB() *pgx.Conn {

	urlString := "postgres://postgresxd:postgresxd@161.35.226.104:7000/postgresxd?pool_max_conns=20"

	conn, _ := pgx.Connect(context.Background(), urlString)

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
