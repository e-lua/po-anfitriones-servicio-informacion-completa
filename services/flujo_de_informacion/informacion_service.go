package informacion

import (

	//REPOSITORIES
	"log"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	address_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/address_x_business"
	banner_x_busines_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/banner_x_business"
	business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/business"
	contact_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/contact_x_business"
	schedule_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/day_x_business"
	payment_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/paymenth_x_business"

	//contact "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/r_contact"
	//day "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/r_day"
	//paymenth "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/r_paymenth"
	//services "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/r_service"
	//typefood "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/r_typefood"
	service_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/service_x_business"
	typefood_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/typefood_x_business"
)

/*----------------------CONSUMER----------------------*/

func UpdateBanners_Consumer_Service(idbanner int, urlphoto string, idbusiness int) error {
	error_add_banner_mo := banner_x_busines_repository.Mo_Update(urlphoto, idbusiness)
	if error_add_banner_mo != nil {
		log.Fatal(error_add_banner_mo)
	}
	return nil
}

/*----------------------SERVICES TO SHOW ALL PUBLIC DATA----------------------*/

func FindPaymenth_Service(idcountry int, idbusiness int) (int, bool, string, []models.Pg_R_PaymentMethod) {

	paymenth, _ := payment_x_business_repository.Pg_Find(idbusiness, idcountry)
	return 200, false, "", paymenth
}

/*
func FindAllService_Service() (int, bool, string, []models.Ar_R_Service) {

	services, _ := services.Ar_Find()

	return 200, false, "", services
}

func FindAllTypeFood_Service(idcountry int) (int, bool, string, []models.Ar_R_TypeFood) {

	typefood, _ := typefood.Ar_Find_ByCountry(idcountry)

	return 200, false, "", typefood
}

func FindAllSchedule_Service() (int, bool, string, []models.Ar_R_Day) {

	day, _ := day.Ar_Find()

	return 200, false, "", day
}

func FindAllContact_Service() (int, bool, string, interface{}) {

	contacto, error_findcontacts := contact.Ar_Find()
	if error_findcontacts != nil {
		return 500, true, "Error interno en el servidor al intentar buscar todos los contatos, detalle: " + error_findcontacts.Error(), contacto
	}

	return 200, false, "", contacto
}
*/
/*----------------------SERVICES TO UPDATE DATA OF BUSINESS----------------------*/

//NOMBRE
func UpdateName_Service(inputObjectIdBusiness int, input_b_name B_Name) (int, bool, string, string) {

	error_updatename_mongo := business_repository.Mo_Update_Name(input_b_name.Name, inputObjectIdBusiness)
	if error_updatename_mongo != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el nombre, detalle: " + error_updatename_mongo.Error(), ""
	}

	/*error_updatename_ar := business_repository.Ar_Update_Name(input_b_name.Name, inputObjectIdBusiness)
	if error_updatename_ar != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el nombre, detalle: " + error_updatename_ar.Error(), ""
	}*/

	return 200, false, "", "Nombre actualizado correctamente"
}
func FindName_Service(inputObjectIdBusiness int) (int, bool, string, string) {

	name, _ := business_repository.Mo_Find_Name(inputObjectIdBusiness)
	return 200, false, "", name
}

//DIRECCION
func UpdateAddress_Service(inputObjectIdBusiness int, input_b_address models.Mo_Address) (int, bool, string, string) {

	error_update := address_x_business_repository.Mo_Update(input_b_address, inputObjectIdBusiness)
	if error_update != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar la direccion, detalle: " + error_update.Error(), ""
	}

	return 200, false, "", "Direccion actualizada correctamente"
}
func FindAddress_Service(inputObjectIdBusiness int) (int, bool, string, models.Mo_Address) {

	b_address, _ := address_x_business_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", b_address
}

//TIPOS DE COMIDA
func UpdateTypeFood_Service(inputObjectIdBusiness int, input_b_typefood []models.Mo_TypeFood) (int, bool, string, string) {

	error_updating_typefood := typefood_x_business_repository.Mo_Update(input_b_typefood, inputObjectIdBusiness)
	if error_updating_typefood != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los tipos de comida, detalle: " + error_updating_typefood.Error(), ""
	}

	/*error_updating_typefood_ar := typefood_x_business_repository.Ar_Add_Edge(input_b_typefood, inputObjectIdBusiness)
	if error_updating_typefood_ar != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los tipos de comida en los nodos, detalle: " + error_updating_typefood_ar.Error(), ""
	}*/

	return 200, false, "", "Se registraron los tipos de comida correctamente"
}
func FindTypeFood_Service(inputObjectIdBusiness int) (int, bool, string, []models.Mo_TypeFood) {

	type_food_x_business, _ := typefood_x_business_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", type_food_x_business
}

//SERVICIOS
func UpdateService_Service(inputObjectIdBusiness int, input_b_service []models.Mo_Service) (int, bool, string, string) {

	error_update_service := service_x_business_repository.Mo_Update(input_b_service, inputObjectIdBusiness)
	if error_update_service != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los servicios, detalle: " + error_update_service.Error(), ""
	}

	/*error_update_service_ar := service_x_business_repository.Ar_Add_Edge(input_b_service, inputObjectIdBusiness)
	if error_update_service_ar != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los servicios en los nodos, detalle: " + error_update_service_ar.Error(), ""
	}*/

	return 200, false, "", "Se registraron los servicios correctamente"
}
func FindService_Service(inputObjectIdBusiness int) (int, bool, string, []models.Mo_Service) {

	service_x_business, _ := service_x_business_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", service_x_business
}

//DELIVERY RANGE
func UpdateDeliveryRange_Service(inputObjectIdBusiness int, b_deliveryrange B_DeliveryRange) (int, bool, string, string) {

	error_update_deliveryrage := business_repository.Mo_Update_DeliveryRange(b_deliveryrange.DeliveryRange, inputObjectIdBusiness)
	if error_update_deliveryrage != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los rango de delivery, detalle: " + error_update_deliveryrage.Error(), ""
	}

	return 200, false, "", "Rango de delivery actualizado correctamente"
}
func FindDeliveryRange_Service(inputObjectIdBusiness int) (int, bool, string, string) {

	deliveryRange, _ := business_repository.Mo_Find_DeliveryRange(inputObjectIdBusiness)

	return 200, false, "", deliveryRange
}

//PAYMENTH METHOD
func UpdatePaymenthMeth_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, int, models.Mo_Business) {

	error_updating_paymenth := payment_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
	if error_updating_paymenth != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los metodos de pago, detalle: " + error_updating_paymenth.Error(), inputObjectIdBusiness, input_mo_business
	}
	error_update_pg := payment_x_business_repository.Pg_Update(input_mo_business, inputObjectIdBusiness)
	if error_update_pg != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los metodos de pago en pg, detalle: " + error_update_pg.Error(), inputObjectIdBusiness, input_mo_business
	}

	return 200, false, "", inputObjectIdBusiness, input_mo_business
}
func FindPaymenthMeth_Service(inputObjectIdBusiness int) (int, bool, string, []models.Mo_PaymenthMeth) {

	payment_x_business, _ := payment_x_business_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", payment_x_business
}

//HORARIO
func UpdateSchedule_Service(inputObjectIdBusiness int, b_schedule []models.Mo_Day) (int, bool, string, string) {

	error_update_schedule := schedule_x_business_repository.Mo_Update(b_schedule, inputObjectIdBusiness)
	if error_update_schedule != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el horario, detalle: " + error_update_schedule.Error(), ""
	}

	return 200, false, "", "Se registraro el horario correctamente"
}
func FindSchedule_Service(inputObjectIdBusiness int) (int, bool, string, []models.Mo_Day) {

	day_x_business, _ := schedule_x_business_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", day_x_business
}

//CONTACTO
func UpdateContact_Service(inputObjectIdBusiness int, b_contact []models.Mo_Contact) (int, bool, string, string) {

	error_updating_contact := contact_x_business_repository.Mo_Update(b_contact, inputObjectIdBusiness)
	if error_updating_contact != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los contactos, detalle: " + error_updating_contact.Error(), ""
	}

	return 200, false, "", "Se registraron los medios de contacto correctamente"
}

func FindContact_Service(inputObjectIdBusiness int) (int, bool, string, []models.Mo_Contact) {

	contact_x_business, _ := contact_x_business_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", contact_x_business
}

func FindBanner_Service(inputObjectIdBusiness int) (int, bool, string, []models.Mo_Banner) {

	banner_x_business, _ := banner_x_busines_repository.Mo_Find(inputObjectIdBusiness)

	return 200, false, "", banner_x_business
}

/*----------------------GET DATA OF THE BUSINESS----------------------*/

/*----------------------GET DATA OF THE BUSINESS WITH ONE ENDPOINT----------------------*/

func GetInformationData_Service(inputidbusiness int) (int, bool, string, models.Mo_Business) {

	business, _ := business_repository.Mo_Find_All_Data(inputidbusiness)

	return 200, false, "", business
}
