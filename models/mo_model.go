package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*------------------------BASIC DATA FOR SEARCH------------------------*/

/*
type Mo_Business_Basic_Data struct {
	IDBusiness     int                                   `bson:"idBusiness" json:"idBusiness"`
	IDCountry      int                                   `bson:"idCountry" json:"idCountry"`
	Name           string                                `bson:"name" json:"name"`
	Latitude       float32                               `bson:"latitude" json:"latitude"`
	Longitude      float32                               `bson:"longitude" json:"longitude"`
	Open           bool                                  `bson:"isOpen" json:"isOpen"`
	Banner         string                                `bson:"banner" json:"banner"`
	TypeOfFood     []Mo_TypeFoodXBusiness_Basic_Data     `bson:"typeOfFood" json:"typeOfFood"`
	Services       []Mo_ServiceXBusiness_Basic_Data      `bson:"services" json:"services"`
	PaymentMethods []Mo_PaymenthMethXBusiness_Basic_Data `bson:"paymentMethods" json:"paymentMethods"`
}
*/

type Mo_Business struct {
	Description     string            `bson:"description" json:"description"`
	Name            string            `bson:"name" json:"name"`
	TimeZone        string            `bson:"timezone" json:"timezone"`
	DeliveryRange   string            `bson:"deliveryrange" json:"deliveryrange"`
	Delivery        Mo_Delivery       `bson:"delivery" json:"delivery"`
	Contact         []Mo_Contact      `bson:"contact" json:"contact"`
	DailySchedule   []Mo_Day          `bson:"schedule" json:"schedule"`
	Address         Mo_Address        `bson:"address" json:"address"`
	Banner          []Mo_Banner       `bson:"banners" json:"banners"`
	TypeOfFood      []Mo_TypeFood     `bson:"typeoffood" json:"typeoffood"`
	Services        []Mo_Service      `bson:"services" json:"services"`
	PaymentMethods  []Mo_PaymenthMeth `bson:"paymentmethods" json:"paymentmethods"`
	Comments        []interface{}     `bson:"comments" json:"comments"`
	Uniquename      string            `bson:"uniquename" json:"uniquename"`
	LegalIdentity   string            `bson:"legalidentity" json:"legalidentity"`
	TypeSuscription int               `bson:"typesuscription" json:"typesuscription"`
	IVA             float32           `bson:"iva" json:"iva"`
}

type Mo_Delivery struct {
	Meters        int    `bson:"meters" json:"meters"`
	Details       string `bson:"details" json:"details"`
	IsRestriction bool   `bson:"isrestriction" json:"isrestriction"`
}

type Mo_Banner struct {
	IdBanner int    `bson:"id" json:"id"`
	UrlImage string `bson:"url" json:"url"`
}

type Mo_Address struct {
	Latitude         float64 `bson:"latitude" json:"latitude"`
	Longitude        float64 `bson:"longitude" json:"longitude"`
	FullAddress      string  `bson:"fulladdress" json:"fulladdress"`
	PostalCode       int     `bson:"postalcode" json:"postalcode"`
	State            string  `bson:"state" json:"state"`
	City             string  `bson:"city" json:"city"`
	ReferenceAddress string  `bson:"referenceaddress" json:"referenceaddress"`
}

type Mo_Day struct {
	IDDia      int    `bson:"id" json:"id"`
	StarTime   string `bson:"starttime" json:"starttime"`
	EndTime    string `bson:"endtime" json:"endtime"`
	IsAvaiable bool   `bson:"available" json:"available"`
}

type Mo_TypeFood struct {
	IDTypeFood int    `bson:"id" json:"id"`
	Name       string `bson:"name" json:"name"`
	UrlImage   string `bson:"url" json:"url"`
	IsAvaiable bool   `bson:"available" json:"available"`
}

type Mo_Service struct {
	IDService  int     `bson:"id" json:"id"`
	Name       string  `bson:"name" json:"name"`
	Price      float32 `bson:"price" json:"price"`
	Url        string  `bson:"url" json:"url"`
	TypeMoney  int     `bson:"typemoney" json:"typemoney"`
	IsAvaiable bool    `bson:"available" json:"available"`
}

type Mo_PaymenthMeth struct {
	IDPaymenth  int    `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`
	Url         string `bson:"url" json:"url"`
	HasNumber   bool   `bson:"hasnumber" json:"hasnumber"`
	IsAvaiable  bool   `bson:"available" json:"available"`
}

type Mo_Contact struct {
	IDContact   int    `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	DataContact string `bson:"data" json:"data"`
	IsAvaiable  bool   `bson:"available" json:"available"`
}

type Mo_BusinessWorker_Mqtt struct {
	IdBusiness  int       `json:"idbusiness"`
	IdWorker    int       `json:"idworker"`
	IdCountry   int       `json:"country"`
	CodeRedis   int       `json:"code"`
	Name        string    `json:"name"`
	IdRol       int       `json:"rol"`
	LastName    string    `json:"lastname"`
	Phone       int       `json:"phone"`
	Password    string    `json:"password"`
	UpdatedDate time.Time `json:"updatedate"`
}

type Mo_ToBanner_Mqtt struct {
	IdBusiness                int    `bson:"idbusiness" json:"idbusiness"`
	IdBanner_Category_Element int    `json:"idbCE"`
	IdType                    int    `json:"idtype"`
	Url                       string `bson:"url" json:"url"`
}

type Mo_Registro_FromMqtt struct {
	IdBusiness     int       `json:"idbusiness"`
	OrdersRejected int       `json:"ordersrejected"`
	Available      bool      `json:"available"`
	CreatedDate    time.Time `json:"createddate"`
}

type Mo_Business_Cards struct {
	IDBusiness     int               `json:"idbusiness"`
	ICountry       int               `json:"idcountry"`
	Name           string            `json:"name"`
	Latitude       float32           `json:"latitude"`
	Longitude      float32           `json:"longitude"`
	Location       Location          `json:"location"`
	Available      bool              `json:"available"`
	IsOpen         bool              `json:"isopen"`
	OrdersRejected int               `json:"ordersrejected"`
	Banner         []Mo_Banner       `bson:"banners" json:"banners"`
	TypeOfFood     []Mo_TypeFood     `bson:"typeoffood" json:"typeoffood"`
	Services       []Mo_Service      `bson:"services" json:"services"`
	PaymentMethods []Mo_PaymenthMeth `bson:"paymentmethods" json:"paymentmethods"`
}

type Location struct {
	GeoJSONType string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Mo_BusinessBanner_Mqtt struct {
	IDBusiness int       `bson:"idbusiness" json:"idbusiness"`
	Banner     Mo_Banner `bson:"banners" json:"banners"`
}

type Mo_View_Information struct {
	IDBusiness int       `bson:"idbusiness" json:"idbusiness"`
	IDComensal int       `bson:"idcomensal" json:"idcomensal"`
	Date       time.Time `bson:"date" json:"date"`
}

type Mo_View_Element struct {
	IDElement  int       `bson:"idelement" json:"idelement"`
	IDComensal int       `bson:"idcomensal" json:"idcomensal"`
	Date       time.Time `bson:"date" json:"date"`
}

type Mo_Comment struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Stars            int                `bson:"stars" json:"stars"`
	Comment          string             `json:"comment" bson:"comment"`
	IDBusiness       int                `json:"idbusiness" bson:"idbusiness"`
	IDComensal       int                `json:"idcomensal" bson:"idcomensal"`
	FullNameComensal string             `json:"fullnamecomensal" bson:"fullnamecomensal"`
	FullNameBusiness string             `json:"fullnamebusiness" bson:"fullnamebusiness"`
	PhoneComensal    int                `json:"phonecomensal" bson:"phonecomensal"`
	Dateregistered   time.Time          `json:"dateregistered" bson:"dateregistered"`
	IsVisible        bool               `json:"isvisible" bson:"isvisible"`
	ISToUpdate       bool               `json:"istoupdate" bson:"istoupdate"`
}

type Mo_Comment_Comensal struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Stars            int                `bson:"stars" json:"stars"`
	Comment          string             `json:"comment" bson:"comment"`
	IDBusiness       int                `json:"idbusiness" bson:"idbusiness"`
	IDComensal       int                `json:"idcomensal" bson:"idcomensal"`
	FullNameComensal string             `json:"fullnamecomensal" bson:"fullnamecomensal"`
	Dateregistered   time.Time          `json:"dateregistered" bson:"dateregistered"`
}

type Mo_Comment_ComensalFound struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Stars   int                `bson:"stars" json:"stars"`
	Comment string             `json:"comment" bson:"comment"`
}

/*================REPORTES===================*/

type Mo_Comment_Reported struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDComment        string             `bson:"idcomment" json:"idcomment"`
	Stars            int                `bson:"stars" json:"stars"`
	Comment          string             `json:"comment" bson:"comment"`
	IDBusiness       int                `json:"idbusiness" bson:"idbusiness"`
	IDComensal       int                `json:"idcomensal" bson:"idcomensal"`
	FullNameComensal string             `json:"fullnamecomensal" bson:"fullnamecomensal"`
	PhoneComensal    int                `json:"phonecomensal" bson:"phonecomensal"`
	Dateregistered   time.Time          `json:"dateregistered" bson:"dateregistered"`
	IsVisible        bool               `json:"isvisible" bson:"isvisible"`
	Datereported     time.Time          `json:"datereported" bson:"datereported"`
	Reason           string             `json:"reason" bson:"reason"`
	IDReason         int                `json:"idreason" bson:"idreason"`
	WasView          bool               `json:"wasview" bson:"wasview"`
}

type Mo_Business_Reported struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDBusiness   int                `json:"idbusiness" bson:"idbusiness"`
	IDComensal   int                `json:"idcomensal" bson:"idcomensal"`
	Datereported time.Time          `json:"datereported" bson:"datereported"`
	Reason       string             `json:"reason" bson:"reason"`
	IDReason     int                `json:"idreason" bson:"idreason"`
	WasView      bool               `json:"wasview" bson:"wasview"`
}
