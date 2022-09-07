package informacion_web

import (

	//REPOSITORIES

	models "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/business"
	comment_x_business "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/comments"
	post_x_business "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/posts"
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

func Web_GetPost_Service(input_data_idbusiness int, limit_int int64) (int, bool, string, []*models.Mo_Post) {
	posts, error_add := post_x_business.Mo_Find(input_data_idbusiness, limit_int)
	if error_add != nil {
		return 500, true, "Error interno en el servidor al intentar obtener los posts, detalle: " + error_add.Error(), posts
	}

	return 200, false, "", posts
}
