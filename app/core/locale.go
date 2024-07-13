package core

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Locale string

const (
	FR Locale = "FR"
	EN Locale = "EN"
)

type LocalizedMessage map[Locale]string

func NewLocalizedMessage() LocalizedMessage {
	return LocalizedMessage{}
}

func (m *LocalizedMessage) Add(locale Locale, message string) {
	(*m)[locale] = message
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (l *LocalizedMessage) Scan(value interface{}) error {
	// Deserialize the value from the database into the Address struct
	// For example, if the value is stored as a JSON string, you would unmarshal it
	if value == nil {
		l = nil
		return nil
	}
	localizedMessageJSON, ok := value.([]byte)
	if !ok {
		return errors.New("unexpected type for localizedMessage")
	}
	return json.Unmarshal(localizedMessageJSON, l)
}

// Value return json value, implement driver.Valuer interface
func (l LocalizedMessage) Value() (driver.Value, error) {
	localizedMessageJSON, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	return localizedMessageJSON, nil
}
