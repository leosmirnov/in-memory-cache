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

type KeyValue struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration int    `json:"expiration"`
}

type ErrorBody struct {
	Error ErrorData `json:"error"`
}

type ErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
