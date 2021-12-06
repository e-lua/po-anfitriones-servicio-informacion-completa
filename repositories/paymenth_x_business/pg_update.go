package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

func Pg_Add(input_mo_business models.Mo_Business, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	//Eliminamos los datos
	/*q := "DELETE FROM Business_R_Paymenth WHERE idbusiness=$1"
	_, err_add := db.Exec(context.Background(), q, idbusiness)

	if err_add != nil {

		defer db.Close()
		return 0, err_add
	}*/

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

	idbusiness_pg, idpaymenth_pg, isavailable_pg := []int{}, []int{}, []bool{}
	for _, v := range input_mo_business.PaymentMethods {
		if v.IsAvaiable {
			idbusiness_pg = append(idbusiness_pg, idbusiness)
			idpaymenth_pg = append(idpaymenth_pg, v.IDPaymenth)
			isavailable_pg = append(isavailable_pg, true)
		}
	}
	query := `INSERT INTO Business_R_Paymenth(idbusiness,idPayment,isavailable) (select * from unnest($1::int[], $2::int[],$3::boolean[]))`
	if _, err := db.Exec(context.Background(), query, idbusiness_pg, idpaymenth_pg, isavailable_pg); err != nil {
		return err
	}

	defer db.Close()
	return nil
}
