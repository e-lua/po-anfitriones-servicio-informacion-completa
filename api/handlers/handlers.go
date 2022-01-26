package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	informacion "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/services/flujo_de_informacion"
	register "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/services/flujo_de_sesion/register_from_initialdata"
	recover "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/services/recover_data"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Consumidor-MQTT
	go Consumer_Data()
	go Consumer_Banner_Mo()
	go Consumer_ViewInformation()
	go Consumer_ViewElement()

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	/*====================FLUJO DE INFORMACIÓN====================*/

	/*===========INFORMACION===========*/

	//V1 FROM V1 TO ...TO ENTITY BUSINESS
	router_business := version_1.Group("/business")
	router_business.GET("", informacion.InformacionRouter_mo.GetInformationData)
	router_business.GET("/basicdata", informacion.InformacionRouter_mo.GetBasicData)
	router_business.GET("/comensal/bnss/:idbusiness", informacion.InformacionRouter_mo.GetInformationData_a_Comensal)

	//V1 FROM BUSINESS TO ...NAME
	router_business.PUT("/name", informacion.InformacionRouter_mo.UpdateName)

	//V1 FROM BUSINESS TO ...UNIQUENAME
	router_business.PUT("/uniquename", informacion.InformacionRouter_mo.UpdateUniqueName)

	//V1 FROM BUSINESS TO ...TYMEZONE
	router_business.PUT("/timezone", informacion.InformacionRouter_mo.UpdateTimeZone)

	//V1 FROM BUSINESS TO ...ADDRESS
	router_business.PUT("/address", informacion.InformacionRouter_mo.UpdateAddress)
	router_business.GET("/address", informacion.InformacionRouter_mo.FindAddress)

	//V1 FROM BUSINESS TO ...TYPEFOOD
	router_business.PUT("/typefood", informacion.InformacionRouter_mo.UpdateTypeFood)
	router_business.GET("/typefood", informacion.InformacionRouter_mo.FindTypeFood)

	//V1 FROM BUSINESS TO ...SERVICE
	router_business.PUT("/service", informacion.InformacionRouter_mo.UpdateService)
	router_business.GET("/service", informacion.InformacionRouter_mo.FindService)

	//V1 FROM BUSINESS TO ...DELIVERYRANGE
	router_business.PUT("/deliveryrange", informacion.InformacionRouter_mo.UpdateDeliveryRange)
	router_business.GET("/deliveryrange", informacion.InformacionRouter_mo.FindDeliveryRange)

	//V1 FROM BUSINESS TO ...DELIVERYRANGE
	router_business.PUT("/paymentmethod", informacion.InformacionRouter_mo.UpdatePaymenthMeth)
	router_business.GET("/paymentmethod", informacion.InformacionRouter_mo.FindPaymenthMeth)

	//V1 FROM BUSINESS TO ...SCHEDULE
	router_business.PUT("/schedule", informacion.InformacionRouter_mo.UpdateSchedule)
	router_business.GET("/schedule", informacion.InformacionRouter_mo.FindSchedule)

	//V1 FROM BUSINESS TO ...PHONECONTACT
	router_business.PUT("/contact", informacion.InformacionRouter_mo.UpdateContact)

	//V1 FROM BUSINESS TO ...RECOVERDATA
	router_business.POST("/recoverdata_all", recover.RecoverRouter_mo.RecoverAll)
	router_business.POST("/recoverdata_one", recover.RecoverRouter_mo.RecoverOne)

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "5000"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Acceso no autorizado")
}

func Consumer_Data() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/businessdata", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop := make(chan bool)

	go func() {
		for d := range msgs {
			var anfitrion models.Mo_BusinessWorker_Mqtt
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&anfitrion)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			register.RegisterFrom_SAInitialData.RegisterInitialData(anfitrion)
		}
	}()

	<-noStop

}

func Consumer_Banner_Mo() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/bannermo", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop3 := make(chan bool)

	go func() {
		for d := range msgs {
			var banner models.Mo_BusinessBanner_Mqtt
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&banner)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformacionRouter_mo.UpdateBanners_Consumer(banner)
		}
	}()

	<-noStop3

}

func Consumer_ViewInformation() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("comensal/viewinformation", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop4 := make(chan bool)

	go func() {
		for d := range msgs {
			var view models.Mqtt_View_Information
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&view)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformacionRouter_mo.UpdateViewInformation_Consumer(view)
		}
	}()

	<-noStop4

}

func Consumer_ViewElement() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("comensal/viewelement", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop5 := make(chan bool)

	go func() {
		for d := range msgs {
			var view models.Mqtt_View_Element
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&view)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformacionRouter_mo.UpdateViewElement_Consumer(view)
		}
	}()

	<-noStop5

}
