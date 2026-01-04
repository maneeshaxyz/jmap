package jmap

type Session struct {
	Capabilities    Capabilities      `json:"capabilities"`
	Accounts        Accounts          `json:"accounts"`
	PrimaryAccounts map[string]string `json:"primaryAccounts,omitempty"`
	Username        string            `json:"username"`
	APIURL          string            `json:"apiUrl"`
	DownloadURL     string            `json:"downloadUrl"`
	UploadURL       string            `json:"uploadUrl"`
	EventSourceURL  string            `json:"eventSourceUrl"`
	State           string            `json:"state"`
}

// set to map[string]any so clients can ignore unknown Capabilities
type Capabilities map[string]any

type CoreCapability struct {
	MaxSizeUpload         uint64   `json:"maxSizeUpload"`
	MaxConcurrentUpload   uint64   `json:"maxConcurrentUpload"`
	MaxSizeRequest        uint64   `json:"maxSizeRequest"`
	MaxConcurrentRequests uint64   `json:"maxConcurrentRequests"`
	MaxCallsInRequest     uint64   `json:"maxCallsInRequest"`
	MaxObjectsInGet       uint64   `json:"maxObjectsInGet"`
	MaxObjectsInSet       uint64   `json:"maxObjectsInSet"`
	CollationAlgorithms   []string `json:"collationAlgorithms"`
}
