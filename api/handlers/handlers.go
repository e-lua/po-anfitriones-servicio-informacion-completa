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
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Consumidor-MQTT
	go Consumer_Data()
	go Consumer_Banner()

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	/*====================FLUJO DE INFORMACIÓN====================*/

	/*===========INFORMACION===========*/

	//V1 FROM V1 TO ...TO ENTITY BUSINESS
	router_business := version_1.Group("/business")
	//1
	router_business.GET("", informacion.InformacionRouter_mo.GetInformationData)

	//V1 FROM BUSINESS TO ...BANNER
	//21
	router_business.GET("/banner", informacion.InformacionRouter_mo.FindBanner)

	//V1 FROM BUSINESS TO ...NAME
	//3
	router_business.PUT("/name", informacion.InformacionRouter_mo.UpdateName)
	//31
	//router_business.GET("/name", informacion.informacionRouter_mo.FindName)

	//V1 FROM BUSINESS TO ...ADDRESS
	//4
	router_business.PUT("/address", informacion.InformacionRouter_mo.UpdateAddress)
	//41
	//router_business.GET("/address", informacion.informacionRouter_mo.FindAddress)

	//V1 FROM BUSINESS TO ...TYPEFOOD
	//5
	router_business.PUT("/typefood", informacion.InformacionRouter_mo.UpdateTypeFood)
	//51
	//router_business.GET("/typefood", informacion.informacionRouter_mo.FindTypeFood)
	//router_business.GET("/typefood/all", informacion.InformacionRouter_mo.FindAllTypeFood)

	//V1 FROM BUSINESS TO ...SERVICE
	//6
	router_business.PUT("/service", informacion.InformacionRouter_mo.UpdateService)
	//61
	//router_business.GET("/service", informacion.informacionRouter_mo.FindService)
	//router_business.GET("/service/all", informacion.informacionRouter_mo.FindAllService)

	//V1 FROM BUSINESS TO ...DELIVERYRANGE
	//7
	router_business.PUT("/deliveryrange", informacion.InformacionRouter_mo.UpdateDeliveryRange)
	//71
	//router_business.GET("/deliveryrange", informacion.informacionRouter_mo.FindDeliveryRange)

	//V1 FROM BUSINESS TO ...DELIVERYRANGE
	//8
	router_business.PUT("/paymentmethod", informacion.InformacionRouter_mo.UpdatePaymenthMeth)
	//81
	//router_business.GET("/paymentmethod", informacion.informacionRouter_mo.FindPaymenthMeth)
	//router_business.GET("/paymentmethod/all", informacion.informacionRouter_mo.FindAllPaymenth)

	//V1 FROM BUSINESS TO ...SCHEDULE
	//9
	router_business.PUT("/schedule", informacion.InformacionRouter_mo.UpdateSchedule)
	//91
	//router_business.GET("/schedule", informacion.informacionRouter_mo.FindSchedule)
	//router_business.GET("/schedule/all", informacion.informacionRouter_mo.FindAllSchedule)

	//V1 FROM BUSINESS TO ...PHONECONTACT
	//101
	router_business.PUT("/contact", informacion.InformacionRouter_mo.UpdateContact)
	//101
	//router_business.GET("/contact", informacion.informacionRouter_mo.FindContact)
	//router_business.GET("/contact/all", informacion.informacionRouter_mo.FindAllContact)

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
		defer ch.Close()
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/businessdata", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

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

}

func Consumer_Banner() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		defer ch.Close()
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/banner", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	go func() {
		for d := range msgs {
			var toCarta models.Mo_ToBanner_Mqtt
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&toCarta)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformacionRouter_mo.UpdateBanners_Consumer(toCarta.IdBanner_Category_Element, toCarta.Url, toCarta.IdBusiness)
		}
	}()

}
