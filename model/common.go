package model

import (
        "database/sql/driver"
        "encoding/json"

        "github.com/pkg/errors"
)

type StringSlice []string

func (p StringSlice) Value() (driver.Value, error) {
        data, err := json.Marshal(p)
        if err != nil {
                return nil, errors.Wrap(err, "json.Marshal")
        }
        return data, nil
}

func (p *StringSlice) Scan(src interface{}) error {
        data, ok := src.([]byte)
        if !ok {
                return nil
        }
        return json.Unmarshal(data, p)
}

type IntSlice []int

func (p IntSlice) Value() (driver.Value, error) {
        data, err := json.Marshal(p)
        if err != nil {
                return nil, errors.Wrap(err, "json.Marshal")
        }
        return data, nil
}

func (p *IntSlice) Scan(src interface{}) error {
        data, ok := src.([]byte)
        if !ok {
                return nil
        }
        return json.Unmarshal(data, p)
}
