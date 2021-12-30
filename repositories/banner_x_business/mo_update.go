package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Update(idbanner int, urlphoto string, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	//Todo sobre el banner
	var banners []models.Mo_Banner
	var banner models.Mo_Banner
	banner.IdBanner = idbanner
	banner.UrlImage = urlphoto
	banners = append(banners, banner)

	//Ponemos el banner en el business
	var business models.Mo_Business

	business.Banner = banners

	updtString := bson.M{
		"$set": bson.M{
			"banners": banners,
		},
	}

	filtro := bson.M{"idbusiness": idbusiness}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
