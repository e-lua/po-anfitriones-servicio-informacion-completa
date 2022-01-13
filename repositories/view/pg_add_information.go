package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Mo_Add_Information(view models.Mqtt_View_Information) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("viewinformation")

	_, error_update := col.InsertOne(ctx, view)

	if error_update != nil {
		return error_update
	}

	return nil
}
