package repositories

import (
	"bytes"
	"encoding/json"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_Update(input_mo_business models.Mo_Business, idbusiness int) error {

	idbusiness_pg, Idtypefood_pg, isavailable_pg := []int{}, []int{}, []bool{}
	for _, v := range input_mo_business.TypeOfFood {
		if v.IsAvaiable {
			idbusiness_pg = append(idbusiness_pg, idbusiness)
			Idtypefood_pg = append(Idtypefood_pg, v.IDTypeFood)
			isavailable_pg = append(isavailable_pg, true)
		}
	}

	//Serializamos el MQTT
	var serialize_typefood models.Mqtt_TypeFood
	serialize_typefood.Idbusiness_pg = idbusiness_pg
	serialize_typefood.Idtypefood_pg = Idtypefood_pg
	serialize_typefood.Isavailable_pg = isavailable_pg
	serialize_typefood.IdBusiness = idbusiness

	//Comenzamos el envio al MQTT

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			defer ch.Close()
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize(serialize_typefood)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/typefood", false, false,
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

	time.Sleep(2 * time.Second)

	return nil
}

//SERIALIZADORA
func serialize(serialize_typefoo models.Mqtt_TypeFood) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_typefoo)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
