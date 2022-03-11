package comments

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Traeremos los tips de anfitriones
func Mo_Find_CommentComensal(idbusiness int, idcomensal int) (models.Mo_Comment_ComensalFound, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("comment")

	condicion := bson.M{
		"idbusiness": idbusiness,
		"idcomensal": idcomensal,
	}

	//Resultado de la query
	var resultado models.Mo_Comment_ComensalFound

	//Asignamos los datos del cursor
	err_find := col.FindOne(ctx, condicion).Decode(&resultado)

	if err_find != nil {
		return resultado, err_find
	}

	return resultado, nil
}
