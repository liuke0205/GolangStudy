package containers

// JSONSerializer 提供JSON序列化
type JSONSerializer interface {
	// ToJSON 输出容器元素的JSON表示
	ToJSON() ([]byte, error)
	// MarshalJSON @implements json.Marshaler
	MarshalJSON() ([]byte, error)
}

// JSONDeserializer 提供JSON反序列化
type JSONDeserializer interface {
	// FromJSON 从输入的JSON表示填充容器的元素
	FromJSON([]byte) error
	// UnmarshalJSON @implements json.Unmarshaler
	UnmarshalJSON([]byte) error
}
