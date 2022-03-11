package comments

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Update(idcomment string) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("comment")

	updtString := bson.M{
		"$set": bson.M{
			"isvisible": false,
		},
	}

	objID, _ := primitive.ObjectIDFromHex(idcomment)
	filtro := bson.M{"_id": objID}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}

func Mo_Update_MainData(input_comment models.Mo_Comment) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("comment")

	updtString := bson.M{
		"$set": bson.M{
			"stars":            input_comment.Stars,
			"comment":          input_comment.Comment,
			"fullnamecomensal": input_comment.FullNameComensal,
			"phonecomensal":    input_comment.PhoneComensal,
			"dateregistered":   input_comment.Dateregistered,
			"isvisible":        input_comment.IsVisible,
		},
	}

	filtro := bson.M{
		"idbusiness": input_comment.IDBusiness,
		"idcomensal": input_comment.IDComensal,
	}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
