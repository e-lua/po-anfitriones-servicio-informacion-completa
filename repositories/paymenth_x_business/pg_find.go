package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Find(idbusiness int, idcountry int) ([]models.Pg_R_PaymentMethod, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT r.idpayment,r.name,r.urlphoto,r.hasnumber,coalesce(bp.phonenumber,''),bp.isavailable from r_paymentmethod AS r LEFT JOIN business_r_paymenth AS bp ON r.idpayment=bp.idpayment WHERE bp.idbusiness=$1 UNION SELECT r.idpayment,r.name,r.urlphoto,false,'',false from r_paymentmethod AS r LEFT JOIN business_r_paymenth AS bp ON r.idpayment=bp.idpayment LEFT JOIN r_countryr_payment AS rr ON rr.idpayment=r.idpayment WHERE r.idpayment NOT IN (SELECT bp.idpayment FROM business_r_paymenth AS bp WHERE bp.idbusiness=$1) AND rr.idcountry=$2"
	rows, error_show := db.Query(context.Background(), q, idbusiness, idcountry)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListPg_Paymenth []models.Pg_R_PaymentMethod

	if error_show != nil {
		return oListPg_Paymenth, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var paymenth models.Pg_R_PaymentMethod
		rows.Scan(&paymenth.IDPaymenth, &paymenth.Name, &paymenth.Url, &paymenth.HasNumber, &paymenth.PhoneNumber, &paymenth.IsAvailable)
		oListPg_Paymenth = append(oListPg_Paymenth, paymenth)
	}

	//Si todo esta bien
	return oListPg_Paymenth, nil

}
