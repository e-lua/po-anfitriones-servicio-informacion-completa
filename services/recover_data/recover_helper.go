package recover

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Response_GetAllBusiness struct {
	Error     bool          `json:"error"`
	DataError string        `json:"dataError"`
	Data      []interface{} `json:"data"`
}

type Response_GetOneBusiness struct {
	Error     bool        `json:"error"`
	DataError string      `json:"dataError"`
	Data      interface{} `json:"data"`
}
