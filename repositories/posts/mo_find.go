package posts

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Traeremos los tips de anfitriones
func Mo_Find(idbusiness int, limit int64) ([]*models.Mo_Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("post")

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado []*models.Mo_Post

	condicion := bson.M{
		"idbusiness": idbusiness,
	}

	opciones := options.Find()
	/*Indicar como ira ordenado*/
	opciones.SetSort(bson.D{{Key: "dateregistered", Value: -1}})
	opciones.SetLimit(limit)

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultado, err
	}

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Mo_Post
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	return resultado, nil
}
