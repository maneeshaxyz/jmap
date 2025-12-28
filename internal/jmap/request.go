package jmap

type JmapRequest struct {
	Using       []string     `json:"using"`
	MethodCalls []Invocation `json:"methodCalls"`
}
