package informacion_web

import (
	"github.com/labstack/echo/v4"
)

var Web_InformacionRouter_mo *webInformacionRouter_mo

type webInformacionRouter_mo struct {
}

/*----------------------SERVIMOS LOS DATOS CON CONSULTA DEL COMENSAL----------------------*/

func (wir *webInformacionRouter_mo) Web_GetInformationData_a_Comensal(c echo.Context) error {

	//Recibimos el id del Business Owner
	uniquename := c.Param("uniquename")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_GetInformationData_a_Comensal_Service(uniquename)
	results := ResponseBusiness{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
