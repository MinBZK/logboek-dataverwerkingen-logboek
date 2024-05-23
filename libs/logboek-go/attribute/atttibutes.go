package attribute

type Attribute struct {
	Key, Value string
}

func New(key, value string) Attribute {
	return Attribute{Key: key, Value: value}
}
