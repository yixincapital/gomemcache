package serialization

import "encoding/json"

/*json存储*/
type JsonProvider struct {
	SerializationFactory
}

func (provider *JsonProvider) Marshal(v ...interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	return bytes, err
}

func (provider *JsonProvider) Unmarshal(data []byte, v ...interface{}) error {
	err := json.Unmarshal(data, &v)
	return err
}
