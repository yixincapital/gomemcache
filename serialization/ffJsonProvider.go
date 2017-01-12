package serialization

import "github.com/pquerna/ffjson/ffjson"

/*json存储*/
type FFJsonProvider struct {
	SerializationFactory
}

func (provider *FFJsonProvider) Marshal(v ...interface{}) ([]byte, error) {
	bytes, err := ffjson.Marshal(v)
	return bytes, err
}

func (provider *FFJsonProvider) Unmarshal(data []byte, v ...interface{}) error {
	err := ffjson.Unmarshal(data, &v)
	return err
}
