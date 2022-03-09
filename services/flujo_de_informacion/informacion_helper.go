package informacion

import "github.com/Aphofisis/po-anfitriones-servicio-informacion-completa/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

//BUSINESSDATA
type ResponseBusiness struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      models.Mo_Business `json:"data"`
}

//NAME
type B_Name struct {
	Name string `json:"name"`
}

//UNIQUENAME
type B_Uniquename struct {
	Uniquename string `json:"uniquename"`
}

//ISOPEN
type B_Open struct {
	IsOpen bool `bson:"isopen" json:"isopen"`
}

//PHOTO-BANNER&PROFILE
type ResponseBannerProfile struct {
	Error     bool                 `json:"error"`
	DataError string               `json:"dataError"`
	Data      B_PhotoBannerProfile `json:"data"`
}

type B_PhotoBannerProfile struct {
	URLProfile string `json:"urlProfile"`
	URLBanner  string `json:"urlBanner"`
}

//ADDRESS
type ResponseAddress struct {
	Error     bool              `json:"error"`
	DataError string            `json:"dataError"`
	Data      models.Mo_Address `json:"data"`
}

type B_Address struct {
	Latitude         float32 `json:"latitude"`
	Longitude        float32 `json:"longitude"`
	FullAddress      string  `json:"fullAddress"`
	PostalCode       int     `json:"postalCode"`
	ReferenceAddress string  `json:"referenceAddress"`
}

//INTERFACE
type ResponseInterface_FromComensal struct {
	Error     bool        `json:"error"`
	DataError string      `json:"dataError"`
	Data      interface{} `json:"data"`
}

//DELIVERYRANGE
type ResponseDeliveryRange struct {
	Error     bool            `json:"error"`
	DataError string          `json:"dataError"`
	Data      B_DeliveryRange `json:"data"`
}

type B_DeliveryRange struct {
	DeliveryRange string `json:"deliveryRange"`
}

//PAYMENTMETH
type ResponsePaymentMeth struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      []models.Mo_PaymenthMeth `json:"data"`
}

//BANNER
type ResponseBanner struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      []models.Mo_Banner `json:"data"`
}

//CONTACT
type ResponseContact struct {
	Error     bool                `json:"error"`
	DataError string              `json:"dataError"`
	Data      []models.Mo_Contact `json:"data"`
}

//CONTACT_ALL
type ResponseContactAll struct {
	Error     bool        `json:"error"`
	DataError string      `json:"dataError"`
	Data      interface{} `json:"data"`
}

//BUSINESS FULL DATA

type ResponseBusinessData struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      models.Mo_Business `json:"data"`
}

/*
type ResponseWithStructBusiness struct {
	NameBusiness      string                            `json:"name"`
	LatitudeBusiness  float32                           `json:"latitude"`
	PostalCode        int                               `json:"postalCode"`
	LongitudeBusiness float32                           `json:"longitude"`
	Fulladdress       string                            `json:"fullAddress"`
	ReferenceAddress  string                            `json:"referenceAddress"`
	Banner            []models.Pg_BannerXBusiness       `json:"banner"`
	TypeOfFood        []models.Pg_TypeFoodXBusiness     `json:"typeOfFood"`
	Services          []models.Pg_ServiceXBusiness      `json:"services"`
	DeliveryRange     string                            `json:"deliveryRange"`
	PaymentMethods    []models.Pg_PaymenthMethXBusiness `json:"paymentMethods"`
	DailySchedule     []models.Pg_DayXBusiness          `json:"schedule"`
	PhoneContact      []models.Pg_ContactxBusiness      `json:"contact"`
}
*/

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type JWT struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type ResponseJWT_Comensal struct {
	Error     bool         `json:"error"`
	DataError string       `json:"dataError"`
	Data      JWT_Comensal `json:"data"`
}

type JWT_Comensal struct {
	Phone      int `json:"phone"`
	Country    int `json:"country"`
	IDComensal int ` json:"comensal"`
}

/*type ResponseCBusinessBasicData_Mo struct {
	Error     bool                            `json:"error"`
	DataError string                          `json:"dataError"`
	Data      []models.Mo_Business_Basic_Data `json:"data"`
}*/

//OBTENER UNIQUENAMES

type Response_Uniquenames struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Post_Notification struct {
	Message  string `bson:"message" json:"message,omitempty"`
	IDUser   int    `bson:"iduser" json:"iduser,omitempty"`
	TypeUser int    `bson:"typeuser" json:"typeuser,omitempty"`
}

type Response_Comments struct {
	Error     bool                 `json:"error"`
	DataError string               `json:"dataError"`
	Data      []*models.Mo_Comment `json:"data"`
}

type Response_Comment_Resume struct {
	Error     bool        `json:"error"`
	DataError string      `json:"dataError"`
	Data      interface{} `json:"data"`
}

type Response_Comments_Comensal struct {
	Error     bool                          `json:"error"`
	DataError string                        `json:"dataError"`
	Data      []*models.Mo_Comment_Comensal `json:"data"`
}
