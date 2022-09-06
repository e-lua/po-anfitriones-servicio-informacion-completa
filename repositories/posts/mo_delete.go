package posts

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Delete(idbusiness int, idpost string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("post")

	objID, _ := primitive.ObjectIDFromHex(idpost)

	condicion := bson.M{
		"_id":        objID,
		"idbusiness": idbusiness,
	}

	_, err := col.DeleteOne(ctx, condicion)
	if err != nil {
		return err
	}

	return nil
}
