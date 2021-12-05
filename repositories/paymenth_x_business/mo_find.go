package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Find(idbusiness int) ([]models.Mo_PaymenthMeth, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"idbusiness": idbusiness})
	condiciones = append(condiciones, bson.M{"$unwind": "$paymentMethods"})

	cursor, _ := col.Aggregate(ctx, condiciones)

	//Resultado de la query
	var resultado []models.Mo_PaymenthMeth

	//Asignamos los datos del cursor
	err_cursor_add := cursor.All(ctx, &resultado)
	if err_cursor_add != nil {
		return resultado, err_cursor_add
	}

	//Si todo esta bien
	return resultado, nil
}
