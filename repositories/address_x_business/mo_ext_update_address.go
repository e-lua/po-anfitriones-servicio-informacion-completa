package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Ext_Update(intpu_mo_business models.Mo_Business, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN_Externo.Database("restoner_externo")
	col := db.Collection("business_cards")

	var anfitrion_mo models.Mo_Business_Cards
	anfitrion_mo.Location.GeoJSONType = "Point"
	anfitrion_mo.Location.Coordinates = []float64{intpu_mo_business.Address.Latitude, intpu_mo_business.Address.Longitude}

	updtString := bson.M{
		"$set": bson.M{
			"location": anfitrion_mo.Location,
		},
	}

	filtro := bson.M{"idbusiness": idbusiness}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
