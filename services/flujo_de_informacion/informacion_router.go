package informacion

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var InformacionRouter_mo *informacionRouter_mo

type informacionRouter_mo struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string, service int, module int, epic int, endpoint int) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt + "&service=" + strconv.Itoa(service) + "&module=" + strconv.Itoa(module) + "&epic=" + strconv.Itoa(epic) + "&endpoint=" + strconv.Itoa(endpoint))
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

func GetJWT_Country(jwt string, service int, module int, epic int, endpoint int) (int, bool, string, int, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt + "&service=" + strconv.Itoa(service) + "&module=" + strconv.Itoa(module) + "&epic=" + strconv.Itoa(epic) + "&endpoint=" + strconv.Itoa(endpoint))
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0, 0
	}
	return 200, false, "", get_respuesta.Data.IdCountry, get_respuesta.Data.IdBusiness
}

func GetJWT_Rol(jwt string, service int, module int, epic int, endpoint int) (int, bool, string, int, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt + "&service=" + strconv.Itoa(service) + "&module=" + strconv.Itoa(module) + "&epic=" + strconv.Itoa(epic) + "&endpoint=" + strconv.Itoa(endpoint))
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0, 0
	}
	return 200, false, "", get_respuesta.Data.IdRol, get_respuesta.Data.IdBusiness
}

func GetJWT_Comensal(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT_Comensal
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}

/*----------------------CONSUMER----------------------*/

func (ir *informacionRouter_mo) UpdateBanners_Consumer(banner models.Mo_BusinessBanner_Mqtt) {
	//Enviamos los datos al servicio
	error_consumer_banner := UpdateBanners_Consumer_Service(banner)
	log.Println(error_consumer_banner)

}

func (ir *informacionRouter_mo) Manual_UpdateBanners_Consumer(c echo.Context) error {

	//Instanciamos una variable del modelo B_Description
	var banner models.Mo_BusinessBanner_Mqtt

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&banner)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el nombre del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_consumer_banner := UpdateBanners_Consumer_Service(banner)
	log.Println(error_consumer_banner)

	return nil
}

func (ir *informacionRouter_mo) Manual_UpdatePost_Consumer(c echo.Context) error {

	//Instanciamos una variable del modelo B_Description
	var post models.Mo_BusinessPost_Mqtt

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&post)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el nombre del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_consumer_post := UpdatePosts_Consumer_Service(post)
	log.Println(error_consumer_post)
	return nil
}

/*----------------------UPDATE DATA OF THE BUSINESS----------------------*/

func (ir *informacionRouter_mo) UpdateDescription(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Description
	var b_description B_Description

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_description)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el nombre del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(b_description.Description) > 300 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateDescription_Service(data_idbusiness, b_description)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateName(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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
	if len(b_name.Name) > 50 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateName_Service(data_idbusiness, b_name)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateLegalIdentity(inputserialize_legalidentity []models.Mqtt_LegalIdentity) {

	//Enviamos los datos al servicio
	error_r := UpdateLegalIdentity_Service(inputserialize_legalidentity)
	log.Println(error_r)

}
func (ir *informacionRouter_mo) UpdateUniqueName(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idrol, data_idbusiness := GetJWT_Rol(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if data_idrol != 1 {
		results := Response{Error: true, DataError: "111" + "Rol incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var uniquename_string B_Uniquename

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&uniquename_string)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el nombre del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	res1, _ := regexp.MatchString(`restoner`, uniquename_string.Uniquename)

	//Evitando modificar la cuenta demo
	if data_idbusiness == 24 {
		results := Response{Error: true, DataError: "Cuenta DEMO"}
		return c.JSON(403, results)
	}

	//Validamos los valores enviados
	if res1 {
		results := Response{Error: true, DataError: "333" + "El valor ingresado no cumple con la regla de negocio, no debe contener el texto restoner"}
		return c.JSON(403, results)
	}

	//Validamos el texto
	counter_arroba := 0
	counter := 0
	uniquename_lower := strings.ToLower(uniquename_string.Uniquename)

	for i := 0; i < len(uniquename_lower); i++ {
		if uniquename_lower[i] > 47 && uniquename_lower[i] < 58 || uniquename_lower[i] > 96 && uniquename_lower[i] < 123 || uniquename_lower[i] == 95 || uniquename_lower[i] == 64 {
			counter = counter + 1
		} else {
			counter = counter + 30
		}
		if uniquename_lower[i] == 64 {
			counter_arroba = counter_arroba + 10
		}
		if uniquename_lower[0] != 64 {
			counter_arroba = counter_arroba + 100
		}
	}

	//Validamos los valores enviados
	if len(uniquename_string.Uniquename) > 25 || len(uniquename_string.Uniquename) < 8 || counter > 27 || counter_arroba != 10 {
		results := Response{Error: true, DataError: "333" + "El valor ingresado no cumple con la regla de negocio, el uniquename debe contener maximo 25 caracteres y minimo 8 caracteres, debe tener un @(arroba) al comienzo del texto y solo un @, y no enviar caracteres especiales"}
		return c.JSON(403, results)
	}

	//Validamos que no exista el uniquename ya creado
	respuesta, _ := http.Get("http://c-busqueda.restoner-api.fun:80/v1/business/uniquenames?uniquename=" + uniquename_string.Uniquename)
	var get_respuesta Response_Uniquenames
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: true, DataError: "Error en el sevidor interno al intentar buscar los nombre unicos, detalles: " + error_decode_respuesta.Error(), Data: ""}
		return c.JSON(500, results)
	}
	if get_respuesta.Data != "" {
		results := Response{Error: true, DataError: "222" + "Este nombre ya ha sido tomado", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateUniqueName_Service(data_idbusiness, uniquename_lower)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateTimeZone(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_business models.Mo_Business

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_business)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el nombre del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	timezone, _ := strconv.Atoi(b_business.TimeZone)

	//Validamos los valores enviados
	if timezone > 13 || timezone < -11 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateTimeZone_Service(data_idbusiness, b_business)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateAddress(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 4)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo de negocio
	var mo_business models.Mo_Business

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_business)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar la latitud, longitud, codigo postal y direccion completa del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 || len(mo_business.Address.FullAddress) < 5 || mo_business.Address.PostalCode < 0 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateAddress_Service(data_idbusiness, mo_business)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateTypeFood(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 5)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo de negocio
	var mo_business models.Mo_Business

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_business)
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
	status, boolerror, dataerror, data := UpdateTypeFood_Service(data_idbusiness, mo_business)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateService(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 6)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo de negocio
	var mo_business models.Mo_Business

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_business)
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
	status, boolerror, dataerror, data := UpdateService_Service(data_idbusiness, mo_business)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateDeliveryRange(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 7)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var b_delivery models.Mo_Delivery

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&b_delivery)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el rango de reparto, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateDeliveryRange_Service(data_idbusiness, b_delivery)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdatePaymenthMeth(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 8)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo B_Name
	var mo_business models.Mo_Business

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_business)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del tiop de comida, el nombre, la url de la imagen y si esta disopnible o no, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdatePaymenthMeth_Service(data_idbusiness, mo_business)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateSchedule(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 9)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo de negocio
	var mo_business models.Mo_Business

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_business)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id de todos los dias, la hora de inicio, la hora de fin, y si esta disponible o no, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateSchedule_Service(data_idbusiness, mo_business)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateContact(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 10)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo de negocio
	var mo_business models.Mo_Business

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_business)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del contacto,el nombre el dato, y si esta disponible o no, revise la estructura o los valores ", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateContact_Service(data_idbusiness, mo_business)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*------------------------------------------------POST---------------------------------------------------------------*/

func (ir *informacionRouter_mo) AddPost(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry, data_idbusiness := GetJWT_Country(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo de negocio
	var mo_post models.Mo_Post

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_post)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar la fecha de eliminación y la url de la imagen, revise la estructura o los valores ", Data: ""}
		return c.JSON(400, results)
	}

	mo_post.IdBusiness = data_idbusiness
	mo_post.Dateregistered = time.Now()
	mo_post.Uuid = uuid.New().String()

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddPost_Service(mo_post)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) GetPost(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry, data_idbusiness := GetJWT_Country(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	limit_string := c.Param("limit")
	limit_int, _ := strconv.ParseInt(limit_string, 10, 64)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetPost_Service(data_idbusiness, limit_int)
	results := Response_Posts{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) DeletePost(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry, data_idbusiness := GetJWT_Country(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	uuidpost := c.Param("uuidpost")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := DeletePost_Service(data_idbusiness, uuidpost)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*-------------------------------------------------------------------------------------------------------------------*/

/*---------------------------------------------COMENTS---------------------------------------------------------------*/

func (ir *informacionRouter_mo) AddComment(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo de negocio
	var mo_comment models.Mo_Comment

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_comment)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del contacto,el nombre el dato, y si esta disponible o no, revise la estructura o los valores ", Data: ""}
		return c.JSON(400, results)
	}

	mo_comment.IDComensal = data_idcomensal

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddComment_Service(mo_comment)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) GetCommentsBusiness(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 10)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	page_string := c.Request().URL.Query().Get("page")
	page_int, _ := strconv.ParseInt(page_string, 10, 64)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCommentsBusiness_Service(data_idbusiness, page_int)
	results := Response_Comments{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) GetCommentsStadistics(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 10)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCommentsStadistics_Service(data_idbusiness)
	results := Response_Comment_Resume{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) GetCommentsComensal(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness_string)

	page_string := c.Request().URL.Query().Get("page")
	page_int, _ := strconv.ParseInt(page_string, 10, 64)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCommentsComensal_Service(idbusiness_int, page_int)
	results := Response_Comments_Comensal{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) GetCommentsOne_Comensal(c echo.Context) error {

	idcomensal_string := c.Request().URL.Query().Get("idcomensal")
	idcomensal_int, _ := strconv.Atoi(idcomensal_string)

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCommentsOne_Comensal_Service(idbusiness_int, idcomensal_int)
	results := Response_CommentFound{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateCommentBusiness(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 10)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del Business Owner
	idcomment := c.Param("idcomment")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCommentBusiness_Service(idcomment)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) UpdateCommentComensal(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del Business Owner
	idcomment := c.Param("idcomment")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCommentComensal_Service(idcomment)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) AddCommentReport(c echo.Context) error {

	//Instanciamos una variable del modelo de negocio
	var mo_reporte_comment models.Mo_Comment_Reported

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_reporte_comment)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del contacto,el nombre el dato, y si esta disponible o no, revise la estructura o los valores ", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if mo_reporte_comment.IDReason < 1 || mo_reporte_comment.Reason == "" {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddCommentReport_Service(mo_reporte_comment)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*------------------------------------------------------------------------------------------------------------*/

func (ir *informacionRouter_mo) AddBusinessReport(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo de negocio
	var mo_reporte_business models.Mo_Business_Reported

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&mo_reporte_business)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el id del contacto,el nombre el dato, y si esta disponible o no, revise la estructura o los valores ", Data: ""}
		return c.JSON(400, results)
	}

	mo_reporte_business.IDComensal = data_idcomensal

	//Validamos los valores enviados
	if mo_reporte_business.IDReason < 1 || mo_reporte_business.Reason == "" {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddBusinessReport_Service(mo_reporte_business)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------GET DATA OF THE BUSINESS----------------------*/

func (ir *informacionRouter_mo) FindName(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 31)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")

	idbusiness_int, _ := strconv.Atoi(idbusiness_string)

	//Validamos los valores enviados
	if idbusiness_int < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAddress_Service(idbusiness_int)
	results := ResponseAddress{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindTypeFood(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry, data_idbusiness := GetJWT_Country(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Obtenemos los datos
	respuesta, _ := http.Get("http://c-busqueda.restoner-api.fun:80/v1/export/typefood?idbusiness=" + strconv.Itoa(data_idbusiness) + "&country=" + strconv.Itoa(data_idcountry))
	var get_respuesta ResponseInterface_FromComensal
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := ResponseInterface_FromComensal{Error: true, DataError: "Error en el servidor interno al intentar decodificar la respuesta", Data: get_respuesta.Data}
		return c.JSON(500, results)
	}

	results := ResponseInterface_FromComensal{Error: get_respuesta.Error, DataError: get_respuesta.DataError, Data: get_respuesta.Data}
	return c.JSON(200, results)
}

func (ir *informacionRouter_mo) FindService(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry, data_idbusiness := GetJWT_Country(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Obtenemos los datos
	respuesta, _ := http.Get("http://c-busqueda.restoner-api.fun:80/v1/export/service?idbusiness=" + strconv.Itoa(data_idbusiness) + "&country=" + strconv.Itoa(data_idcountry))
	var get_respuesta ResponseInterface_FromComensal
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := ResponseInterface_FromComensal{Error: true, DataError: "Error en el servidor interno al intentar decodificar la respuesta", Data: get_respuesta.Data}
		return c.JSON(500, results)
	}

	//Enviamos los datos al servicio
	results := ResponseInterface_FromComensal{Error: get_respuesta.Error, DataError: get_respuesta.DataError, Data: get_respuesta.Data}
	return c.JSON(status, results)
}

func (ir *informacionRouter_mo) FindDeliveryRange(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 71)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcountry <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Obtenemos los datos
	respuesta, _ := http.Get("http://c-busqueda.restoner-api.fun:80/v1/export/payment?idbusiness=" + strconv.Itoa(data_idbusiness) + "&country=" + strconv.Itoa(data_idcountry))
	var get_respuesta ResponseInterface_FromComensal
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := ResponseInterface_FromComensal{Error: true, DataError: "Error en el servidor interno al intentar decodificar la respuesta", Data: get_respuesta.Data}
		return c.JSON(500, results)
	}

	//Enviamos los datos al paymenth
	results := ResponseInterface_FromComensal{Error: get_respuesta.Error, DataError: get_respuesta.DataError, Data: get_respuesta.Data}
	return c.JSON(status, results)

}

func (ir *informacionRouter_mo) FindSchedule(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 91)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Obtenemos los datos
	respuesta, _ := http.Get("http://c-busqueda.restoner-api.fun:80/v1/export/schedule?idbusiness=" + strconv.Itoa(data_idbusiness))
	var get_respuesta ResponseInterface_FromComensal
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := ResponseInterface_FromComensal{Error: true, DataError: "Error en el servidor interno al intentar decodificar la respuesta", Data: get_respuesta.Data}
		return c.JSON(500, results)
	}

	//Enviamos los datos al servicio
	results := ResponseInterface_FromComensal{Error: get_respuesta.Error, DataError: get_respuesta.DataError, Data: get_respuesta.Data}
	return c.JSON(status, results)
}

func (ir *informacionRouter_mo) FindContact(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 101)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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

func (ir *informacionRouter_mo) GetBasicData(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 1)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if data_idbusiness < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Obtenemos los datos
	respuesta, _ := http.Get("http://c-busqueda.restoner-api.fun:80/v1/export/basicdata?idbusiness=" + strconv.Itoa(data_idbusiness))
	var get_respuesta ResponseInterface_FromComensal
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := ResponseInterface_FromComensal{Error: true, DataError: "Error en el servidor interno al intentar decodificar la respuesta", Data: get_respuesta.Data}
		return c.JSON(500, results)
	}

	//Enviamos los datos al servicio
	results := ResponseInterface_FromComensal{Error: get_respuesta.Error, DataError: get_respuesta.DataError, Data: get_respuesta.Data}
	return c.JSON(status, results)
}

/*----------------------SERVIMOS LOS DATOS CON CONSULTA DEL COMENSAL----------------------*/

func (ir *informacionRouter_mo) GetInformationData_a_Comensal(c echo.Context) error {

	//Recibimos el id del Business Owner
	idbusiness := c.Param("idbusiness")

	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetInformationData_a_Comensal_Service(idbusiness_int)
	results := ResponseBusiness{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *informacionRouter_mo) GetPostData_a_Comensal(c echo.Context) error {

	//Recibimos el id del Business Owner
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	limit_string := c.Param("limit")
	limit_int, _ := strconv.ParseInt(limit_string, 10, 64)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetPost_Service(idbusiness_int, limit_int)
	results := Response_Posts{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------INSERTAMOS LAS VISTAS----------------------*/

func (ir *informacionRouter_mo) UpdateViewInformation_Consumer(mo_view_info models.Mqtt_View_Information) {

	//Enviamos los datos al servicio
	error_r := UpdateViewInformation_Consumer_Service(mo_view_info)
	log.Println(error_r)
}

func (ir *informacionRouter_mo) UpdateViewElement_Consumer(mo_view_element models.Mqtt_View_Element) {

	//Enviamos los datos al servicio
	error_r := UpdateViewElement_Consumer_Service(mo_view_element)
	log.Println(error_r)
}
