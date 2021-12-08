package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Add_IntialiData(anfitrionpg models.Mo_BusinessWorker_Mqtt) error {

	db := models.Conectar_Pg_DB()

	//Agregamos el Business

	_, err_add_business := db.Exec(context.Background(), "INSERT INTO Business(idbusiness,idcountry,createdDate,isopen) VALUES ($1,$2,$3,$4) RETURNING idbusiness", anfitrionpg.IdBusiness, anfitrionpg.IdCountry, anfitrionpg.UpdatedDate, false)
	if err_add_business != nil {
		return err_add_business
	}

	return nil
}
