package strategy

import (
	"encoding/json"
	"os"
)

type JsonSave struct{}

func (js *JsonSave) Save(name string, data []byte) error {
	jsonData, _ := json.Marshal(data)
	return os.WriteFile(name, jsonData, 0644)
}
