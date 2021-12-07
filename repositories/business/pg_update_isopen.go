package repositories

import (
	"bytes"
	"encoding/json"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_UpdateIsOpen(isOpen bool, idbusiness int) error {

	//Serializamos el MQTT
	var serialize_open models.Mqtt_IsOpen
	serialize_open.IsaOpen = isOpen
	serialize_open.IdBusiness = idbusiness

	//Comenzamos el envio al MQTT

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			defer ch.Close()
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize_isopen(serialize_open)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/isopen", false, false,
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
func serialize_isopen(serialize_open models.Mqtt_IsOpen) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_open)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
