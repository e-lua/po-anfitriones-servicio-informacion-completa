package posts

import (
	"context"
	"log"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Delete(idbusiness int, uuidpost string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("post")

	id, _ := primitive.ObjectIDFromHex(uuidpost)

	log.Println("IMPRIMIENDO EL ID,", uuidpost)
	log.Println("IMPRIMIENDO EL ID 2,", id)

	filter := bson.M{"uuid": uuidpost}

	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
