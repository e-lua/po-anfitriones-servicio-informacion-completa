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

	//Instanciamos los datos
	idday_pg, idbusiness_pg, starttime_pg, endtime_pg, available_pg, name_pg := []int{}, []int{}, []string{}, []string{}, []bool{}, []string{}

	//Convertimos a formato 24 horas
	for _, day := range input_mo_business.DailySchedule {

		//Convertimos la hora
		startTime, _ := time.Parse("15:04 PM", day.StarTime)
		endTime, _ := time.Parse("15:04 PM", day.EndTime)

		//Append
		idday_pg = append(idday_pg, day.IDDia)
		idbusiness_pg = append(idbusiness_pg, idbusiness)
		starttime_pg = append(starttime_pg, startTime.Format("15:04"))
		endtime_pg = append(endtime_pg, endTime.Format("15:04"))
		available_pg = append(available_pg, day.IsAvaiable)
		name_pg = append(name_pg, day.Name)
	}

	//Serializamos el MQTT
	var serialize_schedule models.Mqtt_Schedule
	serialize_schedule.Idbusiness_pg = idbusiness_pg
	serialize_schedule.Isavailable_pg = available_pg
	serialize_schedule.IdBusiness = idbusiness
	serialize_schedule.Idschedule_pg = idday_pg
	serialize_schedule.Starttime_pg = starttime_pg
	serialize_schedule.Endtime_pg = endtime_pg
	serialize_schedule.Name_pg = name_pg

	//Comenzamos el envio al MQTT
	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize(serialize_schedule)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/horario", false, false,
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
func serialize(serialize_schedule models.Mqtt_Schedule) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_schedule)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
