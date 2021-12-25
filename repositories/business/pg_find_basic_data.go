package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Find_BasicData(idbusiness int) (models.Pg_BasicData, error) {

	var basic_data models.Pg_BasicData

	db := models.Conectar_Pg_DB()
	q := "SELECT b.name,b.timezone,CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END FROM business as b JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND b.idbusiness=$1"
	error_shown := db.QueryRow(context.Background(), q, idbusiness).Scan(&basic_data.Name, &basic_data.TimeZone, &basic_data.IsOpen)

	if error_shown != nil {
		return basic_data, error_shown
	}

	//Si todo esta bien
	return basic_data, nil
}
