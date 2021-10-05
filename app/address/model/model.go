package model

type Create_Response struct {
	Status  bool    `json:"Status"`
	Code    int     `json:"Code,omitempty"`
	Message string  `json:"Message,omitempty"`
	Datas   Address `json:"Address,omitempty"`
}

type Inquiry_Response struct {
	Status  bool      `json:"Status"`
	Code    int       `json:"Code,omitempty"`
	Message string    `json:"Message,omitempty"`
	Datas   []Address `json:"Address,omitempty"`
}

type Update_Response struct {
	Status  bool    `json:"Status"`
	Code    int     `json:"Code,omitempty"`
	Message string  `json:"Message,omitempty"`
	Datas   Address `json:"Address,omitempty"`
}
type Delete_Response struct {
	Status  bool   `json:"Status"`
	Code    int    `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}
