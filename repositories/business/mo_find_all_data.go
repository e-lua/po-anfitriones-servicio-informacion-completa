package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Find_All_Data(idbusiness int) (models.Mo_Business, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	condicion := bson.M{"idbusiness": idbusiness}

	//Resultado de la query
	var resultado models.Mo_Business

	//Asignamos los datos del cursor
	err_find := col.FindOne(ctx, condicion).Decode(&resultado)

	if err_find != nil {
		return resultado, err_find
	}

	//Si todo esta bien
	return resultado, nil
}

func Web_Mo_Find_All_Data(uniquename string) (models.Mo_Business, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	condicion := bson.M{"uniquename": uniquename}

	//Resultado de la query
	var resultado models.Mo_Business

	//Asignamos los datos del cursor
	err_find := col.FindOne(ctx, condicion).Decode(&resultado)

	if err_find != nil {
		return resultado, err_find
	}

	//Si todo esta bien
	return resultado, nil
}
