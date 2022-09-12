package posts

import (
	"context"
	"log"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Delete(idbusiness int, idpost string) error {

	log.Println("PRINT 3", idpost)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("post")

	log.Println("PRINT 4", idpost)

	id, _ := primitive.ObjectIDFromHex(idpost)

	filter := bson.M{"_id": id}

	log.Println("PRINT 5", idpost)

	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
