package serialization

type SerializationFactory interface {
	Marshal(v ...interface{}) ([]byte, error)
	Unmarshal(data []byte, v ...interface{}) error
}