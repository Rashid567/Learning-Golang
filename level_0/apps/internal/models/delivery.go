package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func (s *Delivery) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &s)
		return nil
	case string:
		json.Unmarshal([]byte(v), &s)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (s *Delivery) Value() (driver.Value, error) {
	return json.Marshal(s)
}
