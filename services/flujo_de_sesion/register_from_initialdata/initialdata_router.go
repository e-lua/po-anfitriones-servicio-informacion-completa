package registro

import (

	//MODELS
	"log"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

var RegisterFrom_SAInitialData *registerFrom_SAInitialData

type registerFrom_SAInitialData struct {
}

func (cr *registerFrom_SAInitialData) RegisterInitialData(anfitrion models.Mo_BusinessWorker_Mqtt) {

	//Enviamos los datos al servicio
	error_r := RegisterInitialData(anfitrion)
	log.Println(error_r)
}
