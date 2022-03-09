package comments

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Mo_Add_Business(input_business models.Mo_Business_Reported) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business_reported")

	_, err := col.InsertOne(ctx, input_business)
	if err != nil {
		return err
	}

	return nil
}
