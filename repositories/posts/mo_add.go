package posts

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Mo_Add(input_post models.Mo_Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("post")

	_, err := col.InsertOne(ctx, input_post)
	if err != nil {
		return err
	}

	return nil
}
