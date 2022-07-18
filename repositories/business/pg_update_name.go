package repositories

import (
	"bytes"
	"encoding/json"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_UpdateName(name string, idbusiness int) error {

	//Serializamos el MQTT
	var serialize_n models.Mqtt_Name
	serialize_n.Name = name
	serialize_n.IdBusiness = idbusiness

	//Comenzamos el envio al MQTT
	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize_name(serialize_n)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/name", false, false,
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
func serialize_name(serialize_name models.Mqtt_Name) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_name)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
