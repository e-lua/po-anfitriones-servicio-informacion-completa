package recover

import (
	//REPOSITORIES
	business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/business"
)

/*----------------------TRAEMOS LOS DATOS----------------------*/

func RecoverAll_Service(business_all []interface{}) (int, bool, string, string) {

	//Buscamos si ya he visitado negocios antes
	error_insert_all := business_repository.Mo_Add_Many(business_all)

	if error_insert_all != nil {
		return 500, true, "Error interno en el servidor al intentar registrar los datos recuperados, detalle: " + error_insert_all.Error(), ""
	}

	return 200, false, "", "Datos recuperados registrados correctamente"
}

func RecoverOne_Service(business_one interface{}) (int, bool, string, string) {

	//Buscamos si ya he visitado negocios antes
	error_insert_all := business_repository.Mo_Add(business_one)

	if error_insert_all != nil {
		return 500, true, "Error interno en el servidor al intentar registrar los datos recuperados, detalle: " + error_insert_all.Error(), ""
	}

	return 200, false, "", "Datos recuperados registrados correctamente"
}
