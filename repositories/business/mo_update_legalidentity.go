package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Update_Legalidentity(inputserialize_legalidentity models.Mqtt_LegalIdentity) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	updtString := bson.M{
		"$set": bson.M{
			"legalidentity":   inputserialize_legalidentity.LegalIdentity,
			"typesuscription": inputserialize_legalidentity.TypeSuscription,
			"iva":             inputserialize_legalidentity.IVA,
		},
	}

	filtro := bson.M{"idbusiness": inputserialize_legalidentity.IdBusiness}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
