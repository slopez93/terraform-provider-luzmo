package dtos

type Find struct {
	Where Where `json:"where"`
}

type Where struct {
	ID string `json:"id"`
}

type LuzmoFindResponse[T interface{}] struct {
	Count int `json:"count"`
	Rows  []T
}

type LuzmoCreateRequest[T interface{}] struct {
	Action     string `json:"action"`
	Version    string `json:"version"`
	Key        string `json:"key"`
	Token      string `json:"token"`
	Properties T      `json:"properties"`
}

type LuzmoUpdateRequest[T interface{}] struct {
	Action     string `json:"action"`
	Version    string `json:"version"`
	Key        string `json:"key"`
	Token      string `json:"token"`
	Id         string `json:"id"`
	Properties T      `json:"properties"`
}

type LuzmoDeleteRequest struct {
	Action  string `json:"action"`
	Version string `json:"version"`
	Key     string `json:"key"`
	Token   string `json:"token"`
	Id      string `json:"id"`
}

type LuzmoFindRequest struct {
	Action  string `json:"action"`
	Version string `json:"version"`
	Key     string `json:"key"`
	Token   string `json:"token"`
	Find    Find   `json:"find"`
}
