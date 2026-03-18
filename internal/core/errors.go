package jmap

const (
	ErrorTypeNotJSON           = "urn:ietf:params:jmap:error:notJSON"
	ErrorTypeNotRequest        = "urn:ietf:params:jmap:error:notRequest"
	ErrorTypeUnknownCapability = "urn:ietf:params:jmap:error:unknownCapability"
	ErrorTypeLimit             = "urn:ietf:params:jmap:error:limit"
)

type JmapError struct {
	Type   string `json:"type"`
	Status int    `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
	Limit  string `json:"limit,omitempty"`
}
