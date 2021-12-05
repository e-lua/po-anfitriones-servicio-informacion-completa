package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"github.com/jackc/pgx/v4"
)

func Pg_Add(paymentMethods []models.Mo_PaymenthMeth, idbusiness int) (int64, error) {

	db := models.Conectar_Pg_DB()

	//Eliminamos los datos
	q := "DELETE FROM Business_R_Paymenth WHERE idbusiness=$1"
	_, err_add := db.Exec(context.Background(), q, idbusiness)

	if err_add != nil {

		defer db.Close()
		return 0, err_add
	}

	/*type RowSrc interface  {
		Next() bool
		Values() ([]interface{}, error)
		Err() error
	}


	value := make([]interface{}, 3)

	for _, payment_x_business := range paymentMethods {

		value[0] = idbusiness
		value[1] = payment_x_business.IDPaymenth
		value[2] = true
	}

	rowSrc := RowSrc{
		value,
	}*/

	var list_p_x_b [][]interface{}

	for _, payment_x_business := range paymentMethods {
		var p_x_b []interface{}
		p_x_b[0] = idbusiness
		p_x_b[1] = payment_x_business.IDPaymenth
		p_x_b[2] = true
		list_p_x_b = append(list_p_x_b, p_x_b)
	}

	//Insertamos los datos
	//q_2 := "INSERT INTO Business_R_Paymenth(idbusiness,idPayment,isavailable) VALUES ($1,$2,$3)"
	add_paymenth, err_add := db.CopyFrom(context.Background(), pgx.Identifier{"business_r_paymenth"}, []string{"idbusiness", "idpayment", "isavailable"}, pgx.CopyFromRows(list_p_x_b))

	if err_add != nil {
		defer db.Close()
		return add_paymenth, err_add
	}

	defer db.Close()
	return add_paymenth, nil
}
