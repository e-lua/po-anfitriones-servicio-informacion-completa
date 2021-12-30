package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Update(banner models.Mo_BusinessBanner_Mqtt) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	var banner_business models.Mo_Business
	banner_business.Banner = append(banner_business.Banner, banner.Banner)

	updtString := bson.M{
		"$set": bson.M{
			"banners": banner_business.Banner,
		},
	}

	filtro := bson.M{"idbusiness": banner.IDBusiness}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
