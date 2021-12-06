package models

type Mqtt_PaymentMethod struct {
	IdBusiness     int    `json:"idbusiness"`
	Idbusiness_pg  []int  `json:"idbusiness_pg"`
	Idpaymenth_pg  []int  `json:"idpaymenth_pg"`
	Isavailable_pg []bool `json:"isavailable_pg"`
}
