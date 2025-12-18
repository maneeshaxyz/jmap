package main

// CoreCapability limits
const (
	MaxSizeUpload         = 50_000_000
	MaxConcurrentUpload   = 4
	MaxSizeRequest        = 10_000_000
	MaxConcurrentRequests = 4
	MaxCallsInRequest     = 16
	MaxObjectsInGet       = 500
	MaxObjectsInSet       = 500
)

// Collation algorithms
var CollationAlgorithms = []string{"i;ascii-numeric"}

// User and account info
const (
	DefaultUsername       = "user@example.com"
	DefaultPrimaryAccount = "account1"
	DefaultState          = "some_state_string"
)

// URLs
const (
	APIBaseURL         = "https://api.example.com/jmap/"
	DownloadURLTmpl    = APIBaseURL + "download/{accountId}/{blobId}/{type}/{name}"
	UploadURLTmpl      = APIBaseURL + "upload/{accountId}"
	EventSourceURLTmpl = APIBaseURL + "events?types={types}&closeafter={closeafter}&ping={ping}"
)
