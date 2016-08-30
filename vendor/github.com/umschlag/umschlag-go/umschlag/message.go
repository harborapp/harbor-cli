package umschlag

// Message represents a standard response.
type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
