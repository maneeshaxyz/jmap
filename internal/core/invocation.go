package jmap

type Invocation struct {
	Name         string
	Arguments    map[string]any
	MethodCallId string
}
