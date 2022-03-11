package comments

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Add(input_comment models.Mo_Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("comment")

	updtString := bson.M{
		"$set": bson.M{
			"stars":            input_comment.Stars,
			"comment":          input_comment.Comment,
			"idbusiness":       input_comment.IDBusiness,
			"idcomensal":       input_comment.Comment,
			"fullnamecomensal": input_comment.FullNameComensal,
			"fullnamebusiness": input_comment.FullNameBusiness,
			"phonecomensal":    input_comment.PhoneComensal,
			"dateregistered":   input_comment.Dateregistered,
			"isvisible":        input_comment.IsVisible,
		},
	}

	filtro := bson.M{
		"idbusiness": input_comment.IDBusiness,
		"idcomensal": input_comment.IDComensal,
	}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return err
	}

	return nil
}
