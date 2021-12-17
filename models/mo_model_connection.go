package models

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN objetivo de conexion a la BD
var MongoCN = ConectarBD_Mo()

var (
	once_mo sync.Once
	client  *mongo.Client
)

//Con options seteo la URL de la base de datos || "c" minuscula = será de uso interno
var clientOptions = options.Client().ApplyURI("mongodb://mongodbbusiness_user:mongodb1151@mongo:27017")

// ConectarBD: Se conecta a la base de datos, toma la conexión de clientOptions
func ConectarBD_Mo() *mongo.Client {

	once_mo.Do(func() {
		//TODO crea sin un timeout
		client, _ = mongo.Connect(context.TODO(), clientOptions)

		log.Printf("Conexion exitosa con la BD Mo")
	})

	return client
}

//MongoCN objetivo de conexion a la BD
var MongoCN_Externo = ConectarBD_Mo_Externo()

//Con options seteo la URL de la base de datos || "c" minuscula = será de uso interno
var clientOptions_externo = options.Client().ApplyURI("mongodb://mongodbbusine55_user:mongodb6921@mongo:45045")

// ConectarBD: Se conecta a la base de datos, toma la conexión de clientOptions
func ConectarBD_Mo_Externo() *mongo.Client {

	once_mo.Do(func() {
		//TODO crea sin un timeout
		client, _ = mongo.Connect(context.TODO(), clientOptions_externo)

		log.Printf("Conexion exitosa con la BD Mo")
	})

	return client
}
