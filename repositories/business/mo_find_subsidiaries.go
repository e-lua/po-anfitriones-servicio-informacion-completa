package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Find_Subsidiaries(idbusiness int) ([]*models.Mo_Subsidiary, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado []*models.Mo_Subsidiary

	condicion := bson.M{"idbusiness": idbusiness}

	opciones := options.Find()
	/*Indicar como ira ordenado*/
	opciones.SetSort(bson.D{{Key: "issubsidiary", Value: 1}})

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultado, err
	}

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Mo_Subsidiary
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	//Si todo esta bien
	return resultado, nil
}
