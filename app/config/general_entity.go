package config

type GetResponseSuccess struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type PostResponseSuccess struct {
	Error  bool   `json:"error"`
	ReffId string `json:"reff_id"`
}

type ResponseError struct {
	Error   bool   `json:"error"`
	ReffId  string `json:"reff_id"`
	Message string `json:"message"`
}
