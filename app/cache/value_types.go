package cache

import (
	"encoding/json"
	"errors"
)

type RedisValueType int

const (
	jsonValue RedisValueType = iota
	stringValue
)

func (rvt RedisValueType) setValue(value interface{}) (v string, err error) {
	if rvt == jsonValue {
		return setJsonValue(value)
	} else if rvt == stringValue {
		return setStringValue(value)
	}
	return
}

func setJsonValue(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func setStringValue(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	default:
		return "", errors.New("value is not string type")
	}
}

func (rvt RedisValueType) getValue(value string, v interface{}) (err error) {
	if rvt == jsonValue {
		return json.Unmarshal([]byte(value), v)
	} else if rvt == stringValue {
		v = value
		return
	}
	return
}
