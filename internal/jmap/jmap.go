package jmap

// Builds a Session struct constructed with values in constants.go
func BuildSession() *Session {
	session := &Session{
		Capabilities: Capabilities{
			"urn:ietf:params:jmap:core": CoreCapability{
				MaxSizeUpload:         MaxSizeUpload,
				MaxConcurrentUpload:   MaxConcurrentUpload,
				MaxSizeRequest:        MaxSizeRequest,
				MaxConcurrentRequests: MaxConcurrentRequests,
				MaxCallsInRequest:     MaxCallsInRequest,
				MaxObjectsInGet:       MaxObjectsInGet,
				MaxObjectsInSet:       MaxObjectsInSet,
				CollationAlgorithms:   CollationAlgorithms,
			},
		},
		PrimaryAccounts: map[string]string{
			"urn:ietf:params:jmap:mail": DefaultPrimaryAccount,
		},
		Username:       DefaultUsername,
		APIURL:         APIBaseURL,
		DownloadURL:    DownloadURLTmpl,
		UploadURL:      UploadURLTmpl,
		EventSourceURL: EventSourceURLTmpl,
		State:          DefaultState,
		Accounts:       DefaultAccounts,
	}
	return session
}
