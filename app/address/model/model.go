package model

import (
	"encoding/json"
	"fiber-crud/pkg/utils"

	"github.com/shopspring/decimal"
)

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

type TestDecimal_Request struct {
	Number1 decimal.Decimal `json:"Number1"`
	Number2 decimal.Decimal `json:"Number2"`
}

type TestDecimal_Response struct {
	Number1      decimal.Decimal     `json:"-"`
	Number1_Json json.Number         `json:"Number1"`
	Number2      decimal.Decimal     `json:"-"`
	Number2_Json json.Number         `json:"Number2"`
	Total        decimal.NullDecimal `json:"-"`
	Total_Json   json.Number         `json:"Total"`
}

func (t *TestDecimal_Response) JsonNumber() {
	t.Number1_Json = utils.DecimalToJsonNumber(&t.Number1, 2)
	t.Number2_Json = utils.DecimalToJsonNumber(&t.Number2, 2)
	t.Total_Json = utils.DecimalNullToJsonNumber(&t.Total, 2)
}
