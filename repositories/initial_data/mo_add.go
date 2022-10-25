package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Mo_Add(anfitrion_mo models.Mo_BusinessWorker_Mqtt) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	var business models.Mo_Business
	var banner models.Mo_Banner
	banner.IdBanner = 1
	banner.UrlImage = "https://restoner-public-space.sfo3.cdn.digitaloceanspaces.com/restoner-general/default-image/Portada_defect.png"
	business.IdBusiness = anfitrion_mo.IdBusiness
	business.CreatedDate = anfitrion_mo.UpdatedDate
	business.Available = true
	business.OrdersRejected = 0
	business.IsSubsidiary = anfitrion_mo.IsSubsidiary
	business.SubsidiaryOf = anfitrion_mo.SubsidiaryOf
	business.Banner = append(business.Banner, banner)

	_, error_add := col.InsertOne(ctx, business)

	if error_add != nil {
		return error_add

	}

	return nil

}
