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

	idbusiness_pg, idpaymenth_pg, isavailable_pg, phonenumber_pg := []int{}, []int{}, []bool{}, []string{}
	for _, v := range input_mo_business.PaymentMethods {
		if v.IsAvaiable {
			idbusiness_pg = append(idbusiness_pg, idbusiness)
			idpaymenth_pg = append(idpaymenth_pg, v.IDPaymenth)
			isavailable_pg = append(isavailable_pg, true)
			phonenumber_pg = append(phonenumber_pg, v.PhoneNumber)
		}
	}

	//Serializamos el MQTT
	var serialize_paymenth models.Mqtt_PaymentMethod
	serialize_paymenth.Idbusiness_pg = idbusiness_pg
	serialize_paymenth.Idpaymenth_pg = idpaymenth_pg
	serialize_paymenth.Isavailable_pg = isavailable_pg
	serialize_paymenth.IdBusiness = idbusiness
	serialize_paymenth.PhoneNumber = phonenumber_pg

	//Comenzamos el envio al MQTT

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize(serialize_paymenth)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/paymenth", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes,
			})
		if error_publish != nil {
			log.Error(error_publish)
		}

	}()

	time.Sleep(1 * time.Second)

	return nil
}

//SERIALIZADORA PAYMENTH
func serialize(serialize_paymenth models.Mqtt_PaymentMethod) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_paymenth)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
