package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Mo_Ext_Add(name string, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN_Externo.Database("restoner_externo")
	col := db.Collection("business_cards")

	var registro_busiss_ext models.Mo_Business_Cards

	registro_busiss_ext.IDBusiness = idbusiness
	registro_busiss_ext.Available = true
	registro_busiss_ext.OrdersRejected = 0
	registro_busiss_ext.IsOpen = false
	registro_busiss_ext.Name = name

	_, error_add := col.InsertOne(ctx, registro_busiss_ext)

	if error_add != nil {
		return error_add

	}

	return nil
}
