package recover

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

var RecoverRouter_mo *recoverRouter_mo

type recoverRouter_mo struct {
}

func (rr *recoverRouter_mo) RecoverAll(c echo.Context) error {

	//Aseguramos que no recupere los datos cualquier persona
	key_string := c.Request().Header.Get("Key")
	if key_string != "ods8SAEUYng87dhdfn8hfna9s76fnnsaiosr7ffi9nasm" {
		results := Response{Error: true, DataError: "No puede acceder a este recurso debido a que no tiene autorizacion", Data: ""}
		return c.JSON(403, results)
	}

	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://c-busqueda.restoner-api.fun:6850/v1/recover/all")
	var get_respuesta Response_GetAllBusiness
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: true, DataError: "Error en el servidor interno al intentar obtener los datos del servicio de soporte de datos del anfitrion, detalles " + error_decode_respuesta.Error(), Data: ""}
		return c.JSON(500, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := RecoverAll_Service(get_respuesta.Data)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (rr *recoverRouter_mo) RecoverOne(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")
	if idbusiness_string == "" {
		results := Response{Error: true, DataError: "Debe enviar el id del negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Aseguramos que no recupere los datos cualquier persona
	key := c.Request().Header.Get("Key")
	if key != "ods8SAEUYng87dhdfn8hfna9s76fnnsaiosr7ffi9nasm" {
		results := Response{Error: true, DataError: "No puede acceder a este recurso debido a que no tiene autorizacion", Data: ""}
		return c.JSON(403, results)
	}

	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://c-busqueda.restoner-api.fun:6850/v1/recover/one?idbusiness=" + idbusiness_string)
	var get_respuesta Response_GetOneBusiness
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: true, DataError: "Error en el servidor interno al intentar obtener los datos del servicio de soporte de datos del anfitrion, detalles " + error_decode_respuesta.Error(), Data: ""}
		return c.JSON(500, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := RecoverOne_Service(get_respuesta.Data)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
