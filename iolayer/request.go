package iolayer

type PulRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PutTimoutRequest struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Timeout int    `json:"timeout"`
}
