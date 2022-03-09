package comments

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Mo_Add_Comment(input_comment models.Mo_Comment_Reported) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("comments_reported")

	_, err := col.InsertOne(ctx, input_comment)
	if err != nil {
		return err
	}

	return nil
}
