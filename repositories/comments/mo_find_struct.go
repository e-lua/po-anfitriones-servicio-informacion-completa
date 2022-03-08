package comments

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Traeremos los tips de anfitriones
func Mo_Find_Struct(idbusiness int, pagina int64) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("comment")

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado []interface{}

	/*Condiciones*/
	datacomments := []bson.M{
		{
			"idbusiness": idbusiness,
		},
		{
			"$group": bson.M{
				"_id":        "",
				"avgstars":   bson.M{"$avg": "$stars"},
				"qtycomment": bson.M{"$sum": 1},
			},
		},
	}

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Aggregate(ctx, datacomments)
	if err != nil {
		return resultado, err
	}
	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro interface{}
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	return resultado, nil
}
