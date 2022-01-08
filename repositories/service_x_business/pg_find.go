package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Find(idbusiness int, idcountry int) ([]models.Pg_R_Service, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT r.idservice,r.name,bs.pricing,bs.typemoney,bs.isavailable FROM r_service AS r LEFT JOIN bussinessr_service AS bs ON bs.idservice=r.idservice WHERE bs.idbusiness=$1 UNION SELECT r.idservice,r.name,0,0,false FROM r_service AS r LEFT JOIN bussinessr_service AS bs ON bs.idservice=r.idservice LEFT JOIN r_countryr_service AS rr ON rr.idservice=r.idservice WHERE r.idservice NOT IN (SELECT bs.idservice FROM bussinessr_service AS bs WHERE bs.idbusiness=$1)AND rr.idcountry=$2"
	rows, error_show := db.Query(context.Background(), q, idbusiness, idcountry)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListPg_Service []models.Pg_R_Service

	if error_show != nil {
		return oListPg_Service, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var service models.Pg_R_Service
		rows.Scan(&service.IDservice, &service.Name, &service.Pricing, &service.TypeMoney, &service.IsAvailable)
		oListPg_Service = append(oListPg_Service, service)
	}

	//Si todo esta bien
	return oListPg_Service, nil

}
