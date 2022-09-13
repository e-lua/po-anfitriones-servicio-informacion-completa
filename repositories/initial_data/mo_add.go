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

	var registro models.Mo_Registro_FromMqtt

	registro.IdBusiness = anfitrion_mo.IdBusiness
	registro.CreatedDate = anfitrion_mo.UpdatedDate
	registro.Available = true
	registro.OrdersRejected = 0

	_, error_add := col.InsertOne(ctx, registro)

	if error_add != nil {
		return error_add

	}

	return nil

}
