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

	db := models.Conectar_Pg_DB()

	idbusiness_pg, idpaymenth_pg, isavailable_pg := []int{}, []int{}, []bool{}
	for _, v := range input_mo_business.PaymentMethods {
		if v.IsAvaiable {
			idbusiness_pg = append(idbusiness_pg, idbusiness)
			idpaymenth_pg = append(idpaymenth_pg, v.IDPaymenth)
			isavailable_pg = append(isavailable_pg, true)
		}
	}

	defer db.Close()

	/*
		query := `INSERT INTO Business_R_Paymenth(idbusiness,idPayment,isavailable) (select * from unnest($1::int[], $2::int[],$3::boolean[]))`
		if _, err := db.Exec(context.Background(), query, idbusiness_pg, idpaymenth_pg, isavailable_pg); err != nil {
			return err
		}*/

	//Serializamos el MQTT
	var serialize_paymenth models.Mqtt_PaymentMethod
	serialize_paymenth.Idbusiness_pg = idbusiness_pg
	serialize_paymenth.Idpaymenth_pg = idpaymenth_pg
	serialize_paymenth.Isavailable_pg = isavailable_pg
	serialize_paymenth.IdBusiness = idbusiness

	//Comenzamos el envio al MQTT

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			defer ch.Close()
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

		defer ch.Close()
	}()

	time.Sleep(2 * time.Second)

	return nil
}

//SERIALIZADORA
func serialize(serialize_paymenth models.Mqtt_PaymentMethod) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_paymenth)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
