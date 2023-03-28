package model

import (
	"database/sql/driver"
	"strings"
)

type Array []string

// scan 从数据库读取数据后，对其进行处理，获得go类型的变量
func (m *Array) Scan(val interface{}) error {
	value := val.([]uint8)
	data := strings.Split(string(value), "|")
	*m = data
	return nil
}

// value 将数据存到数据库时，对数据进行处理，获得数据库支持的类型
func (m *Array) Value() (driver.Value, error) {
	str := strings.Join(*m, "|")
	return str, nil
}
