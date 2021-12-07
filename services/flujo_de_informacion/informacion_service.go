package informacion

import (

	//REPOSITORIES
	"log"
	"time"

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	address_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/address_x_business"
	banner_x_busines_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/banner_x_business"
	business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/business"
	contact_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/contact_x_business"
	schedule_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/day_x_business"
	payment_x_business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/paymenth_x_business"
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

/*----------------------SERVICES TO UPDATE DATA OF BUSINESS----------------------*/

//NOMBRE
func UpdateName_Service(inputObjectIdBusiness int, input_b_name B_Name) (int, bool, string, string) {

	error_updatename_mongo := business_repository.Mo_Update_Name(input_b_name.Name, inputObjectIdBusiness)
	if error_updatename_mongo != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el nombre, detalle: " + error_updatename_mongo.Error(), ""
	}

	go func() {
		business_repository.Pg_UpdateName(input_b_name.Name, inputObjectIdBusiness)
	}()

	return 200, false, "", "Nombre actualizado correctamente"
}
func FindName_Service(inputObjectIdBusiness int) (int, bool, string, string) {

	name, _ := business_repository.Mo_Find_Name(inputObjectIdBusiness)
	return 200, false, "", name
}

//ISOPEN
func UpdateOpen_Service(inputObjectIdBusiness int, input_b_open B_Open) (int, bool, string, string) {

	error_updatename_mongo := business_repository.Mo_Update_Open(input_b_open.IsOpen, inputObjectIdBusiness)
	if error_updatename_mongo != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el nombre, detalle: " + error_updatename_mongo.Error(), ""
	}

	go func() {
		business_repository.Pg_UpdateIsOpen(input_b_open.IsOpen, inputObjectIdBusiness)
	}()

	return 200, false, "", "Nombre actualizado correctamente"
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

	go func() {
		typefood_x_business_repository.Pg_Update(input_mo_business, inputObjectIdBusiness)
	}()

	return 200, false, "", "Se registraron los tipos de comida correctamente"
}
func FindTypeFood_Service(idcountry int, idbusiness int) (int, bool, string, []models.Pg_R_TypeFood) {

	typefood, _ := typefood_x_business_repository.Pg_Find(idbusiness, idcountry)

	return 200, false, "", typefood
}

//SERVICIOS
func UpdateService_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_update_service := service_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
	if error_update_service != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los servicios, detalle: " + error_update_service.Error(), ""
	}

	go func() {
		service_x_business_repository.Pg_Update(input_mo_business, inputObjectIdBusiness)
	}()

	return 200, false, "", "Se registraron los servicios correctamente"
}
func FindService_Service(idcountry int, idbusiness int) (int, bool, string, []models.Pg_R_Service) {

	services, _ := service_x_business_repository.Pg_Find(idbusiness, idcountry)

	return 200, false, "", services
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
func UpdatePaymenthMeth_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_updating_paymenth := payment_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
	if error_updating_paymenth != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar los metodos de pago, detalle: " + error_updating_paymenth.Error(), ""
	}

	go func() {
		payment_x_business_repository.Pg_Update(input_mo_business, inputObjectIdBusiness)
	}()

	time.Sleep(2 * time.Second)

	return 200, false, "", "Metodos de pagos cargados correctamente"
}
func FindPaymenth_Service(idcountry int, idbusiness int) (int, bool, string, []models.Pg_R_PaymentMethod) {

	paymenth, _ := payment_x_business_repository.Pg_Find(idbusiness, idcountry)
	return 200, false, "", paymenth
}

//HORARIO
func UpdateSchedule_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_update_schedule := schedule_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
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
func UpdateContact_Service(inputObjectIdBusiness int, input_mo_business models.Mo_Business) (int, bool, string, string) {

	error_updating_contact := contact_x_business_repository.Mo_Update(input_mo_business, inputObjectIdBusiness)
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

func GetInformationData_Service(inputidbusiness int) (int, bool, string, models.Mo_Business) {

	business, _ := business_repository.Mo_Find_All_Data(inputidbusiness)

	return 200, false, "", business
}

/*----------------------GET DATA OF THE BUSINESS WITH ONE ENDPOINT----------------------*/

func GetInformationData_a_Comensal_Service(inputidbusiness_from_comensal int) (int, bool, string, models.Mo_Business) {

	business, _ := business_repository.Mo_Find_All_Data(inputidbusiness_from_comensal)

	return 200, false, "", business
}
