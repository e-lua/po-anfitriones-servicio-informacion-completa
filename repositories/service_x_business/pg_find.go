package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Find(idbusiness int, idcountry int) ([]models.Pg_R_Service, error) {

	db := models.Conectar_Pg_DB()
	q := "select DISTINCT ON(s.idservice)s.idservice,s.name,s.urlphoto,coalesce(bs.isavailable,false) from r_service s LEFT JOIN bussinessr_service bs ON s.idservice=bs.idservice LEFT JOIN r_countryr_service rs ON s.idservice=rs.idservice WHERE bs.idbusiness<>$1 OR s.isavailable=false AND rs.idcountry=$2"
	rows, error_show := db.Query(context.Background(), q, idbusiness, idcountry)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListPg_Service []models.Pg_R_Service

	if error_show != nil {
		defer db.Close()
		return oListPg_Service, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var service models.Pg_R_Service
		rows.Scan(&service.IDservice, &service.Name, &service.Url, &service.IsAvailable)
		oListPg_Service = append(oListPg_Service, service)
	}

	defer db.Close()

	//Si todo esta bien
	return oListPg_Service, nil

}
