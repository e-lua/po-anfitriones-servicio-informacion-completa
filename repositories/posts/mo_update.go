package posts

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_UpdateImage(idbusiness int, idpost string, url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("post")

	objID, _ := primitive.ObjectIDFromHex(idpost)

	updtString := bson.M{
		"$set": bson.M{
			"url": url,
		},
	}

	filtro := bson.M{
		"_id":        objID,
		"idbusiness": idbusiness,
	}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return err
	}

	return nil
}
