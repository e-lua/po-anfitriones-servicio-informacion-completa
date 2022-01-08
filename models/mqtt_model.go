package models

type Mqtt_PaymentMethod struct {
	IdBusiness     int      `json:"idbusiness"`
	Idbusiness_pg  []int    `json:"idbusiness_pg"`
	Idpaymenth_pg  []int    `json:"idpaymenth_pg"`
	PhoneNumber    []string `json:"phonenumber_pg"`
	Isavailable_pg []bool   `json:"isavailable_pg"`
}

type Mqtt_Schedule struct {
	IdBusiness     int      `json:"idbusiness"`
	Idbusiness_pg  []int    `json:"idbusiness_pg"`
	Idschedule_pg  []int    `json:"idschedule_pg"`
	Isavailable_pg []bool   `json:"isavailable_pg"`
	Starttime_pg   []string `json:"starttime_pg"`
	Endtime_pg     []string `json:"endtime_pg"`
}

type Mqtt_Service struct {
	IdBusiness     int       `json:"idbusiness"`
	Idbusiness_pg  []int     `json:"idbusiness_pg"`
	Idservice_pg   []int     `json:"idservice_pg"`
	Pricing_pg     []float32 `json:"pricing_pg"`
	TypeMoney_pg   []int     `json:"typemoney_pg"`
	Isavailable_pg []bool    `json:"isavailable_pg"`
}

type Mqtt_TypeFood struct {
	IdBusiness     int    `json:"idbusiness"`
	Idbusiness_pg  []int  `json:"idbusiness_pg"`
	Idtypefood_pg  []int  `json:"Idtypefood_pg"`
	Isavailable_pg []bool `json:"isavailable_pg"`
}

type Mqtt_Name struct {
	IdBusiness int    `json:"idbusiness"`
	Name       string `json:"name"`
}

type Mqtt_Address struct {
	IdBusiness int     `json:"idbusiness"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type Mqtt_TimeZone struct {
	IdBusiness int    `json:"idbusiness"`
	TimeZone   string `json:"timezone"`
}
