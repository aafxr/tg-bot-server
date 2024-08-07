package types

type Response struct {
	Ok      bool        `json:"ok"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}
