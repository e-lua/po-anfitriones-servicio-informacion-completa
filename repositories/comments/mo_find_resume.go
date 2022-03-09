package comments

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Traeremos los tips de anfitriones
func Mo_Find_Resume(idbusiness int) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("comment")

	var resultado []interface{}

	/*Condiciones*/
	datacomments := []bson.M{
		{
			"$match": bson.D{{Key: "idbusiness", Value: idbusiness}},
		},
		{
			"$match": bson.D{{Key: "isvisible", Value: true}},
		},
		{
			"$group": bson.M{
				"_id":         "",
				"avgstars":    bson.M{"$avg": "$stars"},
				"qtycomments": bson.M{"$sum": 1},
			},
		},
	}

	cursor, err := col.Aggregate(ctx, datacomments)
	if err != nil {
		return resultado, err
	}

	if err = cursor.All(ctx, &resultado); err != nil {
		panic(err)
	}

	return resultado, nil
}
