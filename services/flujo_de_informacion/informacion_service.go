package informacion

import (

	//REPOSITORIES
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	address_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/address_x_business"
	banner_x_busines_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/banner_x_business"
	business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/business"
	comment_x_business "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/comments"
	contact_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/contact_x_business"
	schedule_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/day_x_business"
	payment_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/paymenth_x_business"
	reports "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/reports"
	service_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/service_x_business"
	typefood_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/typefood_x_business"
	view "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/view"
)

/*----------------------CONSUMER----------------------*/

func UpdateBanners_Consumer_Service(banner models.Mo_BusinessBanner_Mqtt) error {
	error_add_banner_mo := banner_x_busines_repository.Mo_Update(banner)
	if error_add_banner_mo != nil {
		log.Fatal(error_add_banner_mo)
	}
	return nil
}

/*----------------------SERVICES TO UPDATE DATA OF BUSINESS----------------------*/

//DESCRIPCION
func UpdateDescription_Service(inputObjectIdBusiness int, input_b_description B_Description) (int, bool, string, string) {

	error_updatename_mongo := business_repository.Mo_Update_Description(input_b_description.Description, inputObjectIdBusiness)
	if error_updatename_mongo != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el nombre, detalle: " + error_updatename_mongo.Error(), ""
	}

	return 200, false, "", "Description actualizada correctamente"
}
func FindDescriptiion_Service(inputObjectIdBusiness int) (int, bool, string, string) {

	description, _ := business_repository.Mo_Find_Description(inputObjectIdBusiness)
	return 200, false, "", description
}

//NOMBRE
func UpdateName_Service(inputObjectIdBusiness int, input_b_name B_Name) (int, bool, string, string) {

	error_updatename_mongo := business_repository.Mo_Update_Name(input_b_name.Name, inputObjectIdBusiness)
	if error_updatename_mongo != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el nombre, detalle: " + error_updatename_mongo.Error(), ""
	}

	go func() {
		business_repository.Pg_UpdateName(input_b_name.Name, inputObjectIdBusiness)
	}()

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizó el nombre de su negocio a " + input_b_name.Name,
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()
	return 200, false, "", "Nombre actualizado correctamente"
}

func FindName_Service(inputObjectIdBusiness int) (int, bool, string, string) {

	name, _ := business_repository.Mo_Find_Name(inputObjectIdBusiness)
	return 200, false, "", name
}

//LEGAL IDENTITY
func UpdateLegalIdentity_Service(inputserialize_legalidentity_multiple []models.Mqtt_LegalIdentity) error {

	for _, inputserialize_legalidentity := range inputserialize_legalidentity_multiple {
		error_update := business_repository.Mo_Update_Legalidentity(inputserialize_legalidentity)
		if error_update != nil {
			log.Fatal(error_update)
		}
	}

	return nil
}

//UNIQUE-NOMBRE
func UpdateUniqueName_Service(inputObjectIdBusiness int, uniquename string) (int, bool, string, string) {

	error_updatename_mongo := business_repository.Mo_Update_Uniquename(uniquename, inputObjectIdBusiness)
	if error_updatename_mongo != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el nombre, detalle: " + error_updatename_mongo.Error(), ""
	}

	go func() {
		business_repository.Pg_UpdateUniquename(uniquename, inputObjectIdBusiness)
	}()

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizó el nombre único de su negocio a " + uniquename,
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Nombre único actualizado correctamente"
}

//ISOPEN
func UpdateTimeZone_Service(inputObjectIdBusiness int, input_business models.Mo_Business) (int, bool, string, string) {

	error_updatename_mongo := business_repository.Mo_Update_TimeZone(input_business, inputObjectIdBusiness)
	if error_updatename_mongo != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar la zona horaria, detalle: " + error_updatename_mongo.Error(), ""
	}

	error_update_pg := business_repository.Pg_Update_TimeZone(input_business.TimeZone, inputObjectIdBusiness)
	if error_update_pg != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar la zona horaria, detalle: " + error_update_pg.Error(), ""
	}

	return 200, false, "", "Zona horaria actualizado correctamente"
}

//DIRECCION
func UpdateAddress_Service(inputObjectIdBusiness int, intpu_mo_business models.Mo_Business) (int, bool, string, string) {

	error_update := address_x_business_repository.Mo_Update(intpu_mo_business, inputObjectIdBusiness)
	if error_update != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar la direccion, detalle: " + error_update.Error(), ""
	}

	go func() {
		address_x_business_repository.Pg_UpdateAddress(intpu_mo_business, inputObjectIdBusiness)
	}()

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizó la dirección a " + intpu_mo_business.Address.FullAddress,
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Direccion actualizada correctamente"
}
func FindAddress_Service(inputObjectIdBusiness int) (int, bool, string, models.Mo_Address) {

	b_address, _ := address_x_business_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", b_address
}

//TIPOS DE COMIDA
func UpdateTypeFood_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_updating_typefood := typefood_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
	if error_updating_typefood != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los tipos de comida, detalle: " + error_updating_typefood.Error(), ""
	}

	error_update_pg := typefood_x_business_repository.Pg_Update(input_mo_business, inputObjectIdBusiness)
	if error_update_pg != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los tipos de comida, detalle: " + error_update_pg.Error(), ""
	}

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizaron los tipos de comida",
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Se registraron los tipos de comida correctamente"
}

//SERVICIOS
func UpdateService_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_update_service := service_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
	if error_update_service != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los servicios, detalle: " + error_update_service.Error(), ""
	}

	error_update_pg := service_x_business_repository.Pg_Update(input_mo_business, inputObjectIdBusiness)
	if error_update_pg != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los servicios, detalle: " + error_update_pg.Error(), ""
	}

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizaron los servicios",
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Se registraron los servicios correctamente"
}

//DELIVERY RANGE
func UpdateDeliveryRange_Service(inputObjectIdBusiness int, b_deliveryrange models.Mo_Delivery) (int, bool, string, string) {

	error_update_deliveryrage := business_repository.Mo_Update_DeliveryRange(b_deliveryrange, inputObjectIdBusiness)
	if error_update_deliveryrage != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los rango de delivery, detalle: " + error_update_deliveryrage.Error(), ""
	}

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizó el rango de delivery",
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Rango de delivery actualizado correctamente"
}
func FindDeliveryRange_Service(inputObjectIdBusiness int) (int, bool, string, string) {

	deliveryRange, _ := business_repository.Mo_Find_DeliveryRange(inputObjectIdBusiness)

	return 200, false, "", deliveryRange
}

//PAYMENTH METHOD
func UpdatePaymenthMeth_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_updating_paymenth := payment_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
	if error_updating_paymenth != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los metodos de pago, detalle: " + error_updating_paymenth.Error(), ""
	}

	error_update_pg := payment_x_business_repository.Pg_Update(input_mo_business, inputObjectIdBusiness)
	if error_update_pg != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los metodos de pago, detalle: " + error_update_pg.Error(), ""
	}

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizaron los métodos de pago",
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Metodos de pagos cargados correctamente"
}

//HORARIO
func UpdateSchedule_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_update_schedule := schedule_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
	if error_update_schedule != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el horario, detalle: " + error_update_schedule.Error(), ""
	}

	error_update_pg := schedule_x_business_repository.Pg_Update(input_mo_business, inputObjectIdBusiness)
	if error_update_pg != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el horario, detalle: " + error_update_pg.Error(), ""
	}

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizó el horario",
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Se registraro el horario correctamente"
}

//CONTACTO
func UpdateContact_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_updating_contact := contact_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
	if error_updating_contact != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los contactos, detalle: " + error_updating_contact.Error(), ""
	}

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se actualizó el contacto",
			"iduser":   inputObjectIdBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Se registraron los medios de contacto correctamente"
}

func FindContact_Service(inputObjectIdBusiness int) (int, bool, string, []models.Mo_Contact) {

	contact_x_business, _ := contact_x_business_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", contact_x_business
}

//COMENTARIO
func AddComment_Service(input_comment models.Mo_Comment) (int, bool, string, string) {

	input_comment.IsVisible = true

	if input_comment.ISToUpdate {

		error_updating_comment := comment_x_business.Mo_Update_MainData(input_comment)
		if error_updating_comment != nil {
			return 500, true, "Error interno en el servidor al intentar actualizar el comentario al negocio, detalle: " + error_updating_comment.Error(), ""
		}

	} else {

		error_updating_comment := comment_x_business.Mo_Add(input_comment)
		if error_updating_comment != nil {
			return 500, true, "Error interno en el servidor al intentar agregar el comentario al negocio, detalle: " + error_updating_comment.Error(), ""
		}

	}

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  input_comment.FullNameComensal + " agregó un comentario 💬 con una calificación de " + strconv.Itoa(input_comment.Stars) + " estrellas",
			"iduser":   input_comment.IDBusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Nuevo comentario 💬",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Agregaste un comentario 💬 con una calificación de " + strconv.Itoa(input_comment.Stars) + " estrellas a " + input_comment.FullNameBusiness,
			"iduser":   input_comment.IDComensal,
			"typeuser": 2,
			"priority": 1,
			"title":    "Nuevo comentario 💬",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 200, false, "", "Se registró el comentario correctamente"
}

func GetCommentsBusiness_Service(input_data_idbusiness int, page_int int64) (int, bool, string, []*models.Mo_Comment) {

	comments, error_find_comments := comment_x_business.Mo_Find(input_data_idbusiness, page_int)
	if error_find_comments != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los comentarios del negocio, detalle: " + error_find_comments.Error(), comments
	}

	return 200, false, "", comments
}

func GetCommentsStadistics_Service(input_data_idbusiness int) (int, bool, string, []interface{}) {

	comments_resume, error_find_comments := comment_x_business.Mo_Find_Resume(input_data_idbusiness)
	if error_find_comments != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los comentarios del negocio, detalle: " + error_find_comments.Error(), comments_resume
	}

	return 200, false, "", comments_resume
}

func GetCommentsComensal_Service(input_data_idbusiness int, page_int int64) (int, bool, string, []*models.Mo_Comment_Comensal) {

	comments_visible, error_find_comments := comment_x_business.Mo_Find_Visible(input_data_idbusiness, page_int)
	if error_find_comments != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los comentarios del negocio, detalle: " + error_find_comments.Error(), comments_visible
	}

	return 200, false, "", comments_visible
}

func GetCommentsOne_Comensal_Service(input_data_idbusiness int, input_data_idcomensal int) (int, bool, string, CommentFound) {

	var comment_found CommentFound

	comment_one, error_find_comments := comment_x_business.Mo_Find_CommentComensal(input_data_idbusiness, input_data_idcomensal)
	if error_find_comments != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los comentarios del negocio, detalle: " + error_find_comments.Error(), comment_found
	}

	if comment_one.Comment != "" {
		comment_found.Hascomment = true
	} else {
		comment_found.Hascomment = false
	}

	comment_found.Comment = comment_one

	return 200, false, "", comment_found
}

func UpdateCommentBusiness_Service(input_idcommment string) (int, bool, string, string) {

	error_updating_comment := comment_x_business.Mo_Update(input_idcommment)
	if error_updating_comment != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el comentario, detalle: " + error_updating_comment.Error(), ""
	}

	return 200, false, "", "Comentario eliminado correctamente"
}

func UpdateCommentComensal_Service(input_idcommment string) (int, bool, string, string) {

	error_updating_comment := comment_x_business.Mo_Update(input_idcommment)
	if error_updating_comment != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el comentario, detalle: " + error_updating_comment.Error(), ""
	}

	return 200, false, "", "Comentario eliminado correctamente"
}

func AddCommentReport_Service(input_comment_reported models.Mo_Comment_Reported) (int, bool, string, string) {

	error_add_report_comment := reports.Mo_Add_Comment(input_comment_reported)
	if error_add_report_comment != nil {
		return 500, true, "Error interno en el servidor al intentar reportar el comentario, detalle: " + error_add_report_comment.Error(), ""
	}

	return 200, false, "", "Comentario reportado exitosamente"
}

func AddBusinessReport_Service(input_comment_reported models.Mo_Business_Reported) (int, bool, string, string) {

	error_add_report_business := reports.Mo_Add_Business(input_comment_reported)
	if error_add_report_business != nil {
		return 500, true, "Error interno en el servidor al intentar reportar el comentario, detalle: " + error_add_report_business.Error(), ""
	}

	return 200, false, "", "Negocio reportado exitosamente"
}

/*----------------------GET DATA OF THE BUSINESS----------------------*/

func GetInformationData_Service(inputidbusiness int) (int, bool, string, models.Mo_Business) {

	business, _ := business_repository.Mo_Find_All_Data(inputidbusiness)

	return 200, false, "", business
}

/*----------------------GET DATA OF THE BUSINESS WITH ONE ENDPOINT----------------------*/

func GetInformationData_a_Comensal_Service(inputidbusiness_from_comensal int) (int, bool, string, models.Mo_Business) {

	business, _ := business_repository.Mo_Find_All_Data(inputidbusiness_from_comensal)

	datas, error_updating_comment := comment_x_business.Mo_Find_Resume(inputidbusiness_from_comensal)
	if error_updating_comment != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el comentario, detalle: " + error_updating_comment.Error(), business
	}

	business.Comments = datas

	return 200, false, "", business
}

/*----------------------GET DATA OF THE BUSINESS WITH ONE ENDPOINT----------------------*/

func UpdateViewInformation_Consumer_Service(input_view models.Mqtt_View_Information) string {

	error_add_view := view.Mo_Add_Information(input_view)
	if error_add_view != nil {
		return "Error interno en el servidor al intentar registrar la vista del negocio, detalle: " + error_add_view.Error()
	}
	return ""
}

func UpdateViewElement_Consumer_Service(input_view models.Mqtt_View_Element) string {

	error_add_view := view.Mo_Add_Element(input_view)
	if error_add_view != nil {
		return "Error interno en el servidor al intentar registrar la vista del negocio, detalle: " + error_add_view.Error()
	}
	return ""
}
