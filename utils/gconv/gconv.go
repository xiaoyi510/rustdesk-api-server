package gconv

func Bytes(data interface{}) []byte {
	if data == nil {
		return nil
	}
	switch value := data.(type) {
	case string:
		return []byte(value)
	case []byte:
		return value
	default:
		return nil
	}
}
