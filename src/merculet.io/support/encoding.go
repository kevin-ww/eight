package support

import (
	"encoding/json"
)

func UnMarshal(data []byte, target interface{}) (interface{}, error) {
	e := json.Unmarshal(data, target)
	if e != nil {
		return nil, e
	}
	return target, nil
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
