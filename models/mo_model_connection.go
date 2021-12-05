package models

import (
	/* Context: es un espacio en memoria donde podre ir a compartiendo, setear un
	   conexto de ejecución, por ejemplo, que una ejecución no supere los 15 seg, esto
	   para evitar los cuelgues. Por lo tanto, nos sirve para comunicar información
	   entre ejcuciones y además nos permirte setear una serie de valores com un Timeout.
	   Se tiene que ejcutar en 15 segudos, por lo que si ocurre un cuelgue en la BD, no se
	   colgará la API, ya que esto antes causaba que todas las API se cuelguen.
	*/
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN objetivo de conexion a la BD
var MongoCN = ConectarBD_Mo()

//Con options seteo la URL de la base de datos || "c" minuscula = será de uso interno
var clientOptions = options.Client().ApplyURI("mongodb://mongodbbusiness_user:mongodb1151@mongo:27017")

// ConectarBD: Se conecta a la base de datos, toma la conexión de clientOptions
func ConectarBD_Mo() *mongo.Client {
	//TODO crea sin un timeout
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printf("Conexion exitosa con la BD Mo")
	return client
}

//ChequeoConnection es el Ping a la BD
func ChequeoConnection_Mo() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}
