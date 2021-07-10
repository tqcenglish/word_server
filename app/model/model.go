package model

type ModbusConfig struct {
	Name string `json:"name"`

	Low       string `json:"low"` // Low 由 LowNumber 与 LowType 构成，可以直接填写
	LowNumber string `json:"lowNumber"`
	LowType   string `json:"lowType"`

	High       string `json:"high"` // High 由 HighNumber 与 HighType 构成，可以直接填写
	HighNumber string `json:"highNumber"`
	HighType   string `json:"highType"`
}

type DIS struct {
	DIS []ModbusConfig `json:"dis"`
}

type PlayMP3Param struct {
	URLValue string `json:"url"`
}
