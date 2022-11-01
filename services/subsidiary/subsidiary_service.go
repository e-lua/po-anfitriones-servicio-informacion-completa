package subsidiary

import (
	"github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"
	business_repository "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/repositories/business"
)

func GetSubsidiaries_Service(idbusiness int) (int, bool, string, []*models.Mo_Subsidiary) {

	subsidiary, error_updating_comment := business_repository.Mo_Find_Subsidiaries(idbusiness)
	if error_updating_comment != nil {
		return 500, true, "Error interno en el servidor al intentar listar las sucursales, detalle: " + error_updating_comment.Error(), subsidiary
	}

	return 200, false, "", subsidiary
}
