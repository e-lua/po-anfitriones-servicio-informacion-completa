package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Find(idbusiness int, idcountry int) ([]models.Pg_R_PaymentMethod, error) {

	db := models.Conectar_Pg_DB()
	q := "select DISTINCT ON(p.idpayment)p.idpayment,p.name,p.urlphoto,p.hasnumber,bp.phonenumber,coalesce(bp.isavailable,false) from r_paymentmethod p LEFT JOIN business_r_paymenth bp ON p.idpayment=bp.idpayment LEFT JOIN r_countryr_payment rp ON p.idpayment=rp.idpayment WHERE bp.idbusiness<>$1 OR p.isavailable=false AND rp.idcountry=$2"
	rows, error_show := db.Query(context.Background(), q, idbusiness, idcountry)

	//.Scan(&typeF_x_Business.IDTypeFood, &typeF_x_Business.NameFood, &typeF_x_Business.URLPhoto, &typeF_x_Business.Weight)

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
