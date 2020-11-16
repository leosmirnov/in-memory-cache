package types

type DataBody struct {
	Data interface{} `json:"data"`
}

type Value struct {
	Value string `json:"value"`
}

type KeysList struct {
	Keys []string `json:"keys"`
}

type ErrorBody struct {
	Error ErrorData `json:"error"`
}

type ErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type KeyValue struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration string `json:"expiration"`
}

func (kv KeyValue) Validate() bool {
	if kv.Key == "" || kv.Value == "" {
		return false
	}
	return true
}
