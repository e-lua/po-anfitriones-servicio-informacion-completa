package registro

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	initial_data_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/initial_data"
	//REPOSITORIES
)

func RegisterInitialData(input_anfitrion models.Mo_BusinessWorker_Mqtt) error {

	//Insertamos los datos en Mo
	error_add_business := initial_data_repository.Mo_Add(input_anfitrion)
	if error_add_business != nil {
		log.Println(error_add_business)
		return nil
	}

	//Serializamos el MQTT
	/*var serialize_create models.Mqtt_CreateInitialData
	serialize_create.IDBusiness = input_anfitrion.IdBusiness
	serialize_create.Country = input_anfitrion.IdCountry*/

	//Comenzamos el envio al MQTT
	/*go func() {
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize(serialize_create)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/createpg", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes,
			})
		if error_publish != nil {
			log.Error(error_publish)
		}

	}()

	time.Sleep(1 * time.Second)*/

	//Dejamos de lado el MQTT y nos enfocaremos en crear el negocio

	business_data := map[string]interface{}{
		"idbusiness":   input_anfitrion.IdBusiness,
		"country":      input_anfitrion.IdCountry,
		"issubsidiary": input_anfitrion.IsSubsidiary,
		"subsidiaryof": input_anfitrion.SubsidiaryOf,
	}
	json_data, _ := json.Marshal(business_data)
	http.Post("http://c-busqueda.restoner-api.fun/v1/business/create", "application/json", bytes.NewBuffer(json_data))

	return nil
}

//SERIALIZADORA SCHEDULE
/*func serialize(serialize_initialdata models.Mqtt_CreateInitialData) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_initialdata)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}*/
