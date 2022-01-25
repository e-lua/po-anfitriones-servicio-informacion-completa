package repositories

import (
	"bytes"
	"encoding/json"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_UpdateUniquename(uniquename string, idbusiness int) error {

	//Serializamos el MQTT
	var serialize_unique models.Mqtt_Uniquename
	serialize_unique.Uniquename = uniquename
	serialize_unique.IdBusiness = idbusiness

	//Comenzamos el envio al MQTT

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize_uniquename(serialize_unique)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/uniquename", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes,
			})
		if error_publish != nil {
			log.Error(error_publish)
		}

	}()

	return nil
}

//SERIALIZADORA
func serialize_uniquename(serialize_unique models.Mqtt_Uniquename) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_unique)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
