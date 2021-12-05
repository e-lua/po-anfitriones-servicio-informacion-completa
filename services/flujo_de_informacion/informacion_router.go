package informacion

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
)

var InformacionRouter_mo *informacionRouter_mo

type informacionRouter_mo struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string, service int, module int, epic int, endpoint int) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://143.198.75.79:5000/v1/trylogin?jwt=" + jwt + "&service=" + strconv.Itoa(service) + "&module=" + strconv.Itoa(module) + "&epic=" + strconv.Itoa(epic) + "&endpoint=" + strconv.Itoa(endpoint))
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

func GetJWT_Country(jwt string, service int, module int, epic int, endpoint int) (int, bool, string, int, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://143.198.75.79:5000/v1/trylogin?jwt=" + jwt + "&service=" + strconv.Itoa(service) + "&module=" + strconv.Itoa(module) + "&epic=" + strconv.Itoa(epic) + "&endpoint=" + strconv.Itoa(endpoint))
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0, 0
	}
	return 200, false, "", get_respuesta.Data.IdCountry, get_respuesta.Data.IdBusiness
}

/*----------------------CONSUMER----------------------*/

func (ir *informacionRouter_mo) UpdateBanners_Consumer(idbanner int, urlphoto string, idbusiness int) {
	//Enviamos los datos al servicio
	error_consumer_banner := UpdateBanners_Consumer_Service(idbanner, urlphoto, idbusiness)
	if error_consumer_banner != nil {
		log.Fatal(error_consumer_banner)
	}
}

/*----------------------SHOW ALL PUBLIC DATA----------------------*/
/*
func (ir *informacionRouter_mo) FindAllPaymenth(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry, data_idbusiness := GetJWT_Country(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllPaymenth_Service(data_idcountry, data_idbusiness)
	results := ResponsePaymethAll{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}*/

/*
func (ir *informacionRouter_mo) FindAllService(c echo.Context) error {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllService_Service()
	results := ResponseServiceAll{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindAllTypeFood(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry := GetJWT_Country(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllTypeFood_Service(data_idcountry)
	results := ResponseTypeFoodAll{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindAllSchedule(c echo.Context) error {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllSchedule_Service()
	results := ResponseScheduleAll{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindAllContact(c echo.Context) error {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllContact_Service()
	results := ResponseContactAll{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}*/

/*----------------------UPDATE DATA OF THE BUSINESS----------------------*/

func (ir *informacionRouter_mo) UpdateName(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_name B_Name

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_name)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el nombre del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(b_name.Name) > 20 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateName_Service(data_idbusiness, b_name)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateAddress(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 4)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_address models.Mo_Address

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_address)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar la latitud, longitud, codigo postal y direccion completa del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 || len(b_address.FullAddress) < 10 || b_address.PostalCode < 0 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateAddress_Service(data_idbusiness, b_address)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateTypeFood(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 5)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_TypeFood []models.Mo_TypeFood

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_TypeFood)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del tipo de comida e indicar si esta habilitado, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateTypeFood_Service(data_idbusiness, b_TypeFood)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateService(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 6)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_service []models.Mo_Service

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_service)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id, precio, tipo de moneda y disponibilidad del servicio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateService_Service(data_idbusiness, b_service)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateDeliveryRange(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 7)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_deliveryrange B_DeliveryRange

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_deliveryrange)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del tipo de comida, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateDeliveryRange_Service(data_idbusiness, b_deliveryrange)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdatePaymenthMeth(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 8)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_paymenthmeth []models.Mo_PaymenthMeth

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_paymenthmeth)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del tipo de comida, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdatePaymenthMeth_Service(data_idbusiness, b_paymenthmeth)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateSchedule(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 9)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_schedule []models.Mo_Day

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_schedule)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del tipo de comida, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateSchedule_Service(data_idbusiness, b_schedule)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateContact(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 10)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_contact []models.Mo_Contact

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_contact)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del tipo de comida, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateContact_Service(data_idbusiness, b_contact)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------GET DATA OF THE BUSINESS----------------------*/

func (ir *informacionRouter_mo) FindName(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 31)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindName_Service(data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindAddress(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 41)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAddress_Service(data_idbusiness)
	results := ResponseAddress{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindTypeFood(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 51)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindTypeFood_Service(data_idbusiness)
	results := ResponseTypeFood{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindService(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 61)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindService_Service(data_idbusiness)
	results := ResponseService{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindDeliveryRange(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 71)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindDeliveryRange_Service(data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindPaymenthMeth(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry, data_idbusiness := GetJWT_Country(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindPaymenth_Service(data_idcountry, data_idbusiness)
	results := ResponsePaymethAll{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindSchedule(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 91)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindSchedule_Service(data_idbusiness)
	results := ResponseSchedule{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindBanner(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 21)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindBanner_Service(data_idbusiness)
	results := ResponseBanner{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}
func (ir *informacionRouter_mo) FindContact(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 101)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindContact_Service(data_idbusiness)
	results := ResponseContact{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------GET DATA OF THE BUSINESS WITH ONE ENDPOINT----------------------*/

func (ir *informacionRouter_mo) GetInformationData(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 1)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetInformationData_Service(data_idbusiness)
	results := ResponseBusiness{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}
