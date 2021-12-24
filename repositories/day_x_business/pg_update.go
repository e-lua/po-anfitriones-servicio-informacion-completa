package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Update(input_mo_business models.Mo_Business, idbusiness int) error {

	//Instanciamos los datos
	idday_pg, idbusiness_pg, starttime_pg, endtime_pg, available_pg := []int{}, []int{}, []string{}, []string{}, []bool{}

	//Convertimos a formato 24 horas
	for _, day := range input_mo_business.DailySchedule {

		//Convertimos la hora
		startTime, _ := time.Parse("15:04 PM", day.StarTime)
		endTime, _ := time.Parse("15:04 PM", day.EndTime)

		//Append
		idday_pg = append(idday_pg, day.IDDia)
		idbusiness_pg = append(idbusiness_pg, idbusiness)
		starttime_pg = append(starttime_pg, startTime.Format("15:04"))
		endtime_pg = append(endtime_pg, endTime.Format("15:04"))
		available_pg = append(available_pg, day.IsAvaiable)
	}

	//Conexion con la BD
	db := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db.Begin(context.Background())
	if error_tx != nil {
		return error_tx
	}

	//ELIMINAR DATOS
	q_delete := `DELETE FROM BusinessSchedule WHERE idbusiness=$1`
	_, err := tx.Exec(context.Background(), q_delete, idbusiness)
	if err != nil {
		return err
	}

	//HORARIO
	q_schedulerange := `INSERT INTO BusinessSchedule(idschedule,idbusiness,starttime,endtime,available,zonetime) (SELECT * FROM unnest($1::int[],$2::int[],$3::varchar(14)[],$4::varchar(14)[],$5::boolean[]));`
	if _, err_schedule := tx.Exec(context.Background(), q_schedulerange, idday_pg, idbusiness_pg, starttime_pg, endtime_pg, available_pg); err_schedule != nil {
		return err_schedule
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(context.Background())
	if err_commit != nil {
		return err_commit
	}
	return nil
}
