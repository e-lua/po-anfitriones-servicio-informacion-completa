package informacion_web

import (

	//REPOSITORIES

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/business"
	comment_x_business "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/comments"
)

func Web_GetInformationData_a_Comensal_Service(uniquename string) (int, bool, string, models.Mo_Business) {

	business, _ := business_repository.Web_Mo_Find_All_Data(uniquename)

	datas, error_updating_comment := comment_x_business.Mo_Find_Resume(business.IdBusiness)
	if error_updating_comment != nil {
		return 500, true, "Error interno en el servidor al intentar actualizar el comentario, detalle: " + error_updating_comment.Error(), business
	}

	business.Comments = datas

	return 200, false, "", business
}
