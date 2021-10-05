package model

type Create_Response struct {
	Status  bool     `json:"Status"`
	Code    int      `json:"Code,omitempty"`
	Message string   `json:"Message,omitempty"`
	Datas   *Address `json:"Datas,omitempty"`
}

type Inquiry_Response struct {
	Status  bool       `json:"Status"`
	Code    int        `json:"Code,omitempty"`
	Message string     `json:"Message,omitempty"`
	Datas   *[]Address `json:"Datas,omitempty"`
}

type Update_Response struct {
	Status  bool     `json:"Status"`
	Code    int      `json:"Code,omitempty"`
	Message string   `json:"Message,omitempty"`
	Datas   *Address `json:"Datas,omitempty"`
}
type Delete_Response struct {
	Status  bool   `json:"Status"`
	Code    int    `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}
