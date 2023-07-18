package clients

import jsoniter "github.com/json-iterator/go"

var jsonClient = NewJsonClient()

func NewJsonClient() jsoniter.API {

	return jsoniter.ConfigCompatibleWithStandardLibrary
}

func MarshallJson(data interface{}) ([]byte, error) {
	return jsonClient.Marshal(data)
}

func UnmarshallJson[T any](data []byte) (*T, error) {
	t := new(T)
	err := jsonClient.Unmarshal(data, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
