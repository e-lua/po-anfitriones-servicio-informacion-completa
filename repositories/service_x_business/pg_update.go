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

	idbusiness_pg, idservice_pg, pricing_pg, typemoney_pg, isavailable_pg := []int{}, []int{}, []float32{}, []int{}, []bool{}
	for _, v := range input_mo_business.Services {
		if v.IsAvaiable {
			idbusiness_pg = append(idbusiness_pg, idbusiness)
			idservice_pg = append(idservice_pg, v.IDService)
			pricing_pg = append(pricing_pg, v.Price)
			typemoney_pg = append(typemoney_pg, v.TypeMoney)
			isavailable_pg = append(isavailable_pg, true)
		}
	}

	//Serializamos el MQTT
	var serialize_service models.Mqtt_Service
	serialize_service.Idbusiness_pg = idbusiness_pg
	serialize_service.Idservice_pg = idservice_pg
	serialize_service.Pricing_pg = pricing_pg
	serialize_service.TypeMoney_pg = typemoney_pg
	serialize_service.Isavailable_pg = isavailable_pg
	serialize_service.IdBusiness = idbusiness

	//Comenzamos el envio al MQTT

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize(serialize_service)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/service", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes,
			})
		if error_publish != nil {
			log.Error(error_publish)
		}

	}()

	time.Sleep(2 * time.Second)

	return nil
}

//SERIALIZADORA
func serialize(serialize_service models.Mqtt_Service) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_service)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
