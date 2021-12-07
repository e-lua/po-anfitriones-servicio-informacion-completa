package models

import "time"

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
	Name           string            `bson:"name" json:"name"`
	DeliveryRange  string            `bson:"deliveryrange" json:"deliveryrange"`
	IsOpen         bool              `bson:"isopen" json:"isopen"`
	Contact        []Mo_Contact      `bson:"contact" json:"contact"`
	DailySchedule  []Mo_Day          `bson:"schedule" json:"schedule"`
	Address        Mo_Address        `bson:"address" json:"address"`
	Banner         []Mo_Banner       `bson:"banners" json:"banners"`
	TypeOfFood     []Mo_TypeFood     `bson:"typeoffood" json:"typeoffood"`
	Services       []Mo_Service      `bson:"services" json:"services"`
	PaymentMethods []Mo_PaymenthMeth `bson:"paymentmethods" json:"paymentmethods"`
}

type Mo_Banner struct {
	IdBanner string `bson:"id" json:"id"`
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
	Name       string `bson:"name" json:"name"`
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
	IsOpen         bool      `json:"isopen"`
}
