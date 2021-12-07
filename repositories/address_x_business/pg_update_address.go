package repositories

import (
	"bytes"
	"encoding/json"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_UpdateAddress(address models.Mo_Address, idbusiness int) error {

	//Serializamos el MQTT
	var serialize_add models.Mqtt_Address
	serialize_add.Latitude = address.Latitude
	serialize_add.IdBusiness = idbusiness
	serialize_add.Longitude = address.Longitude
	//Comenzamos el envio al MQTT

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			defer ch.Close()
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize_address(serialize_add)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/address", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes,
			})
		if error_publish != nil {
			log.Error(error_publish)
		}

		defer ch.Close()
	}()

	return nil
}

//SERIALIZADORA
func serialize_address(inputserialize_add models.Mqtt_Address) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(inputserialize_add)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
