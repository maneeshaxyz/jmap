package jmap

type JmapError struct {
	Type        string `json:"type"`
	Status      int    `json:"status,omitempty"`
	Detail      string `json:"detail,omitempty"`
	Limit       string `json:"limit,omitempty"`
	Description string `json:"description,omitempty"`
}
