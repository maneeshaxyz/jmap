package jmap

// Contains all the constants required for JMAP business logic.
// Currently all the constants satisfy the minimum reccommendations in the RFC

const (
	MaxSizeUpload         = 50_000_000
	MaxConcurrentUpload   = 4
	MaxSizeRequest        = 10_000_000
	MaxConcurrentRequests = 4
	MaxCallsInRequest     = 16
	MaxObjectsInGet       = 500
	MaxObjectsInSet       = 500
)

var CollationAlgorithms = []string{"i;ascii-numeric"}

const (
	DefaultUsername       = "user@example.com"
	DefaultPrimaryAccount = "account1"
	DefaultState          = "some_state_string"
)

const (
	APIBaseURL         = "https://api.example.com/jmap/"
	DownloadURLTmpl    = APIBaseURL + "download/{accountId}/{blobId}/{type}/{name}"
	UploadURLTmpl      = APIBaseURL + "upload/{accountId}"
	EventSourceURLTmpl = APIBaseURL + "events?types={types}&closeafter={closeafter}&ping={ping}"
)

var DefaultAccounts = Accounts{
	"account1": {
		Name:       "user@example.com",
		IsPersonal: true,
		IsReadOnly: false,
		AccountCapabilities: map[string]any{
			"urn:ietf:params:jmap:core": struct{}{}, // empty object
		},
	},
}
