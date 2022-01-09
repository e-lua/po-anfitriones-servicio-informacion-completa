package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Mo_Add(anfitrion_mo models.Mo_BusinessWorker_Mqtt) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_anfitriones")
	col := db.Collection("business")

	var registro models.Mo_Registro_FromMqtt

	registro.IdBusiness = anfitrion_mo.IdBusiness
	registro.CreatedDate = anfitrion_mo.UpdatedDate
	registro.Available = true
	registro.OrdersRejected = 0
	registro.IsOpen = false

	_, error_add := col.InsertOne(ctx, registro)

	if error_add != nil {
		return error_add

	}

	//Serializamos el MQTT
	var serialize_create models.Mqtt_CreateInitialData
	serialize_create.IDBusiness = anfitrion_mo.IdBusiness
	serialize_create.Country = anfitrion_mo.IdCountry

	//Comenzamos el envio al MQTT
	go func() {
		//Comienza el proceso de MQTT
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

	time.Sleep(1 * time.Second)

	return nil

}

//SERIALIZADORA SCHEDULE
func serialize(serialize_initialdata models.Mqtt_CreateInitialData) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_initialdata)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
