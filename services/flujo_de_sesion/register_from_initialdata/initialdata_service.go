package registro

import (

	//MDOELS
	"log"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"

	//REPOSITORIES
	initial_data_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/initial_data"
)

func RegisterInitialData(input_anfitrion models.Mo_BusinessWorker_Mqtt) error {

	//Insertamos los datos en Mo
	error_add_business := initial_data_repository.Mo_Add(input_anfitrion)
	if error_add_business != nil {
		log.Println(error_add_business)
	}

	return nil
}
