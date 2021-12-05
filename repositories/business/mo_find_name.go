package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Find_Name(idbusiness int) (string, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	condicion := bson.M{"idBusiness": idbusiness}

	//Resultado de la query
	var resultado models.Mo_Business

	//Asignamos los datos del cursor
	err_find := col.FindOne(ctx, condicion).Decode(&resultado)

	if err_find != nil {
		return resultado.Name, err_find
	}

	//Si todo esta bien
	return resultado.Name, nil
}
