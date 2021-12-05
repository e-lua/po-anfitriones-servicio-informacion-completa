package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Add(services []models.Mo_Service, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	//Eliminamos los datos
	q := "DELETE FROM BussinessR_Service WHERE idbusiness=$1"
	_, err_add := db.Exec(context.Background(), q, idbusiness)

	if err_add != nil {

		defer db.Close()
		return err_add
	}

	//Insertamos los datos
	/*	q_2 := "INSERT INTO BussinessR_Service(idbusiness,idService,isavailable) VALUES ($1,$2,$3)"
		add_service, err_add := db.Prepare(q_2)

		if err_add != nil {
			defer db.Close()
			return err_add
		}

		for _, service_x_business := range services {
			add_service.Exec(idbusiness, service_x_business.IDService, true)
		}*/

	defer db.Close()
	return nil
}
