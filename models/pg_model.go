package models

type Pg_R_PaymentMethod struct {
	IDPaymenth  int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`
	HasNumber   bool   `json:"hasnumber"`
	IsAvailable bool   `json:"available"`
}

type Pg_R_Schedule struct {
	IDSchedule int    `json:"idschedule"`
	Name       string `json:"name"`
	Starttime  string `json:"starttime"`
	Endtime    string `json:"endtime"`
	Available  bool   `json:"available"`
}

type Pg_PaymentMethod_X_Business struct {
	IDPaymenth  int
	IDBusiness  int
	IsAvailable bool
}

type Pg_R_Service struct {
	IDservice   int     `json:"id"`
	Name        string  `json:"name"`
	Pricing     float32 `json:"price"`
	TypeMoney   int     `json:"typemoney"`
	Url         string  `json:"url"`
	IsAvailable bool    `json:"available"`
}

type Pg_R_TypeFood struct {
	IDTypefood  int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	IsAvailable bool   `json:"available"`
}

type Pg_BasicData struct {
	IsOpen   bool   `json:"isopen"`
	Name     string `json:"name"`
	TimeZone string `json:"timezone"`
}
